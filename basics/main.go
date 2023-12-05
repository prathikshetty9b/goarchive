package main

import (
	"fmt"

	"github.com/prathikshetty9b/goarchive/pkg/basics"
)

func display[T any](s T) {
	fmt.Println(s)
}

func main() {
	//Grade Problem
	// grade, err := basics.GetGrade()
	// if err != nil {
	// 	fmt.Print(err)
	// 	return
	// }
	// fmt.Print("Grade :", grade)

	basics.GetDay(5)

	//arrays
	array := [5]int{1, 2, 3, 4, 5}
	array[2] = 21
	slice := []int{2, 3, 4}
	slice[1] = 8
	fmt.Println(array)
	fmt.Println(slice)

	text := "Hello"

	display(string(text[1]))

}
