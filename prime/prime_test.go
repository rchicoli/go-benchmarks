package prime

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrimeNumbers(t *testing.T) {
	tests := []struct {
		name string
		max  int
		want []int
	}{
		{"primeOfOne", 5, []int{2, 33}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrimeNumbers(tt.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("primeNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkPrimeNumbers100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimeNumbers(100)
	}
}

var table = []struct {
	input int
}{
	{input: 100},
	{1000},
	{input: 200},
}

func BenchmarkPrimeNumbersInputs(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				PrimeNumbers(v.input)
			}
		})
	}
}

func BenchmarkSieveOfEratosthenes(b *testing.B) {
	// func BenchmarkPrimeNumbersInputs(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SieveOfEratosthenes(v.input)
			}
		})
	}
}
