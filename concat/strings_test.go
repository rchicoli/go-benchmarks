package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// https://www.soroushjp.com/2015/01/27/beautifully-simple-benchmarking-with-go/
// https://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go

// BenchmarkConcat-8                1000000             39818 ns/op          503992 B/op          1 allocs/op
// BenchmarkBuffer-8               200000000                6.28 ns/op            2 B/op          0 allocs/op
// BenchmarkCopy-8                 500000000                3.04 ns/op            0 B/op          0 allocs/op
// BenchmarkStringBuilder-8        200000000                7.73 ns/op            5 B/op          0 allocs/op

func BenchmarkConcat(b *testing.B) {
	// b.Skip()
	var str string
	for n := 0; n < b.N; n++ {
		str += "x"
	}
	// b.StopTimer()

	if s := strings.Repeat("x", b.N); str != s {
		b.Errorf("unexpected result; got=%s, want=%s", str, s)
	}
}

func BenchmarkFmt(b *testing.B) {
	var str string
	for n := 0; n < b.N; n++ {
		str = fmt.Sprintf("%s%s", str, "x")
	}
	if s := strings.Repeat("x", b.N); str != s {
		b.Errorf("unexpected result; got=%s, want=%s", str, s)
	}
}

func BenchmarkBuffer(b *testing.B) {

	var buffer bytes.Buffer
	for n := 0; n < b.N; n++ {
		buffer.WriteString("x")
	}
	// b.StopTimer()

	if s := strings.Repeat("x", b.N); buffer.String() != s {
		b.Errorf("unexpected result; got=%s, want=%s", buffer.String(), s)
	}
}

func BenchmarkCopy(b *testing.B) {

	bs := make([]byte, b.N)
	bl := 0
	for n := 0; n < b.N; n++ {
		// bl += copy(bs[bl:], "x")
		bl += copy(bs[bl:], "x")
	}

	if s := strings.Repeat("x", b.N); string(bs) != s {
		b.Errorf("unexpected result; got=%s, want=%s", string(bs), s)
	}
}

// Go 1.10
func BenchmarkStringBuilder(b *testing.B) {
	var strBuilder strings.Builder

	// b.ResetTimer()
	for n := 0; n < b.N; n++ {
		strBuilder.WriteString("x")
	}
	// b.StopTimer()

	if s := strings.Repeat("x", b.N); strBuilder.String() != s {
		b.Errorf("unexpected result; got=%s, want=%s", strBuilder.String(), s)
	}
}
