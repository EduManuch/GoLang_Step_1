/*
Задача №2

Вход:
Пользователь должен ввести "правильный пароль", состоящий из:
цифр, букв латинского алфавита(строчные и прописные) и
специальных символов  special = "_!@#$%^&"

Всего 4 набора различных символов.
В пароле обязательно должен быть хотя бы один символ из каждого набора.
Длина пароля от 8(мин) до 15(макс) символов.
Максимальное количество попыток ввода неправильного пароля - 5.
Каждый раз выводим номер попытки.
*Желательно выводить пояснение, почему пароль не принят и что нужно исправить.

digits = "0123456789"
lowercase = "abcdefghiklmnopqrstvxyz"
uppercase = "ABCDEFGHIKLMNOPQRSTVXYZ"
special = "_!@#$%^&"

Выход:
Написать, что пароль правильный и он принят.

Пример:
хороший пароль -> o58anuahaunH!
хороший пароль -> aaaAAA111!!!
плохой пароль -> saucacAusacu8

Реализацию оформить через функцию:
1. checkPassword(pass string) (bool, errors <- на усмотрение)
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	fmt.Println("Создайте пароль")
	input := bufio.NewScanner(os.Stdin)

	var try uint = 1
	for input.Scan() {
		if try > 5 {
			fmt.Println("Превышено количество попыток")
			break
		}
		fmt.Printf("Попытка %v: \n", try)
		try++

		passed, err := checkPassword(input.Text())
		for _, val := range err {
			fmt.Println(val)
		}
		if passed {
			fmt.Println("Пароль успешно создан")
			break
		}
	}
}

func checkPassword(pass string) (bool, []string) {
	var err []string
	if utf8.RuneCountInString(pass) < 8 || utf8.RuneCountInString(pass) > 15 {
		err = append(err, "Длина пароля должна быть от 8 до 15 символов")
	}

	digits := strToMap("0123456789")
	lowercase := strToMap("abcdefghiklmnopqrstvxyz")
	uppercase := strToMap("ABCDEFGHIKLMNOPQRSTVXYZ")
	special := strToMap("_!@#$%^&")
	var digitsOk, lowerOk, upperOk, specialOk bool

	for _, val := range pass {
		s := string(val)
		if _, ok := digits[s]; ok {
			digitsOk = true
		}
		if _, ok := lowercase[s]; ok {
			lowerOk = true
		}
		if _, ok := uppercase[s]; ok {
			upperOk = true
		}
		if _, ok := special[s]; ok {
			specialOk = true
		}
	}

	if !digitsOk {
		err = append(err, "Пароль должен содержать минимум одну цифру")
	}
	if !lowerOk {
		err = append(err, "Пароль должен содержать минимум одну строчную букву")
	}
	if !upperOk {
		err = append(err, "Пароль должен содержать минимум одну заглавную букву")
	}
	if !specialOk {
		err = append(err, "Пароль должен содержать минимум один спецсимвол")
	}
	if len(err) != 0 {
		return false, err
	}
	return true, err
}

func strToMap(val string) map[string]int {
	result := make(map[string]int)
	for _, v := range val {
		result[string(v)] = 1
	}
	return result
}
