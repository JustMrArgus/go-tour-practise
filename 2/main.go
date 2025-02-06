package main

import (
	"fmt"
)

func print_all_even(n int) {
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd. what a mess")
		}
	}
}

func sum_n(n int) int {
	sum := 0
	for i := 1; i < n; i++ {
		sum += i
	}
	return sum
}

func hadnle_command(command string) {
	switch command {
	case "Add":
		fmt.Println("Add operation")
	case "Update":
		fmt.Println("Update operation")
	case "Delete":
		fmt.Println("Delete operation")
	default:
		fmt.Println("Unknown operation")
	}
}

func check_number(n int) {
	switch {
	case n > 0:
		fmt.Println("Greater than zero")
	case n == 0:
		fmt.Println("Equal to zero")
	case n < 0:
		fmt.Println("Lower than zero")
	}
}

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	counter := 1
	for counter <= 10 {
		fmt.Println(counter)
		counter++
	}

	print_all_even(10)

	infinite_counter := 0
	for {
		fmt.Println(infinite_counter)
		if infinite_counter == 5 {
			break
		}
		infinite_counter++
	}

	fmt.Println(sum_n(5))

	hadnle_command("Add")

	check_number(0)

	fmt.Println("Start")
	defer fmt.Println("I guess i will be last")
	defer fmt.Println("I guess i will be second")
	defer fmt.Println("I guess i will be first")
	fmt.Println("End")
}
