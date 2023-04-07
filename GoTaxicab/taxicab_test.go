package GoTaxicab

import (
	"testing"
)

const BigTestLimit = 6
const SmallTestLimit = 5

func TestBigRelativeSumAlgorithm(t *testing.T) {
	for i, n := range TA[:BigTestLimit] {
		t.Logf("Testing TA%v", i+1)
		solCount := len(BigRelativeSumAlgorithm(n))
		if solCount != i+1 {
			t.Fatalf("Found %v solutions for TA(%v), expected %v", solCount, n, i+1)
		}
	}
}

func TestBigCubeSumAlgorithm(t *testing.T) {
	for i, n := range TA[:BigTestLimit] {
		t.Logf("Testing TA%v", i+1)
		solCount := len(BigCubeSumAlgorithm(n))
		if solCount != i+1 {
			t.Fatalf("Found %v solutions for TA(%v), expected %v", solCount, n, i+1)
		}
	}
}

func TestSmallCubeSumAlgorithm(t *testing.T) {
	for i, n := range TA[:SmallTestLimit] {
		t.Logf("Testing TA%v", i+1)
		solutions := SmallCubeSumAlgorithm(n.Uint64())
		//t.Log(solutions)
		solCount := len(solutions)
		if solCount != i+1 {
			t.Fatalf("Found %v solutions for TA(%v), expected %v", solCount, n, i+1)
		}
	}
}

func TestSmallFactorAlgorithm(t *testing.T) {
	for i, n := range TA[:SmallTestLimit] {
		t.Logf("Testing TA%v", i+1)
		solutions := SmallFactorAlgorithm(n.Uint64())
		//t.Log(solutions)
		solCount := len(solutions)
		if solCount != i+1 {
			t.Fatalf("Found %v solutions for TA(%v), expected %v", solCount, n, i+1)
		}
	}
}

const uint64BenchTarget = 4 // <5

func BenchmarkBigRelativeSumAlgorithm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BigRelativeSumAlgorithm(TA[uint64BenchTarget])
	}
}

func BenchmarkBigCubeSumAlgorithm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BigCubeSumAlgorithm(TA[uint64BenchTarget])
	}
}

func BenchmarkSmallCubeSumAlgorithm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SmallCubeSumAlgorithm(TA[uint64BenchTarget].Uint64())
	}
}

func BenchmarkSmallFactorAlgorithm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SmallFactorAlgorithm(TA[uint64BenchTarget].Uint64())
	}
}
