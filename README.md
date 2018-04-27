# Go Benchmarks

This is a collection of multiple benchmarks. Some of them are just copied from other projects or from the internet.

How to run:

```bash
go test -v -bench=. -benchtime 1s -benchmem ./...
```

Benchmark Output:

```bash
# Concat Strings
BenchmarkConcat-8                1000000             39818 ns/op          503992 B/op          1 allocs/op
BenchmarkBuffer-8               200000000                6.28 ns/op            2 B/op          0 allocs/op
BenchmarkCopy-8                 500000000                3.04 ns/op            0 B/op          0 allocs/op
BenchmarkStringBuilder-8        200000000                7.73 ns/op            5 B/op          0 allocs/op

# Bytes vs Strings
BenchmarkBytesToStrings-8               100000000               30.5 ns/op            32 B/op          1 allocs/op
BenchmarkBytesToStringsWithSprintf-8    20000000               130 ns/op              64 B/op          2 allocs/op
BenchmarkBytesCompareSame-8             300000000                4.39 ns/op            0 B/op          0 allocs/op
BenchmarkBytesCompareDifferent-8        300000000                4.66 ns/op            0 B/op          0 allocs/op
BenchmarkStringsCompareSame-8           2000000000               1.91 ns/op            0 B/op          0 allocs/op
BenchmarkStringsCompareDifferent-8      200000000                8.99 ns/op            0 B/op          0 allocs/op
BenchmarkBytesContains-8                100000000               11.5 ns/op             0 B/op          0 allocs/op
BenchmarkStringsContains-8              100000000               10.6 ns/op             0 B/op          0 allocs/op
BenchmarkBytesIndex-8                   200000000                7.17 ns/op            0 B/op          0 allocs/op
BenchmarkStringIndex-8                  300000000                5.55 ns/op            0 B/op          0 allocs/op
BenchmarkBytesReplace-8                 20000000                69.3 ns/op            32 B/op          1 allocs/op
BenchmarkStringsReplace-8               20000000               109 ns/op              64 B/op          2 allocs/op
BenchmarkBytesConcat2-8                 50000000                44.3 ns/op            32 B/op          1 allocs/op
BenchmarkStringsConcat2-8               20000000                60.9 ns/op            32 B/op          1 allocs/op
BenchmarkStringsJoin2-8                 30000000                50.4 ns/op            32 B/op          1 allocs/op
BenchmarkMapHints-8                     200000000                6.96 ns/op            0 B/op          0 allocs/op
BenchmarkMapsHints_Dont-8               100000000               12.3 ns/op             0 B/op          0 allocs/op

```

## Sources

  - https://github.com/Tkanos/strings-vs-bytes/blob/master/bench_test.go
  - https://www.soroushjp.com/2015/01/27/beautifully-simple-benchmarking-with-go/
  - https://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go