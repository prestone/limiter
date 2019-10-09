# 💂Limiter
Limit any request is simple now. Sometimes you want to limit api requests or any data. So this is a simple solution for limits. I use it in production.

### Install
```go get -u github.com/prestone/limiter```

### Example
```go
// ok, lets create a new limiter
// with 10 request per minute max
// its a memory struct
limit := limiter.New(10, time.Second)

// now we can use it for checking strin, int and any ID
// for example you want limit your api request
// you can make it simple by ip
ok := limit.String("192.168.0.1")

// or by userID
ok := limit.Int(42)

// or any you want
ok := limit.Interface([]byte("nice")))

//its checks period and count request
// so answer is just bool
// is it ok or limit
if !ok{
  //limit here
} else {
  //it ok now
}
  
//its a simple limiter help me in my projects
//and here is a benchmarks
BenchmarkInt-8                   7979979               154 ns/op               0 B/op          0 allocs/op
BenchmarkIntRandom-8             4268584               269 ns/op               0 B/op          0 allocs/op
BenchmarkString-8                6635058               191 ns/op               0 B/op          0 allocs/op
BenchmarkInterfaceUint-8         7213478               175 ns/op               0 B/op          0 allocs/op
BenchmarkInterfaceString-8       5515800               217 ns/op               0 B/op          0 allocs/op
BenchmarkInterfaceBytes-8        4535649               258 ns/op              36 B/op          2 allocs/op
BenchmarkInterfaceSliceInt-8     1251056               954 ns/op             112 B/op          7 allocs/op
BenchmarkGetFreeParallel-8       3537357               336 ns/op

```

## Thanks
I hope it help you in your an awesome projects.
