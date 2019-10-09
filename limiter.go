package limiter

import (
	"fmt"
	"sync"
	"time"

	"github.com/OneOfOne/xxhash"
)

const (
	shards = 1024
	sleep  = time.Minute * 10
)

func New(limit int, per time.Duration) (a *Limiter) {
	a = new(Limiter)
	a.hide.list = make([]map[int][]int, shards)
	for x := 0; x < shards; x++ {
		a.hide.list[x] = make(map[int][]int)
	}
	a.limit = limit
	a.period = per
	return
}

type Limiter struct {
	hide struct {
		sync.Mutex
		list []map[int][]int
	}
	period  time.Duration
	limit   int //per period
	current int
}

func (a *Limiter) delete(id int) {
	delete(a.hide.list[id%shards], id)
}

func (a *Limiter) cleaner() {
	for {
		time.Sleep(sleep)
		a.hide.Lock()
		for _, users := range a.hide.list {
			now := int(time.Now().Unix())
			for id, limit := range users {
				if limit[0] < now {
					a.delete(id)
					continue
				}
			}
		}
		a.hide.Unlock()
	}
}

func (a *Limiter) Int(id int) (res bool) {

	defer a.hide.Unlock()
	a.hide.Lock()

	shard := id % shards

	switch a.hide.list[shard][id] {
	case nil:
		//first time
		a.hide.list[shard][id] = []int{
			int(time.Now().Add(a.period).Unix()),
			1,
		}
		return true
	default:
		switch a.hide.list[shard][id][0] < int(time.Now().Unix()) {
		case true:
			//new period
			a.hide.list[shard][id] = []int{
				int(time.Now().Add(a.period).Unix()),
				1,
			}
			return true
		default:
			//ok
			if a.hide.list[shard][id][1] < a.limit {
				a.hide.list[shard][id][1]++
				return true
			}

			//limit
			return false
		}
	}
}

func (a *Limiter) String(sid string) (res bool) {
	return a.Int(int(xxhash.Checksum32([]byte(sid))))
}

func (a *Limiter) Interface(id interface{}) (res bool) {
	switch id.(type) {
	case int:
		return a.Int(id.(int))
	case int8:
		return a.Int(int(id.(int8)))
	case int16:
		return a.Int(int(id.(int16)))
	case int32:
		return a.Int(int(id.(int32)))
	case int64:
		return a.Int(int(id.(int64)))
	case uint:
		return a.Int(int(id.(uint)))
	case uint8:
		return a.Int(int(id.(uint8)))
	case uint16:
		return a.Int(int(id.(uint16)))
	case uint32:
		return a.Int(int(id.(uint32)))
	case uint64:
		return a.Int(int(id.(uint64)))
	case string:
		return a.String(id.(string))
	case []byte:
		return a.String(string(id.([]byte)))
	default:
		return a.String(fmt.Sprint(id))
	}
}
