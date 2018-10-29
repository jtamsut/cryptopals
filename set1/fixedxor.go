package main

import (
	"math/big"
)

func fixedXOR(a, b string) int64 {

	l := big.NewInt(0)
	m := big.NewInt(0)

	l.SetString(a, 16)
	m.SetString(b, 16)

	//need to XOR and return

}

func main() {
	a := "1c0111001f010100061a024b53535009181c"
	b := "686974207468652062756c6c277320657965"
	fixedXOR(a, b)
	// fmt.Printf("%b", fixedXOR(a, b))
}
