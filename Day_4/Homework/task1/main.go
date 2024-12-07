/*
Сформировать данные для отправки заказа из
магазина по накладной и вывести на экран:
1) Наименование товара (минимум 1, максимум 100)
2) Количество (только числа)
3) ФИО покупателя (только буквы)
4) Контактный телефон (10 цифр)
5) Адрес: индекс(ровно 6 цифр), город, улица, дом, квартира

Эти данные не могут быть пустыми.
Проверить правильность заполнения полей.

Реализовать конструктор и несколько методов у типа "Накладная"

Пример:
invoice = NewInvoice()

или

order = NewOrder()
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"unicode/utf8"
)

type Invoice struct {
	product  string
	quantity string
	customer Person
	phone    string
	address  Address
}

func NewInvoice(fio, address []string, product, quantity, phone *string) *Invoice {
	customer := NewPerson(fio)
	addr := NewAddress(address)

	inv := Invoice{*product, *quantity, *customer, *phone, *addr}
	inv.validateProduct()
	inv.validateQuantity()
	inv.validatePhone()
	return &inv
}

func (i *Invoice) validateProduct() {
	for utf8.RuneCountInString(i.product) < 1 || utf8.RuneCountInString(i.product) > 100 {
		fmt.Println("Длина наименования товара должна быть минимум 1 и максимум 100 символов. Повторите ввод:")
		fmt.Scan(&i.product)
	}
}

func (i *Invoice) validateQuantity() {
	re := regexp.MustCompile(`^[1-9]+$`)
	for !re.MatchString(i.quantity) {
		fmt.Println("Количество должно состоять из цифр. Повторите ввод:")
		fmt.Scan(&i.quantity)
	}
}

func (i *Invoice) validatePhone() {
	re := regexp.MustCompile(`^\d{10}$`)
	fmt.Println(i.phone)
	for !re.MatchString(i.phone) {
		fmt.Println("Неверно введен номер телефона. Повторите ввод:")
		fmt.Scan(&i.phone)
	}
}

func (i *Invoice) print() {
	fmt.Println("=====Заказ=====")
	fmt.Println("Товар:", i.product)
	fmt.Println("Количество:", i.quantity)
	fmt.Println("ФИО:", i.customer.lastName, i.customer.firstName, i.customer.middleName)
	fmt.Println("Номер телефона:", i.phone)
	fmt.Printf("Адрес: %v, город %v, улица %v, дом %v, кв. %v\n", i.address.postcode, i.address.city, i.address.street, i.address.house, i.address.apparment)
	fmt.Println("===============")
}

type Person struct {
	firstName  string
	lastName   string
	middleName string
}

func NewPerson(fio []string) *Person {
	re := regexp.MustCompile(`[A-Za-zА-Яа-я\-]`)
	for {
		if !re.MatchString(fio[0]) {
			fmt.Println("Некорретно введена фамилия. Повторите ввод:")
			fmt.Scan(&fio[0])
			continue
		}
		if !re.MatchString(fio[1]) {
			fmt.Println("Некорретно введено имя. Повторите ввод:")
			fmt.Scan(&fio[1])
			continue
		}
		if !re.MatchString(fio[2]) {
			fmt.Println("Некорретно введено отчество. Повторите ввод:")
			fmt.Scan(&fio[2])
			continue
		}
		break
	}
	return &Person{fio[0], fio[1], fio[2]}
}

type Address struct {
	postcode, city, street, house, apparment string
}

func NewAddress(addr []string) *Address {

	reStr := regexp.MustCompile(`[A-Za-zА-Яа-я\-]`)
	reIdx := regexp.MustCompile(`^\d{6}$`)
	for {
		if !reIdx.MatchString(addr[0]) {
			fmt.Println("Некорретно введен почтовый индекс. Повторите ввод:")
			fmt.Scan(&addr[0])
			continue
		}
		if !reStr.MatchString(addr[1]) {
			fmt.Println("Некорретно введен город. Повторите ввод:")
			fmt.Scan(&addr[1])
			continue
		}
		if addr[2] == "" {
			fmt.Println("Не введенa улица. Повторите ввод:")
			fmt.Scan(&addr[2])
			continue
		}
		if addr[3] == "" {
			fmt.Println("Не введен дом. Повторите ввод:")
			fmt.Scan(&addr[3])
			continue
		}
		if addr[4] == "" {
			fmt.Println("Не введена квартира. Повторите ввод:")
			fmt.Scan(&addr[4])
			continue
		}
		break
	}

	return &Address{addr[0], addr[1], addr[2], addr[3], addr[4]}
}

const limitTry uint8 = 5

func main() {

	var product, quantity, phone string
	fio := make([]string, 0, 3)
	address := make([]string, 0, 5)
	var try uint8

	fmt.Println("=======Начало ввода заказа=======")
	scanner := bufio.NewScanner(os.Stdin)

	for try <= limitTry {
		try++

		if product == "" {
			fmt.Print("Введите наименование товара: ")
			scanner.Scan()
			product = scanner.Text()
		}
		if quantity == "" {
			fmt.Print("Введите введите количество: ")
			scanner.Scan()
			quantity = scanner.Text()
		}
		if len(fio) == 0 {
			fmt.Print("Введите фамилию: ")
			scanner.Scan()
			fio = append(fio, scanner.Text())

			fmt.Print("Введите имя: ")
			scanner.Scan()
			fio = append(fio, scanner.Text())

			fmt.Print("Введите отчество: ")
			scanner.Scan()
			fio = append(fio, scanner.Text())
		}
		if phone == "" {
			fmt.Print("Введите телефон: ")
			scanner.Scan()
			phone = scanner.Text()
		}
		if len(address) == 0 {
			fmt.Print("Введите почтовый индекс: ")
			scanner.Scan()
			address = append(address, scanner.Text())

			fmt.Print("Введите город: ")
			scanner.Scan()
			address = append(address, scanner.Text())

			fmt.Print("Введите улицу: ")
			scanner.Scan()
			address = append(address, scanner.Text())

			fmt.Print("Введите номер дома: ")
			scanner.Scan()
			address = append(address, scanner.Text())

			fmt.Print("Введите номер квартиры: ")
			scanner.Scan()
			address = append(address, scanner.Text())
		}

		if product != "" && quantity != "" && len(fio) != 0 &&
			phone != "" && len(address) != 0 {
			break
		}
	}
	inv := NewInvoice(fio, address, &product, &quantity, &phone)
	inv.print()
}
