package GoTaxicab

import "testing"

func TestBigRelativeSumAlgorithm(t *testing.T) {
	for i, n := range TA[:4] {
		t.Logf("Testing TA%v", i+1)
		solCount := len(BigRelativeSumAlgorithm(n))
		if solCount != i+1 {
			t.Fatalf("Found %v solutions for TA(%v), expected %v", solCount, n, i+1)
		}
	}
}

func TestSmallCubeSumAlgorithm(t *testing.T) {
	for i, n := range TA[:4] {
		t.Logf("Testing TA%v", i+1)
		solCount := len(SmallCubeSumAlgorithm(n.Uint64()))
		if solCount != i+1 {
			t.Fatalf("Found %v solutions for TA(%v), expected %v", solCount, n, i+1)
		}
	}
}

const uint64BenchTarget = 4

func BenchmarkBigRelativeSumAlgorithm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BigRelativeSumAlgorithm(TA[uint64BenchTarget])
	}
}

func BenchmarkSmallCubeSumAlgorithm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SmallCubeSumAlgorithm(TA[uint64BenchTarget].Uint64())
	}
}
