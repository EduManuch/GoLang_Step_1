package main

import "fmt"

func main() {
	// 1. Указатели - переменная, хранящая в качестве значения - адрес в памяти другой переменной

	// 2. Определение указателя на что-то
	var variable int = 30
	var pointer *int = &variable // &... - операция взятия адреса в памяти
	// Выше создан указатель на переменную variable
	fmt.Printf("Type of pointer: %T\n", pointer)
	fmt.Printf("Value of pointer: %v\n", pointer)

	// 3. А какое нулевое значение для указателя
	var zeroPointer *int // zeroPointer имеет значение nil(указатель в никуда)
	fmt.Printf("Type: %T and value: %v\n", zeroPointer, zeroPointer)
	if zeroPointer == nil {
		zeroPointer = &variable
		fmt.Printf("After init Type: %T and value: %v\n", zeroPointer, zeroPointer)
	}

	// 4. Разыменовывание указателя(получение значения):
	// *pointer - возвращает значение, хранимое по адресу
	var numericValue int = 32
	pointerToNumeric := &numericValue
	fmt.Printf("Value of pointerToNumeric: %v and it's address %v\n",
				*pointerToNumeric, pointerToNumeric)

	// 5. Создание указателей на нулевые значения типов
	// var zeroVar int
	// var zeroPoint *int = &zeroVar

	zeroNewPoint := new(int)
	// Создает под капотом zeroValue для int и возвращает адрес, где этот 0 хранится
	fmt.Printf("Value of zeroNewPoint: %v and it's address %v\n", *zeroNewPoint, zeroNewPoint)

	// 6. Изменение значения хранимого по адресу через указатель
	zeroPtoInt := new(int)
	fmt.Printf("Value of zeroPtoInt: %v and it's address %v\n", *zeroPtoInt, zeroPtoInt)
	*zeroPtoInt += 40
	fmt.Printf("Value of zeroPtoInt: %v and it's address %v\n", *zeroPtoInt, zeroPtoInt)
	b := 345
	a := &b
	c := &b
	*a++
	*c += 98
	fmt.Println("b:", b)

	// 7. Указательная арифметика ПОЛНОСТЬЮ ОТСУТСТВУЕТ
	// У вас на руках адрес одной ячейки - вы можете через этот адрес перейти в другие ячейки

	// 8. Передача указателей в функции
	// Большой прирост производительности за счет того, 
	// что передается не значение, которое должно копироваться,
	// а передается лишь адрес в памяти, за которым уже храниться какое-то значение
	sample := 1
	fmt.Println("Origin value of sample:", sample)
	changeParam(&sample)
	fmt.Println("Origin value of sample:", sample)

	// 9. Возврат указателя из функции
    ptr1 := returnPointer()		
    ptr2 := returnPointer()
	fmt.Printf("ptr1 -> Type: %T, address: %v and value: %v\n", ptr1, ptr1, *ptr1)
	fmt.Printf("ptr2 -> Type: %T, address: %v and value: %v\n", ptr2, ptr2, *ptr2)

	arr := [3]int{1, 2, 3}
	fmt.Println("Before calling mutation:", arr)
	mutation(&arr)
	fmt.Println("After calling mutation:", arr)

	newArr := [3]int{2, 5, 8}
	fmt.Println("Before calling mutation:", newArr)
	mutationSlice(newArr[:])
	fmt.Println("After calling mutation:", newArr)
}

// 8.1 Определение функции, принимающей указатель как параметр
func changeParam(value *int) {
	*value += 100
}

func returnPointer() *int {
	var numeric int = 321
	return &numeric // В момент возврата Go перемещает данную переменную в кучу
}

func mutation(arr *[3]int) {
	// (*arr)[1] = 900
	// (*arr)[2] = 1000
	// Но можно писать и так, потому что Go сам разыменует указатель на массив
	// (из-за того, что функция принимает *arr)
	arr[1] = 900
	arr[2] = 1000
}

// 11. Лучше используйте слайсы(это идеоматично с точки зрения Go)
func mutationSlice(sls []int) {
	sls[1] = 900
	sls[2] = 1000
}