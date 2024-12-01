/*

=======
Задачи:
=======


3.1 Пользователь вводит числа a и b (b больше a).
    Вывести все целые числа, расположенные между ними.

3.2 Доработать предыдущую задачу так, чтобы выводились только числа,
    делящиеся на 5 без остатка.

3.3 Пользователь вводит число. Вывести таблицу умножения на это число (от 1 до 10)

3.4 В цикле получать от пользователя оценки по четырём экзаменам.
    Вывести сумму набранных им баллов.
    Функцию fmt.Scan() в коде используем только один раз.

3.5 В бесконечном цикле приложение запрашивает у пользователя числа.
    Ввод завершается, как только пользователь вводит число "-1".
    После завершения ввода приложение выводит сумму чисел.
    Используем конструкцию:
    for {
      // body
    }

3.6 Получить от пользователя натуральное число, посчитать сумму цифр в нём.
    Решить с помощью индексов, т.е. работаем с числом, как со строкой.

3.7 Вводим строку без знаков препинания(5 слов).
    Найти самое длинное слово в строке и вывести сколько в нем букв.

Пример:
In: Скажи как дела в учебе
Out: учебе, 5

In: Закрепляем циклы в языке Golang
Out: Закрепляем, 10
*/

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 	// 3.1, 3.2
	// 	var a, b int

	// 	fmt.Println("Enter numbers:")
	// 	fmt.Scan(&a, &b)

	// outer:
	// 	for {
	// 		a += 1
	// 		switch {
	// 		case a == b:
	// 			fmt.Println(a)
	// 			break outer
	// 		case a%5 == 0:
	// 			fmt.Println(a)
	// 		}
	// 	}

	// 	// 3.3
	// 	var c int
	// 	fmt.Println("Enter number again:")
	// 	fmt.Scan(&c)
	// 	for i := 1; i < 10; i++ {
	// 		fmt.Printf("%d * %d = %d\n", c, i, c*i)
	// 	}

	// 	// 3.4
	// 	var sum int
	// 	for i := 0; i < 4; i++ {
	// 		var score int
	// 		fmt.Println("Enter scores:")
	// 		fmt.Scan(&score)
	// 		sum += score
	// 	}
	// 	fmt.Printf("Sum of scores: %d\n", sum)

	// 	// 3.5
	// 	var summa int
	// outer2:
	// 	for {
	// 		var d int
	// 		fmt.Println("Enter next number:")
	// 		fmt.Scan(&d)
	// 		if d == -1 {
	// 			break outer2
	// 		}
	// 		summa += d
	// 	}
	// 	fmt.Println("Total sum is", summa)

	// // 3.6
	// var num uint
	// fmt.Println("Enter new num:")
	// fmt.Scan(&num)

	// var sumNat uint
	// numStr := fmt.Sprintf("%d", num)
	// for i := range numStr {
	// 	char := fmt.Sprintf("%c", numStr[i])
	// 	digit, err := strconv.Atoi(char)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		break
	// 	}
	// 	sumNat += uint(digit)
	// }
	// fmt.Println("Sum of natural is", sumNat)

	//3.7
	var w [5]string
	fmt.Println("Enter 5 words:")
	for i := range w {
		fmt.Scan(&w[i])
	}

	var wMax string
	var wMaxLen int
	for i := range w {
		tmpLen := utf8.RuneCountInString(w[i])
		if tmpLen > wMaxLen {
			wMax = w[i]
			wMaxLen = tmpLen
		}
	}
	fmt.Printf("Word %v has max length %v\n", wMax, wMaxLen)

}
