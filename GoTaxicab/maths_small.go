package GoTaxicab

import (
	"math/bits"
)

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
