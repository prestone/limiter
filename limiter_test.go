package limiter

//testing

import (
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/monopolly/console"
)

func TestNew(t *testing.T) {
	function, _, _, _ := runtime.Caller(0)
	fn := runtime.FuncForPC(function).Name()
	var log = console.New()
	log.OK(fmt.Sprintf("%s\n", fn[strings.LastIndex(fn, ".Test")+5:]))
	a := assert.New(t)
	_ = a

	r := New(time.Second*20, 5)
	a.True(r.Int(1))
	a.True(r.Int(1))
	a.True(r.Int(1))
	a.True(r.Int(1))
	a.True(r.Int(1))
	a.False(r.Int(1))
	a.False(r.Int(1))

	r = New(time.Second, 5)
	a.True(r.Int(1))
	a.True(r.Int(1))
	a.True(r.Int(1))
	a.True(r.Int(1))
	a.True(r.Int(1))
	a.False(r.Int(1))
	a.False(r.Int(1))

	time.Sleep(time.Second * 2)
	a.True(r.Int(1))
	a.True(r.Int(1))
	a.True(r.Int(1))
	a.True(r.Int(1))
	a.True(r.Int(1))
	a.False(r.Int(1))
	a.False(r.Int(1))

	log.OK("час")
	r = New(time.Hour, 10)
	//ok
	a.True(r.Int(1))
	a.True(r.Int(1))
	a.True(r.Int(1))
	a.True(r.Int(1))
	a.True(r.Int(1))

	time.Sleep(time.Second)
	//ok
	a.True(r.Int(1))
	a.True(r.Int(1))
	a.True(r.Int(1))
	a.True(r.Int(1))
	a.True(r.Int(1))

	//limit
	a.False(r.Int(1))
	a.False(r.Int(1))

}

func BenchmarkInt(b *testing.B) {
	r := New(time.Hour*20, 500000)
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r.Int(1)
	}
}

func BenchmarkIntRandom(b *testing.B) {
	r := New(time.Hour*20, 500000)
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r.Int(rand.Intn(10000))
	}
}

func BenchmarkString(b *testing.B) {
	r := New(time.Hour*20, 500000)
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r.String("192.168.0.1")
	}
}

func BenchmarkInterfaceUint(b *testing.B) {
	r := New(time.Hour*20, 500000)
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r.Interface(uint(425))
	}
}

func BenchmarkInterfaceString(b *testing.B) {
	r := New(time.Hour*20, 500000)
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r.Interface("nice toyec")
	}
}

func BenchmarkInterfaceBytes(b *testing.B) {
	r := New(time.Hour*20, 500000)
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r.Interface([]byte("nice"))
	}
}

func BenchmarkInterfaceSliceInt(b *testing.B) {
	r := New(time.Hour*20, 500000)
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r.Interface([]int{1, 2, 3, 4})
	}
}

func BenchmarkGetFreeParallel(b *testing.B) {
	r := New(time.Hour*20, 500000)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			r.Int(1)
		}
	})
}
