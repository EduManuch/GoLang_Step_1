package main

import "fmt"

func main() {
	//Декларирование и инициализация пользовательским значением
	var age int
	fmt.Println("My age is:", age)
	var height int = 183
	fmt.Println("My height is:", height)

	//В чем "полустрогость" типизации?
	var uid = 12345
	fmt.Println("My uid:", uid)
	fmt.Printf("%T\n", uid)

	// Константы
	const price, tax float32 = 275.00, 27.50
	const quantity, inStock = 2, true
	fmt.Println("Total:", 2 * quantity* (price + tax))

	//Variables
	var cost float32 = 275.00
	var total float32 = 27.50
	cost = 300
	fmt.Println(cost * total)

	// // Error block
	// var value = 275.00
	// var taxes float32 = 27.50
	// fmt.Println(value * taxes)

	// OK block
	var value = 275.00
	var taxes float32 = 27.50
	fmt.Println(value * float64(taxes))

	//Множественное присваивание через :=
	aArg, bArg := 10, 30
	fmt.Println(aArg, bArg)
	aArg, bArg = 30, 40
	fmt.Println(aArg, bArg)
	// aArg, bArg := 10, 30
	// fmt.Println(aArg, bArg)

	//Исключение из этого правила
	bArg, cArg := 300, 400
	fmt.Println(aArg, bArg, cArg)

}