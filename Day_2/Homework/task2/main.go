/*
Задача №1.
Программа получает на вход последовательность из 5 целых чисел.

Вам нужно определить вид последовательности:
 - возрастающая
 - убывающая
 - случайная
 - постоянная

В качестве ответа следуют выдать прописными латинскими буквами тип последовательности:
1. ASCENDING (строго возрастающая)
2. WEAKLY ASCENDING (нестрого возрастающая, то есть неубывающая)
3. DESCENDING (строго убывающая)
4. WEAKLY DESCENDING (нестрого убывающая, то есть невозрастающая)
5. CONSTANT (постоянная)
7. RANDOM (случайная)

Примеры входных и выходных данных:
In: 11 9 4 2 -1  Out: DESCENDING
In: 3 8 8 11 12  Out: WEAKLY ASCENDING
In: 2 -1 7 21 1  Out: RANDOM
In: 5 5 5 5 5     Out: CONSTANT

Подсказка: используем метод строки strings.Split()
*/

package main

import "fmt"

func main() {
	var arr [5]int

	fmt.Println("Enter 5 numbers:")
	for i := range arr {
		fmt.Scan(&arr[i])
	}

	switch {
	case arr[0] == arr[1] &&
		arr[1] == arr[2] &&
		arr[2] == arr[3] &&
		arr[3] == arr[4]:
		fmt.Println("CONSTANT")
	case arr[0] > arr[1] &&
		arr[1] > arr[2] &&
		arr[2] > arr[3] &&
		arr[3] > arr[4]:
		fmt.Println("DESCENDING")
	case arr[0] < arr[1] &&
		arr[1] < arr[2] &&
		arr[2] < arr[3] &&
		arr[3] < arr[4]:
		fmt.Println("ASCENDING")
	case (arr[0] > arr[1] || arr[0] == arr[1]) &&
		(arr[1] > arr[2] || arr[1] == arr[2]) &&
		(arr[2] > arr[3] || arr[2] == arr[3]) &&
		(arr[3] > arr[4] || arr[3] == arr[4]):
		fmt.Println("WEAKLY DESCENDING")
	case (arr[0] < arr[1] || arr[0] == arr[1]) &&
		(arr[1] < arr[2] || arr[1] == arr[2]) &&
		(arr[2] < arr[3] || arr[2] == arr[3]) &&
		(arr[3] < arr[4] || arr[3] == arr[4]):
		fmt.Println("WEAKLY ASCENDING")
	default:
		fmt.Println("RANDOM")
	}
}
