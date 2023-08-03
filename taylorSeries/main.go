package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	var n int

	fmt.Print("Введите значение x: ")
	fmt.Scan(&x)

	fmt.Print("Введите количество знаков после запятой: ")
	fmt.Scan(&n)

	epsilon := 1 / math.Pow(10, float64(n))

	sum := 1.0

	term := 1.0
	i := 1
	for math.Abs(term) >= epsilon {
		term *= x / float64(i)
		sum += term
		i++
	}

	fmt.Println("Результат:", sum)
}
