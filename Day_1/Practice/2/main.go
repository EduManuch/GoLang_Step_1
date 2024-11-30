package main

import "fmt"

func main() {
	var n uint8
	fmt.Scan(&n)

	fmt.Println("Первая цифра", n/100)
	fmt.Println("Вторая цифра", n/10%10)
	fmt.Println("Третья цифра", n%10)
}
