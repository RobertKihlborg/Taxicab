package GoTaxicab

import (
	"testing"
)

func TestSmallPrimeFactorize(t *testing.T) {
	num := TA[4].Uint64()
	t.Log(SmallPrimeFactorize(num))
}

func TestNextFactorization(t *testing.T) {
	target := TA[3].Uint64()
	p, e := SmallPrimeFactorize(target)
	ce := make([]uint64, len(p))

	f1 := uint64(1)
	f2 := target

	continueLooping := true
	i := 0
	for continueLooping {
		//t.Logf("(f1, f2): (%v, %v), exponents: %v", f1, f2, ce)
		i++
		continueLooping = nextFactorization(&f1, &f2, &p, &e, &ce)
	}
	t.Logf("Total iterations: %v", i)
}

func BenchmarkSmallPrimeFactorize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SmallPrimeFactorize(TA[4].Uint64())
	}
}
