package main

import (
	"fmt"
)

func main() {

	a, b := 15, 15

	switch {
	case a > b:
		fmt.Println("A é mais velho que B.")
	case a < b:
		fmt.Println("A é mais novo que B.")

	default:
		fmt.Println("A e B tem a mesma idade.")

	}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	nomes := []string{"Paulo", "Priscila", "Hugo", "Pandora"}

	for i := 0; i < len(nomes); i++ {
		fmt.Println(nomes[i])
	}

}
