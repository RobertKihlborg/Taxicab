package GoTaxicab

import "testing"

func TestRelativeSumAlgorithm(t *testing.T) {
	for i, n := range TA[:4] {
		t.Logf("Testing TA%v", i+1)
		solCount := len(RelativeSumAlgorithm(n))
		if solCount != i+1 {
			t.Fatalf("Found %v solutions for TA(%v), expected %v", solCount, n, i+1)
		}
	}
}
