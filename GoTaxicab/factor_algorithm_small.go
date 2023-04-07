package GoTaxicab

func nextFactorization(f1, f2 *uint64, primes, exponents, currentExponent *[]uint64) bool {
	facCount := len(*primes)

	for ixToInc := 0; ixToInc < facCount; ixToInc++ {
		p := (*primes)[ixToInc]
		// If exponent can be increased
		if (*currentExponent)[ixToInc] < (*exponents)[ixToInc] {
			*f1 *= p
			*f2 /= p
			(*currentExponent)[ixToInc]++
			return true
		}

		// If exponent increasing causes overflow
		totalChangeFactor := SmallPow(p, (*currentExponent)[ixToInc])
		*f1 /= totalChangeFactor
		*f2 *= totalChangeFactor

		(*currentExponent)[ixToInc] = 0

	}

	return false
}

func SmallFactorAlgorithm(target uint64) [][]uint64 {
	var res [][]uint64

	minf1 := SmallCubeRootFloored(target)
	maxf1 := 2*SmallCubeRootFloored(target>>1) + 1

	primes, exponents := SmallPrimeFactorize(target)
	currentExponents := make([]uint64, len(exponents))

	f1 := uint64(1)
	f2 := target
	targetMod6 := target % 6

	continueLooping := true
	for continueLooping {
		if f1%6 == targetMod6 && minf1 <= f1 && f1 <= maxf1 {
			sqrk := (4*f2 - f1*f1) / 3
			k := SmallSquareRootFloored(sqrk)
			if k*k == sqrk && k%2 == f1%2 {
				res = append(res, []uint64{(f1 - k) / 2, (f1 + k) / 2})
			}
		}
		continueLooping = nextFactorization(&f1, &f2, &primes, &exponents, &currentExponents)
	}

	return res
}
