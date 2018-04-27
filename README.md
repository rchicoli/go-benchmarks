# Go Benchmarks

How to run:

```bash
go test -v -bench=. -benchtime 1s -benchmem
```

Output:

```bash
BenchmarkConcat-8                1000000             39818 ns/op          503992 B/op          1 allocs/op
BenchmarkBuffer-8               200000000                6.28 ns/op            2 B/op          0 allocs/op
BenchmarkCopy-8                 500000000                3.04 ns/op            0 B/op          0 allocs/op
BenchmarkStringBuilder-8        200000000                7.73 ns/op            5 B/op          0 allocs/op
```