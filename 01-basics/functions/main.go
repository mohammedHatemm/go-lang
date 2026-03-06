package main

import (
	"fmt"
)

func main() {
	// Start writing your code here
	// ابدأ اكتب الكود بتاعك هنا

	// Run: docker compose run --rm go go run 01-basics/functions/main.go

	fmt.Println("=== Functions ===")
	sayHello()
	fmt.Println(add(2, 3))
	fmt.Println(greet("John"))
	result := add(5, 7)
	fmt.Println(result)
	checkAge(20)
	fmt.Println(checkAge(20))
	fmt.Println(sum(1, 2, 3, 4, 5))
	counterFunc := counter()
	fmt.Println(counterFunc()) // Output: 1
	fmt.Println(counterFunc()) // Output: 2
	fmt.Println(counterFunc()) // Output: 3
	deferTest()

}

func sayHello() {
	fmt.Println("Hello, World!")
}

func add(a, b int) int {
	return a + b
}

func greet(name string) string {
	return "Hello, " + name + "!"
}

func checkAge(age int ) bool {
	if age >= 18 {
		return true
	} else {
		return false
	}
}
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}

}
func deferTest() {
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")
}
