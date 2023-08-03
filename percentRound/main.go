package main

import (
	"fmt"
	"math"
)

func main() {

	var deposit float64
	fmt.Print("Введите сумму вклада: ")
	fmt.Scan(&deposit)

	var percent int
	fmt.Print("Введите ежемесячный процент: ")
	fmt.Scan(&percent)

	var termYear int
	fmt.Print("Введите срок вклада (кол. лет): ")
	fmt.Scan(&termYear)

	termMonths := termYear * 12
	var bankRest float64

	for i := 0; i < termMonths; i++ {
		deposit += deposit * float64(percent) / 100
		fmt.Println("Депозит на круге", i, "равен", deposit)

		rest := deposit - float64(math.Floor(deposit*100))/100
		fmt.Println("Остаток на круге", i, "равен", rest)

		deposit -= rest
		bankRest += rest
	}

	fmt.Println("Сумма вклада на конец срока:", deposit)
	fmt.Println("Банковский остаток за счёт окргуления копеек:", bankRest)
}
