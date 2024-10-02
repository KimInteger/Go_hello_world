package find

func FindPrimesInRange(start, end int) []int {
	if start < 2 {
		start = 2
	}

	isPrime := make([]bool, end+1)
	for i := range isPrime {
		isPrime[i] = true
	}

	for i := 2; i*i <= end; i++ {
		if isPrime[i] {
			for j := i * i; j <= end; j += i {
				isPrime[j] = false
			}
		}
	}

	primes := []int{}
	for i := start; i <= end; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes
}
