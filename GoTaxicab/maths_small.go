package GoTaxicab

import (
	"math/bits"
)

func SmallPow(x, y uint64) uint64 {
	res := uint64(1)
	for i := uint64(0); i < y; i++ {
		res *= x
	}
	return res
}

func SmallCube(x uint64) uint64 {
	return x * x * x
}

func SmallSquare(x uint64) uint64 {
	return x * x
}

func SmallCubeSum(x, y uint64) uint64 {
	return SmallCube(x) + SmallCube(y)
}

func SmallCubeRootFloored(x uint64) uint64 {
	if x == 0 {
		return 0
	}

	xLog2Floored := uint64(bits.Len64(x) - 1)
	lowGuessLog2 := xLog2Floored / 3
	highGuessLog2 := lowGuessLog2 + 1

	lowGuess := uint64(1 << lowGuessLog2)
	highGuess := uint64(1 << highGuessLog2)

	for highGuess > lowGuess+1 {
		newGuess := (highGuess + lowGuess) / 2

		if SmallCube(newGuess) > x {
			highGuess = newGuess
		} else {
			lowGuess = newGuess
		}
	}
	return lowGuess
}

func SmallSquareRootFloored(x uint64) uint64 {
	if x == 0 {
		return 0
	}

	xLog2Floored := uint64(bits.Len64(x) - 1)
	lowGuessLog2 := xLog2Floored / 2
	highGuessLog2 := lowGuessLog2 + 1

	lowGuess := uint64(1 << lowGuessLog2)
	highGuess := uint64(1 << highGuessLog2)

	for highGuess > lowGuess+1 {
		newGuess := (highGuess + lowGuess) / 2

		if SmallSquare(newGuess) > x {
			highGuess = newGuess
		} else {
			lowGuess = newGuess
		}
	}
	return lowGuess
}

// SmallPrimeFactorize takes a number x and returns its prime factorization as two lists.
// The first list contains the primes, and the second list contains the exponents.
// For example, SmallPrimeFactorize(12) = [2, 3], [2, 1] since 12 = 2^2*3^1
// x = 1 -> [],[]
// x = 0 -> nil, nil
func SmallPrimeFactorize(x uint64) ([]uint64, []uint64) {
	if x == 0 {
		return nil, nil
	}
	if x == 1 {
		return []uint64{}, []uint64{}
	}

	var primes []uint64
	var exponents []uint64

	factor := uint64(2)
	for SmallSquare(factor) <= x {
		if x%factor == 0 {
			// At least one factor exists
			x = x / factor
			exp := uint64(1)
			// Check how many of that factor there are
			for x%factor == 0 {
				x = x / factor
				exp++
			}
			primes = append(primes, factor)
			exponents = append(exponents, exp)
		}
		factor++
	}

	if x != 1 {
		primes = append(primes, x)
		exponents = append(exponents, 1)
	}

	return primes, exponents
}
