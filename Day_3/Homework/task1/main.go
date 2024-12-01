/*
Задача №1
Написать функцию, которая расшифрует строку.
code = "220411112603141304"
Каждые две цифры - это либо буква латинского алфавита в нижнем регистре либо пробел.
Отчет с 00 -> 'a' и до 25 -> 'z', 26 -> ' '(пробел).
Вход: строка из цифр. Выход: Текст.
Проверка работы функции выполняется через другую строку.

Задача №1.1
Реализовать и функцию зашифровки

codeToString(code) -> "???????'

stringToCode("hello") -> "??????"
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	codeMap := mapForCode()

	forDecode := "220411112603141304"
	decodeResult := codeToString(forDecode, codeMap)
	fmt.Println(decodeResult)

	forCode := "well done"
	encodeResult := stringToCode(forCode, codeMap)
	fmt.Println(encodeResult)
}

func mapForCode() map[int]string {
	m := make(map[int]string)
	alph := "abcdefghijklmnopqrstuvwxyz "
	for i, val := range alph {
		m[i] = string(val)
	}
	return m
}

func codeToString(code string, codeMap map[int]string) (result string) {
	for i := range code {
		if i%2 != 0 {
			key, _ := strconv.Atoi(code[i-1 : i+1])
			result += codeMap[int(key)]
		}
	}
	return
}

func stringToCode(forCode string, codeMap map[int]string) (result string) {
	for _, val := range forCode {
		for k, v := range codeMap {
			if v == string(val) {
				result += fmt.Sprintf("%02d", k)
			}
		}
	}
	return
}
