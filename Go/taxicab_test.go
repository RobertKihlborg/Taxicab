package Go

import "testing"

func TestPartitions(t *testing.T) {
	for i, n := range TA[:7] {
		t.Logf("Testing TA%v", i+1)
		solCount := len(Partitions(n))
		if solCount != i+1 {
			t.Fatalf("Found %v solutions for TA(%v), expected %v", solCount, n, i+1)
		}
	}
}
