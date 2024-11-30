/*
Задача № 2. Получить реверсную запись трехзначного числа
Пример:
вход: 346, выход: 643
вход: 120, выход: 021
вход: 100, выход: 001
*/

package main

import "fmt"

func main() {

	var n, a, b, c uint16

	fmt.Print("Enter your number: ")
	fmt.Scan(&n)

	if n < 100 || n > 999 {
		fmt.Println("Incorrect number")
		return
	}

	a = n / 100
	b = n / 10 % 10
	c = n % 10
	new_n := c*100 + b*10 + a
	fmt.Println(new_n)
}
