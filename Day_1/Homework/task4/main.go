/*
Задача № 4. Проверить, является ли четырехзначное число палиндромом
Пример:
Вход: 1221  Выход: 1221 - палиндром
Вход: 1234  Выход: 1234 - не палиндром
*/

package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter your naumber: ")
	fmt.Scan(&n)

	if n/1000 == 0 {
		fmt.Println("Incorrect number")
		return
	}
	a, b, c, d := n/1000, n%1000/100, n%100/10, n%10
	if a == d && b == c {
		fmt.Println("Palindrom")
	} else {
		fmt.Println("Not Palindrom")
	}
}
