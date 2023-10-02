package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func squareOfNumber(num int) int {
	return num * num
}

func testBasics() {
	var message string

	fmt.Println("Input your message")
	//fmt.Scanf("%s", &message)
	message = "Hello"

	fmt.Printf("Message: %s\n", message)

	// Excplicit type conversion
	var number1 int = 10
	var number2 float64 = 100

	sum := number1 + int(number2)

	//Swaping two numbers
	a, b := 10, 20
	fmt.Printf("A = %d, B = %d\n", a, b)

	a, b = b, a
	fmt.Println("After Swaping")
	fmt.Printf("A = %d, B = %d\n", a, b)

	fmt.Println("Sum after explictit type conversion:", sum)

	// User input
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your year of birth")
	scanner.Scan()

	//use parseInt function to get integer from string
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	fmt.Println("Age = ", 2023-input)

	// If else statement with assignment expression before condition
	if x := 20; x > 30 {
		fmt.Println("Inside If ")
	} else if y := 30; x < 25 && y > 20 {
		fmt.Println("Inside else if")
	} else {
		fmt.Println("Inside else")
	}

	// for loop
	for x := squareOfNumber(2); x < 7; {
		fmt.Println(x)

	}
}

func testArray() {
	var arr [5]int
	fmt.Printf("Array: %v\n Type: %T\n", arr, arr)

	// use [...] to set the size equal to nummber of elements inside curly braces
	array1 := [...]int{0, 1, 3}
	fmt.Printf("Array: %v\n Type: %T\n", array1, array1)

	// initialize the elements of index 0 and 3 only
	arrayOfIntegers := [5]int{0: 7, 3: 9}

	fmt.Printf("Array: %v\n Type: %T\n", arrayOfIntegers, arrayOfIntegers)

	// multidemensional array
	array2D := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Printf("Array: %v\n Type: %T\n", array2D, array2D)

	// multidemensional array without size
	array2D1 := [...][3]int{{1, 2}, {4, 5, 6}}
	fmt.Printf("Array: %v\n Type: %T\n", array2D1, array2D1)
}

func testSlices() {
	array := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Array: %v\nType: %T\nLength: %d\nCapacity: %d\n", array, array, len(array), cap(array))

	slice := array[2:7]
	fmt.Printf("Slice: %v\nType: %T\nLength: %d\nCapacity: %d\n", slice, slice, len(slice), cap(slice))
	fmt.Printf("Array: %v\nType: %T\nLength: %d\nCapacity: %d\n", array, array, len(array), cap(array))

	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice: %v\nType: %T\nLength: %d\nCapacity: %d\n", slice1, slice1, len(slice1), cap(slice1))

	slice1 = append(slice1, 7, 8, 9)
	fmt.Printf("Slice: %v\nType: %T\nLength: %d\nCapacity: %d\n", slice1, slice1, len(slice1), cap(slice1))

	slice1 = append(slice1, 1, 2, 3)
	fmt.Printf("Slice: %v\nType: %T\nLength: %d\nCapacity: %d\n", slice1, slice1, len(slice1), cap(slice1))

	for i, ele := range slice1 {
		fmt.Printf("Index: %d Element: %d\n", i, ele)
	}

}

func testMaps() {
	var map1 map[string]int
	map1 = map[string]int{
		"key1": 1,
		"key2": 2,
	}

	map2 := map[string]int{
		"key1": 1,
		"key2": 2,
	}

	map3 := make(map[string]int)

	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	value, ok := map3["key1"]
	fmt.Printf("Value: %d \nisKeyPresent:%t\n", value, ok)

	value, ok1 := map1["key1"]
	fmt.Printf("Value: %d \nisKeyPresent:%t\n", value, ok1)
}

func printFunc(myfunc func(x int) int) {
	fmt.Println(myfunc(7))
}

func returnFunc(x string) func() {
	return func() {
		fmt.Println(x)
	}
}

func testFunction() {
	testFunc := testMaps
	testFunc()

	test := func(x int) int {
		return x * x
	}

	fmt.Println(test(5))

	//Calling a funcion along with definition
	func(x int) {
		fmt.Println(x * x)
	}(6)

	//passing a function as an argument to a function
	printFunc(test)
	printFunc(func(x int) int {
		return x + x
	})

	//return a function
	returnFunc("Hello")()
	returnVar := returnFunc("GoodBye")
	returnVar()
}
func changeString(str *string) {
	*str = "Changed"
}

func testPointers() {
	toChange := "Change me"
	changeString(&toChange)

	print(toChange)

}

type Point struct {
	x int
	y int
}

func changeX(p *Point) {
	p.x = 100
}

func testStruct() {
	var p1 Point = Point{10, 20}
	p2 := Point{x: 5}
	fmt.Println(p1)
	changeX(&p1)
	fmt.Println(p1)
	fmt.Println(p2)
}

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (c *Circle) area() float64 {
	return 2 * 3.14 * c.radius
}

func (r *Rectangle) area() float64 {
	return r.height * r.width
}

func division(num1, num2 int) {
	defer handlePanic()
	// if num2 is 0, program is terminated due to panic
	if num2 == 0 {
		panic("Cannot divide a number by zero")
	} else {
		result := num1 / num2
		fmt.Println("Result: ", result)
	}
}

// recover function to handle panic
func handlePanic() {

	// detect if panic occurs or not
	a := recover()

	if a != nil {
		fmt.Println("RECOVER", a)
	}

}

func main() {

	//Basics
	//testBasics()

	// arrays
	//testArray()
	//testSlices()
	//testMaps()
	//testFunction()
	//testPointers()
	//testStruct()

	c := Circle{5}
	r := Rectangle{4, 5}

	fmt.Println(c.area())
	fmt.Println(r.area())

	shape := []Shape{&c, &r}

	for _, s := range shape {
		fmt.Println(s.area())
	}

	var i interface{}

	i = 10
	value, ok := i.(int)

	fmt.Println(value, ok)

	i = "Hello"

	value2, ok := i.(string)
	fmt.Println(value2, ok)

	division(4, 2)
	division(8, 0)
	division(2, 8)

}
