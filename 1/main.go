package main

import (
	"fmt"
	"math/rand"
)

var Version string = "1.0"
var Name string = "Go Practise tour"

const Pi = 3.14

func add(x, y int) int {
	return x + y
}

func count_symbols(usr_str string) int {
	counter := 0
	for i := 0; i < len(usr_str); i++ {
		counter++
	}
	return counter
}

func add_and_substruct(x, y int) (int, int) {
	return x + y, x - y
}

// naked return example
func multiplie_and_divide(x, y int) (first_result int, second_result float32) {
	first_result = x * y
	second_result = float32(x) / float32(y)
	return
}

func main() {
	random_num := rand.Intn(100)
	fmt.Println(random_num)

	fmt.Println(add(2, 6))

	fmt.Println(count_symbols("Hello World!"))

	fmt.Println(add_and_substruct(10, 5))

	fmt.Println(multiplie_and_divide(15, 2))

	var counter int = 10
	fmt.Println(counter)

	var name, age = "Alice", 25
	fmt.Println(name, age)

	temperature := 37.5
	fmt.Println(temperature)

	var toBe bool = false
	fmt.Println(toBe)

	var zero_int int
	var zero_bool bool
	var zero_string string
	fmt.Println(zero_bool, zero_int, zero_string)

	floated_counter := float32(counter) + 0.1
	fmt.Println(floated_counter)

	complex := 5i + 1
	fmt.Printf("%f has type %T\n", complex, complex)

	fmt.Println(Pi)
}
