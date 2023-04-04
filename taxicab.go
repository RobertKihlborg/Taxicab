package taxicab

import (
	"fmt"
	"math/big"
	"strings"
)

type Int = big.Int

const (
	ta1  = "2"
	ta2  = "1_729"
	ta3  = "87_539_319"
	ta4  = "6_963_472_309_248"
	ta5  = "48_988_659_276_962_496"
	ta6  = "24_153_319_581_254_312_065_344"
	pta7 = "24_885_189_317_885_898_975_235_988_544"
)

var TA []*Int
var (
	One, Two, Three, Four, Five, Six, Seven, Twelve, Eighteen *Int
)

func init() {
	taStrings := []string{ta1, ta2, ta3, ta4, ta5, ta6, pta7}
	for i, s := range taStrings {
		s = strings.ReplaceAll(s, "_", "")
		taStrings[i] = s
		n, _ := new(Int).SetString(s, 10)
		TA = append(TA, n)
	}

	One = big.NewInt(1)
	Two = big.NewInt(2)
	Three = big.NewInt(3)
	Four = big.NewInt(4)
	Five = big.NewInt(5)
	Six = big.NewInt(6)
	Seven = big.NewInt(7)
	Twelve = big.NewInt(12)
	Eighteen = big.NewInt(18)

}

func main() {
	target := TA[5]
	fmt.Printf("Searching for taxicab solutions to %v\n", target)
	solutions := Partitions(target)
	fmt.Printf("--- %v hits: ---\n", len(solutions))
	for _, s := range solutions {
		fmt.Printf("[%v, %v]\n", s.a, s.b)
	}
}

type BigNumPair struct {
	a, b *Int
}

func GetStartX(y, target *Int) *Int {

	x := new(Int).Mod(new(Int).Sub(target, y), Six)
	discrepancy := CubeRootFloored(new(Int).Sub(target, Cube(y)))
	discrepancy.Sub(discrepancy, new(Int).Mod(discrepancy, Six))
	x.Add(x, discrepancy)
	return x
}

func Partitions(target *Int) []BigNumPair {
	var res []BigNumPair

	maxY := CubeRootFloored(target)
	minY := CubeRootFloored(new(Int).Rsh(target, 1))
	y := maxY
	x := GetStartX(y, target)
	//fmt.Printf("y range: %v, %v \n", minY, maxY)

	relativeSum := new(Int).Sub(new(Int).Add(Cube(x), Cube(y)), target)

	for y.Cmp(minY) >= 0 {
		//fmt.Printf("Testing x=%v and y=%v\n", x, y)
		switch relativeSum.Sign() {
		case 1:
			// (y-1)^3 - y^3 + (x+1)^3 - x^3    =   3(x^2 - y^2) + 3(x+y) = 3 (x^2 + x + y - y^2)
			change := new(Int).Mul(Three, Sum([]*Int{
				Square(x),
				x,
				y,
				SwapSign(Square(y))}))
			relativeSum.Add(relativeSum, change)
			y.Sub(y, One)
			x.Add(x, One)

		case -1:
			// (x+6)^3 - x^3    =   18*x^2 + 108x + 216		= 18 (x^2 + 6x + 12)
			change := new(Int).Mul(Eighteen,
				Sum([]*Int{
					Square(x),
					new(Int).Mul(x, Six),
					Twelve}))
			relativeSum.Add(relativeSum, change)
			x.Add(x, Six)

		case 0:
			//fmt.Printf("Success with %v, %v\n", x, y)
			res = append(res, BigNumPair{a: new(Int).Set(x), b: new(Int).Set(y)})

			// (y-1)^3 - y^3 + (x+7)^3 - x^3    =   3(7x^2 + 49x + y - y^2 + 114)
			change := new(Int).Mul(
				Three, Sum([]*Int{
					new(Int).Mul(Square(x), Seven),
					new(Int).Mul(x, big.NewInt(49)),
					y,
					SwapSign(Square(y)),
					big.NewInt(114)}))

			relativeSum.Add(relativeSum, change)

			y.Sub(y, One)
			x.Add(x, Seven)
		}
	}

	return res
}
