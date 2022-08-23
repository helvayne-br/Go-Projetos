package main

import (
	"fmt"
	"strconv"
)

func hello(name string) { println("Olá", name, ", tudo bem?") }
func sum(a, b int) int {
	fmt.Print("Sua idade atual é:")
	return a + b
}
func convertAndSum(a int, b string) (total int, err error) {
	i, err := strconv.Atoi(b)
	if err != nil {
		return
	}
	total = a + i

	return
}
func main() {
	hello("Hugão")
	fmt.Println(sum(4, 2))
	total, err := convertAndSum(10, "15")
	fmt.Println("convertendo string em inteiro mais soma, é igual a :", total, err)

}
