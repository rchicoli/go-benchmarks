package count

import (
	"fmt"
	"testing"
)

var tests = []struct {
	name string
	str  string
	want int
}{
	{"ThisAndThat", "this and that", 3},
	{"ThisAndThatWithTab", "	this and that", 3},
	{"ThisAndThatWhiteSpaces", " this and that ", 3},
}

func TestWordCount(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s/WordCountRegexp", tt.name), func(t *testing.T) {
			if got := WordCountRegexp(tt.str); got != tt.want {
				t.Errorf("WordCountRegexp() = %v, want %v", got, tt.want)
			}
		})
		t.Run(fmt.Sprintf("%s/WordCount", tt.name), func(t *testing.T) {
			if got := WordCount(tt.str); got != tt.want {
				t.Errorf("WordCount() = %v, want %v", got, tt.want)
			}
		})
		t.Run(fmt.Sprintf("%s/WordCountUnicode", tt.name), func(t *testing.T) {
			if got := WordCountUnicode(tt.str); got != tt.want {
				t.Errorf("WordCountUnicode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCountRegexp(b *testing.B) {
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				WordCountRegexp(tt.str)
			}
		})
	}
}

func BenchmarkCount(b *testing.B) {
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				WordCount(tt.str)
			}
		})
	}
}

func BenchmarkCountUnicode(b *testing.B) {
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				WordCountUnicode(tt.str)
			}
		})
	}
}
