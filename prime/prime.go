package prime

import "math"

func PrimeNumbers(max int) []int {
	var primes []int

	for i := 2; i < max; i++ {
		isPrime := true

		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, i)
		}
	}

	return primes
}

func SieveOfEratosthenes(max int) []int {
	b := make([]bool, max)

	var primes []int

	for i := 2; i < max; i++ {
		if b[i] {
			continue
		}

		primes = append(primes, i)

		for k := i * i; k < max; k += i {
			b[k] = true
		}
	}

	return primes
}
