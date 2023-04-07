package GoTaxicab

import (
	"testing"
)

func TestSmallPrimeFactorize(t *testing.T) {
	for i := uint64(0); i < 20; i++ {
		p, e := SmallPrimeFactorize(i)
		t.Logf("%v, %v, %v", i, p, e)
	}
}
