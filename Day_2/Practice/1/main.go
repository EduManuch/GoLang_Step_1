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

import "fmt"

func main() {
	var num, sum int
	var statement string

	fmt.Println("Enter number: ")
	fmt.Scan(&num)

	for num != 0 {
		digit := num % 10

		if len(statement) == 0 {
			statement = fmt.Sprintf("%d =", digit)
		} else {
			statement = fmt.Sprintf("%d + ", digit) + statement
		}
		sum += digit
		num /= 10
	}
	fmt.Printf("Сумма числа: %v\n", sum)
	fmt.Printf("%s = %v\n", statement, sum)
}
