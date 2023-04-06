package GoTaxicab

import (
	"math/big"
)

var (
	c1, c2, c3, c4, c5, c6, c7, c8, c12, c18, c49, c114 *Int
)

func init() {

	c1 = big.NewInt(1)
	c2 = big.NewInt(2)
	c3 = big.NewInt(3)
	c4 = big.NewInt(4)
	c5 = big.NewInt(5)
	c6 = big.NewInt(6)
	c7 = big.NewInt(7)
	c8 = big.NewInt(7)
	c12 = big.NewInt(12)
	c18 = big.NewInt(18)
	c49 = big.NewInt(49)
	c114 = big.NewInt(114)
}

func Square(x *Int) *Int {
	return new(Int).Mul(x, x)
}

func SetToSquare(z, x *Int) *Int {
	return z.Exp(x, c2, nil)
}

func Cube(x *Int) *Int {
	return new(Int).Mul(x, new(Int).Mul(x, x))
}

func SetToCube(z, x *Int) *Int {
	return z.Exp(x, c3, nil)
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
	if x.BitLen() == 0 {
		return new(Int)
	}

	nLog2Floored := uint(x.BitLen() - 1)
	lowGuessLog2 := nLog2Floored / 3
	highGuessLog2 := lowGuessLog2 + 1

	lowGuess := new(Int).Lsh(c1, lowGuessLog2)
	highGuess := new(Int).Lsh(c1, highGuessLog2)

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
