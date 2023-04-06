package GoTaxicab

func getStartX(y, target *Int) *Int {
	// x = (target - y) % 6
	x := new(Int).Sub(target, y)
	x.Mod(x, c6)

	discrepancy := CubeRootFloored(new(Int).Sub(target, Cube(y)))
	discrepancy.Sub(discrepancy, new(Int).Mod(discrepancy, c6))
	x.Add(x, discrepancy)
	return x
}

func RelativeSumAlgorithm(target *Int) []BigNumPair {
	var res []BigNumPair

	maxY := CubeRootFloored(target)
	minY := CubeRootFloored(new(Int).Rsh(target, 1))
	y := maxY
	x := getStartX(y, target)
	//fmt.Printf("y range: %v, %v \n", minY, maxY)

	relativeSum := new(Int).Sub(new(Int).Add(Cube(x), Cube(y)), target)

	for y.Cmp(minY) >= 0 {
		//fmt.Printf("Testing x=%v and y=%v\n", x, y)
		switch relativeSum.Sign() {
		case 1:
			// (y-1)^3 - y^3 + (x+1)^3 - x^3    =   3(x^2 - y^2) + 3(x+y) = 3 (x^2 + x + y - y^2)
			change := new(Int).Mul(c3, Sum([]*Int{
				Square(x),
				x,
				y,
				SwapSign(Square(y))}))
			relativeSum.Add(relativeSum, change)
			y.Sub(y, c1)
			x.Add(x, c1)

		case -1:
			// (x+6)^3 - x^3    =   18*x^2 + 108x + 216		= 18 (x^2 + 6x + 12)
			change := new(Int).Mul(c18,
				Sum([]*Int{
					Square(x),
					new(Int).Mul(x, c6),
					c12}))
			relativeSum.Add(relativeSum, change)
			x.Add(x, c6)

		case 0:
			//fmt.Printf("Success with %v, %v\n", x, y)
			res = append(res, BigNumPair{a: new(Int).Set(x), b: new(Int).Set(y)})

			// (y-1)^3 - y^3 + (x+7)^3 - x^3    =   3(7x^2 + 49x + y - y^2 + 114)
			change := new(Int).Mul(
				c3, Sum([]*Int{
					new(Int).Mul(Square(x), c7),
					new(Int).Mul(x, c49),
					y,
					SwapSign(Square(y)),
					c114}))

			relativeSum.Add(relativeSum, change)

			y.Sub(y, c1)
			x.Add(x, c7)
		}
	}

	return res
}
