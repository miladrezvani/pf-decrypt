package main

import (
	"fmt"
	"strings"

	"github.com/miladrezvani/playfairCipher"
)

func main() {
	var keyword = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var word string
	fmt.Print("inpute ciphertext :")
	// word = "UGLPLPGYFQYTFQTOHFUGGYBAMUBMLUUGKRAOHBAWQUAFOCUDZY"
	fmt.Scanln(&word)
	i := 0
	perm := Permutation(1,keyword)
			var result string
		for {
			pwd := perm()
				if len(pwd) == 0 {
					i++
					if i > 25 {
						break
					}
					perm = Permutation(i, keyword)
					continue
				}
			fmt.Println(pwd)
			pf := playfairCipher.NewMatrix(pwd, word)
			pf.CreateTable()
			fmt.Println(pf.Matrix)
			result = pf.Decrypt()
			test := strings.Index(result, "CRYPTOGRAPHY")
			if test != -1 {
				fmt.Println(pwd)
				break
			}
		}
		fmt.Println(result)
}


 func Permutation(n int, c string) func() string {
    r := []rune(c)
    p := make([]rune, n)
    x := make([]int, len(p))
    return func() string {
        p := p[:len(x)]
        for i, xi := range x {
            p[i] = r[xi]
        }
        for i := len(x) - 1; i >= 0; i-- {
            x[i]++
            if x[i] < len(r) {
                break
            }
            x[i] = 0
            if i <= 0 {
                x = x[0:0]
                break
            }
        }
        return string(p)
    }
}