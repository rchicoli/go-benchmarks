package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var result interface{}

// https://github.com/Tkanos/strings-vs-bytes/blob/master/bench_test.go

// BenchmarkBytesToStrings-8               100000000               30.5 ns/op            32 B/op          1 allocs/op
// BenchmarkBytesToStringsWithSprintf-8    20000000               130 ns/op              64 B/op          2 allocs/op
// BenchmarkBytesCompareSame-8             300000000                4.39 ns/op            0 B/op          0 allocs/op
// BenchmarkBytesCompareDifferent-8        300000000                4.66 ns/op            0 B/op          0 allocs/op
// BenchmarkStringsCompareSame-8           2000000000               1.91 ns/op            0 B/op          0 allocs/op
// BenchmarkStringsCompareDifferent-8      200000000                8.99 ns/op            0 B/op          0 allocs/op
// BenchmarkBytesContains-8                100000000               11.5 ns/op             0 B/op          0 allocs/op
// BenchmarkStringsContains-8              100000000               10.6 ns/op             0 B/op          0 allocs/op
// BenchmarkBytesIndex-8                   200000000                7.17 ns/op            0 B/op          0 allocs/op
// BenchmarkStringIndex-8                  300000000                5.55 ns/op            0 B/op          0 allocs/op
// BenchmarkBytesReplace-8                 20000000                69.3 ns/op            32 B/op          1 allocs/op
// BenchmarkStringsReplace-8               20000000               109 ns/op              64 B/op          2 allocs/op
// BenchmarkBytesConcat2-8                 50000000                44.3 ns/op            32 B/op          1 allocs/op
// BenchmarkStringsConcat2-8               20000000                60.9 ns/op            32 B/op          1 allocs/op
// BenchmarkStringsJoin2-8                 30000000                50.4 ns/op            32 B/op          1 allocs/op
// BenchmarkMapHints-8                     200000000                6.96 ns/op            0 B/op          0 allocs/op
// BenchmarkMapsHints_Dont-8               100000000               12.3 ns/op             0 B/op          0 allocs/op

//BenchmarkBytesToStrings convert a []bytes in a string
//https://golang.org/pkg/bytes/#Compare
func BenchmarkBytesToStrings(b *testing.B) {

	s1 := []byte("string to convert")

	b.ResetTimer()
	b.ReportAllocs()
	var r string
	for n := 0; n < b.N; n++ {
		r = string(s1)
	}

	result = r
}

func BenchmarkBytesToStringsWithSprintf(b *testing.B) {

	s1 := []byte("string to convert")

	b.ResetTimer()
	b.ReportAllocs()
	var r string
	for n := 0; n < b.N; n++ {
		r = fmt.Sprintf("%s", s1)
	}

	result = r
}

//BenchmarkBytesCompareSame compare 2 []bytes.
//https://golang.org/pkg/bytes/#Compare
func BenchmarkBytesCompareSame(b *testing.B) {

	s1 := []byte("string to compare")
	s2 := []byte("string to compare")

	b.ResetTimer()
	b.ReportAllocs()
	var r int
	for n := 0; n < b.N; n++ {
		r = bytes.Compare(s1, s2)
	}

	result = r
}

//BenchmarkBytesCompareDifferent compare 2 []bytes.
//https://golang.org/pkg/bytes/#Compare
func BenchmarkBytesCompareDifferent(b *testing.B) {

	s1 := []byte("string to compare A")
	s2 := []byte("string to compare B")

	b.ResetTimer()
	b.ReportAllocs()
	var r int
	for n := 0; n < b.N; n++ {
		r = bytes.Compare(s1, s2)
	}

	result = r
}

//BenchmarkStringsCompareSame compare 2 strings.
//https://golang.org/pkg/strings/#Compare
func BenchmarkStringsCompareSame(b *testing.B) {

	s1 := "string to compare"
	s2 := "string to compare"

	b.ResetTimer()
	b.ReportAllocs()
	var r int
	for n := 0; n < b.N; n++ {
		r = strings.Compare(s1, s2)
	}

	result = r
}

//BenchmarkStringsCompareDifferent compare 2 strings.
//https://golang.org/pkg/strings/#Compare
func BenchmarkStringsCompareDifferent(b *testing.B) {

	s1 := "string to compare A"
	s2 := "string to compare B"

	b.ResetTimer()
	b.ReportAllocs()
	var r int
	for n := 0; n < b.N; n++ {
		r = strings.Compare(s1, s2)
	}

	result = r
}

//BenchmarkBytesContains check contains method
//https://golang.org/pkg/bytes/#Contains
func BenchmarkBytesContains(b *testing.B) {

	s1 := []byte("string to compare")
	s2 := []byte("comparc")

	b.ResetTimer()
	b.ReportAllocs()
	var r bool
	for n := 0; n < b.N; n++ {
		r = bytes.Contains(s1, s2)
	}

	result = r
}

//BenchmarkStringsContains check contains method
//https://golang.org/pkg/strings/#Contains
func BenchmarkStringsContains(b *testing.B) {

	s1 := "string to compare"
	s2 := "comparc"

	b.ResetTimer()
	b.ReportAllocs()
	var r bool
	for n := 0; n < b.N; n++ {
		r = strings.Contains(s1, s2)
	}

	result = r
}

//BenchmarkBytesIndex check contains index
//https://golang.org/pkg/bytes/#Index
func BenchmarkBytesIndex(b *testing.B) {

	s1 := []byte("string to compare")
	s2 := []byte("e")

	b.ResetTimer()
	b.ReportAllocs()
	var r int
	for n := 0; n < b.N; n++ {
		r = bytes.Index(s1, s2)
	}

	result = r
}

//BenchmarkStringIndex check contains index
//https://golang.org/pkg/strings/#Index
func BenchmarkStringIndex(b *testing.B) {

	s1 := "string to compare"
	s2 := "e"

	b.ResetTimer()
	b.ReportAllocs()
	var r int
	for n := 0; n < b.N; n++ {
		r = strings.Index(s1, s2)
	}

	result = r
}

//BenchmarkBytesReplace check replace method
//https://golang.org/pkg/bytes/#Replace
func BenchmarkBytesReplace(b *testing.B) {

	s1 := []byte("string to comparc")
	s2 := []byte("comparc")
	s3 := []byte("compare")

	b.ResetTimer()
	b.ReportAllocs()
	var r []byte
	for n := 0; n < b.N; n++ {
		r = bytes.Replace(s1, s2, s3, -1)
	}

	result = r
}

//BenchmarkStringsReplace check replace method
//https://golang.org/pkg/strings/#Replace
func BenchmarkStringsReplace(b *testing.B) {

	s1 := "string to comparc"
	s2 := "comparc"
	s3 := "compare"

	b.ResetTimer()
	b.ReportAllocs()
	var r string
	for n := 0; n < b.N; n++ {
		r = strings.Replace(s1, s2, s3, -1)
	}

	result = r
}

//BenchmarkBytesConcat concats 2 bytes
func BenchmarkBytesConcat2(b *testing.B) {

	s1 := []byte("string to compare")
	s2 := []byte("string to add")
	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		r := make([]byte, 0, len(s1)+len(s2))
		r = append(r, s1...)
		r = append(r, s2...)
	}
}

//BenchmarkStringsConcat concats 2 strings
func BenchmarkStringsConcat2(b *testing.B) {

	s1 := "string to compare"
	s2 := "string to add"
	b.ResetTimer()
	b.ReportAllocs()
	var r string
	for n := 0; n < b.N; n++ {
		r = s1 + s2
	}

	result = r
}

//BenchmarkStringsJoin joins 2 strings
//https://golang.org/pkg/strings/#Join
func BenchmarkStringsJoin2(b *testing.B) {

	s1 := "string to compare"
	s2 := "string to add"
	b.ResetTimer()
	b.ReportAllocs()
	var r string
	for n := 0; n < b.N; n++ {
		j := []string{s1, s2}
		r = strings.Join(j, "")
	}

	result = r
}

//BenchmarkMapHints bench mymap[string(abytes)]
func BenchmarkMapHints(b *testing.B) {
	mymap := make(map[string]string)
	mymap["key"] = "value"
	abytes := []byte("key")

	b.ResetTimer()
	b.ReportAllocs()

	var v string
	for n := 0; n < b.N; n++ {
		v, _ = mymap[string(abytes)]
	}

	result = v
}

//BenchmarkMapsHints_Dont bench key := string(abytes)
//v, _ = mymap[key]
func BenchmarkMapsHints_Dont(b *testing.B) {
	mymap := make(map[string]string)
	mymap["key"] = "value"
	abytes := []byte("key")

	b.ResetTimer()
	b.ReportAllocs()

	var v string
	for n := 0; n < b.N; n++ {
		key := string(abytes)
		v, _ = mymap[key]
	}

	result = v

}
