package main

import "fmt"

/*
Оператор defer откладывает выполнение функции до возврата из окружающей функции.
Аргументы отложенного вызова оцениваются немедленно,
но вызов функции не выполняется до возврата из окружающей функции.
*/

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

}
