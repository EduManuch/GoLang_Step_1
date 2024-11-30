/*
Задача №1
Вход:

	расстояние(50 - 10000 км),
	расход в литрах (5-25 литров) на 100 км и
	стоимость бензина(константа) = 48 руб

Выход: стоимость поездки в рублях
*Проверка условий по желанию
*/
package main

import "fmt"

func main() {
	const price float64 = 48
	var (
		distance float64
		fuel     float64
	)

	fmt.Print("Enter distance: ")
	fmt.Scan(&distance)
	fmt.Print("Enter fuel: ")
	fmt.Scan(&fuel)

	if distance < 50 || distance > 10000 {
		fmt.Println("Incorrect distance")
		return
	}
	if fuel < 5 || fuel > 25 {
		fmt.Println("Incorrect fuel")
		return
	}

	cost := distance / 100 * fuel * price
	fmt.Printf("Total cost is %.2f\n", cost)
}
