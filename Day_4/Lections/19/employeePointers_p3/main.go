package main

import "fmt"

type Employee struct {
	name string
	salary int
}

// 1. Метод, в котором получатель копируется и в его теле происходит работа с локальной копией
func (e Employee) SetName(newName string) {
	e.name = newName
}

// 2. Метод, в котором получатель передается по ссылке (в теле метода работаем с ссылкой на экземпляр)
func (e *Employee) SetSalary(newSalary int) {
	e.salary = newSalary
}

func main() {
	e := Employee{"Alex", 3000}
	fmt.Println("Before setting new name:", e)
	e.SetName("Yan")
	fmt.Println("After setting new name:", e)

	// 3. Вызов метода у ссылки на сотрудника
	e.SetSalary(4500)
	// 5. Аналогично явному вызову (&e).SetSalary(4500)
	fmt.Println("After setting new salary:", e)
}
