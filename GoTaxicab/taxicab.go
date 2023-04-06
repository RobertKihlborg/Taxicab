package GoTaxicab

import (
	"log"
	"math/big"
	"os"
	"strings"
)

type Int = big.Int

var TA []*Int

func init() {
	readTaxicabFiles()
}

func readTaxicabFiles() {
	taFileBytes, err := os.ReadFile("../Taxicabs.txt")
	if err != nil {
		log.Fatal(err)
	}
	taFileString := string(taFileBytes)
	taStrings := strings.Fields(taFileString)

	for i, s := range taStrings {
		taStrings[i] = s
		n, success := new(Int).SetString(s, 10)
		if !success {
			log.Fatalf("Failiure to read string %v", s)
		}
		TA = append(TA, n)
	}

}

type BigNumPair struct {
	a, b *Int
}
