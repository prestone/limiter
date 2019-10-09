package limiter

//testing

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	a := assert.New(t)
	_ = a

	r := New(5, time.Second*20)
	a.True(r.Int(1))
	a.True(r.Int(1))
	a.True(r.Int(1))
	a.True(r.Int(1))
	a.True(r.Int(1))
	a.False(r.Int(1))
	a.False(r.Int(1))

	r = New(5, time.Second)
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

	r = New(10, time.Hour)
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
	r := New(5000, time.Hour)
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r.Int(1)
	}
}

func BenchmarkIntRandom(b *testing.B) {
	r := New(5000, time.Hour)
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r.Int(rand.Intn(10000))
	}
}

func BenchmarkString(b *testing.B) {
	r := New(5000, time.Hour)
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r.String("192.168.0.1")
	}
}

func BenchmarkInterfaceUint(b *testing.B) {
	r := New(5000, time.Hour)
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r.Interface(uint(425))
	}
}

func BenchmarkInterfaceString(b *testing.B) {
	r := New(5000, time.Hour)
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r.Interface("nice toyec")
	}
}

func BenchmarkInterfaceBytes(b *testing.B) {
	r := New(5000, time.Hour)
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r.Interface([]byte("nice"))
	}
}

func BenchmarkInterfaceSliceInt(b *testing.B) {
	r := New(5000, time.Hour)
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r.Interface([]int{1, 2, 3, 4})
	}
}

func BenchmarkGetFreeParallel(b *testing.B) {
	r := New(5000, time.Hour)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			r.Int(1)
		}
	})
}
