package main

import "fmt"

// 1. Буферизированный канал - канал с буфером, в который можно напихать со стороны отправителя не 1 пак данных,
//  а столько, сколько позволяет в буфер.
// Шаблон
// ch := make(chan int, capacityIntValue)

func main() {
	ch := make(chan string, 5) // Создадим канал вместимостью 5
	ch <- "Bob"                // Не блокируемся , т.к. можно запихнуть в буфер еще 4 элемента
	ch <- "Alex"               // Не блокируемся, т.к. можно еще запихнуть в буфер 3 элемента
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}