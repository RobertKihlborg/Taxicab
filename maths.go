package taxicab

import (
	"math/big"
)

func Square(x *Int) *Int {
	return new(Int).Mul(x, x)
}

func Cube(x *Int) *Int {
	return new(Int).Mul(x, new(Int).Mul(x, x))
}

func SwapSign(z *Int) *Int {
	z.Neg(z)
	return z
}

func Sum(x []*Int) *Int {
	a := new(Int)

	for _, n := range x {
		a.Add(a, n)
	}
	return a
}

func CubeRootFloored(x *Int) *Int {
	if len(x.Bits()) == 0 {
		return new(Int)
	}

	nLog2Floored := uint(x.BitLen() - 1)
	lowGuessLog2 := nLog2Floored / 3
	highGuessLog2 := lowGuessLog2 + 1

	lowGuess := new(Int).Lsh(big.NewInt(1), lowGuessLog2)
	highGuess := new(Int).Lsh(One, highGuessLog2)

	one := big.NewInt(1)
	// highGuess > lowGuess + 1
	for highGuess.Cmp(new(Int).Add(lowGuess, one)) == 1 {
		// newGuess = (highGuess + lowGuess) / 2
		newGuess := new(Int).Add(lowGuess, highGuess)
		newGuess.Rsh(newGuess, 1)

		if Cube(newGuess).Cmp(x) == 1 {
			highGuess = newGuess
		} else {
			lowGuess = newGuess
		}
	}

	return lowGuess
}
