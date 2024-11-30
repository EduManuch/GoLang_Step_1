/*
Задача № 3. Вывести на экран в порядке возрастания три введенных числа
Пример:
Вход: 1 9 2
Выход: 1 2 9

Про sort мы пока ничего не знаем!
Это задача на условный оператор
*/

package main

import "fmt"

func main() {
	var a, b, c uint16

	fmt.Print("Enter three numbers: ")
	fmt.Scan(&a, &b, &c)

	if a < b && a < c {
		if b < c {
			fmt.Println(a, b, c)
		} else {
			fmt.Println(a, c, b)
		}
	}
	if b < a && b < c {
		if a < c {
			fmt.Println(b, a, c)
		} else {
			fmt.Println(b, c, a)
		}
	}
	if c < a && c < b {
		if a < b {
			fmt.Println(c, a, b)
		} else {
			fmt.Println(c, b, a)
		}
	}

}
