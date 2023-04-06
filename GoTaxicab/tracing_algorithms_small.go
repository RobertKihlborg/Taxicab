package GoTaxicab

func smallGetStartX(y, target uint64) uint64 {
	// x = (target - y) % 6
	x := (target - y) % 6

	discrepancy := SmallCubeRootFloored(target - SmallCube(y))
	discrepancy -= discrepancy % 6
	x += discrepancy
	return x

}

func SmallCubeSumAlgorithm(target uint64) [][]uint64 {
	var res [][]uint64

	maxY := SmallCubeRootFloored(target)
	minY := SmallCubeRootFloored(target >> 1)
	y := maxY
	x := smallGetStartX(y, target)
	//fmt.Printf("y range: %v, %v \n", minY, maxY)

	for y >= minY {
		cubeSum := SmallCubeSum(x, y)
		//fmt.Printf("Testing x=%v and y=%v\n", x, y)
		if cubeSum > target {
			x++
			y--
		} else if cubeSum < target {
			x += 6
		} else {
			res = append(res, []uint64{x, y})
			x += 7
			y--
		}
	}

	return res
}
