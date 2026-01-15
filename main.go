package main

import "fmt"
import "github.com/TheAditya-10/go/mathutils"

func main() {
	fmt.Println("Hello World, let's Go")
	var name string = "Aditya"
	var age int = 19

	// var a int = 10
	// var b float64 = 3.14
	// var c bool = true
	// var d string = "backend"

	// var x int      // 0
	// var y string   // ""
	// var z bool     // false

	const pi = 3.14
	const AppName = "GoApp"

	if age >= 18 {
		fmt.Println(name, "is an adult.")
	} else {
		fmt.Println(name, "is a minor.")
	}

	if x:= 10; x > 5 {
		fmt.Println(x, "is greater than 5")
	} else {
		fmt.Println(x, "is not greater than 5")
	}

	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Other day")
	}

	for i:= 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}

	i:= 0
	for i < 5 {
		fmt.Println("Count:", i)
		i++
	}

	for {
		fmt.Println("Infinite Loop")
		break
	}
	sum := add(5, 7)
	fmt.Println("Sum:", sum)

	result, ok := divide(10, 2)
	if ok {
		fmt.Println("Division Result:", result)
	} else {
		fmt.Println("Division by zero is not allowed.")
	}
	product := mathutils.Multiply(4, 5)
	fmt.Println("Product:", product)
}

func add(a int, b int) int {
	return a + b
}

func divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}