/*
Задача №1
Вводим любое натуральное число.
Нужно посчитать сумму цифр числа, с помощью цикла for


Пример:
In: 4521
Out: 12
*Задание 1.1: 4 + 5 + 2 + 1 = 12 - добавить к выводу сумму как выражение
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var num, sum int
	var statement string

	fmt.Println("Enter number: ")
	fmt.Scan(&num)

	for num != 0 {
		digit := num % 10
		statement = fmt.Sprintf("%d + ", digit) + statement
		sum += digit
		num /= 10
	}

	statement, _ = strings.CutSuffix(statement, " + ")
	fmt.Println(statement, "=", sum)
}
