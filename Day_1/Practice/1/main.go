package main

import "fmt"

const minus = 1

func main() {

	fmt.Println("Введите имя и возраст")
	var (
		name string
		age  int
	)
	fmt.Scan(&name, &age)
	fmt.Printf("Your name is %s\nYour age is %d\n", name, age-minus)

}
