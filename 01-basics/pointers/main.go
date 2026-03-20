package main

import (
	"fmt"
)

func main() {
	// Start writing your code here
	// ابدأ اكتب الكود بتاعك هنا

	// Run: docker compose run --rm go go run 01-basics/pointers/main.go

	fmt.Println("=== Pointers ===")
	age:=25;
	pointerAddress := &age
	fmt.Println(pointerAddress)
	fmt.Println(*pointerAddress)

// double(age)
fmt.Println(age);
double(&age)
	fmt.Println(age);


}
func double(n *int){
	*n=(*n)*2
	// n*2
	// fmt.Println(n)
}
