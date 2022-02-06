package reverse

import (
	"strings"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{"ReverseFoobar", "foobar", "raboof"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.str); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestReverseUnicode(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{"ReverseFoobar", "foobar", "raboof"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseUnicode(tt.str); got != tt.want {
				t.Errorf("reverseUnicode() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseUnicode2(tt.str); got != tt.want {
				t.Errorf("reverseUnicode2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseRange(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{"ReverseFoobar", "foobar", "raboof"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseRange(tt.str); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkReverse(b *testing.B) {
	str := "foobar"
	var rev string
	for i := 0; i < b.N; i++ {
		rev += Reverse(str)
	}
	if s := strings.Repeat("raboof", b.N); rev != s {
		b.Errorf("reverse() = %v, want %v", rev, s)
	}
}

func BenchmarkReverseUnicode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseUnicode("foobar")
	}
}

func BenchmarkReverseUnicode2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseUnicode2("foobar")
	}
}

func BenchmarkReverseRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse("foobar")
	}
}
