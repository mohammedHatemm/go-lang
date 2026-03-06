package main
import (
	"fmt"
	// "strconv"
	// "runtime"
	// "net/http"
	)

func main() {

	// Run: docker compose run --rm go go run 01-basics/control-flow/main.go
// ### Exercise 1: Grade Calculator

	var grade int = 85
	switch  {
	case grade >= 90:
		fmt.Println("A")
	case grade >= 80:
		fmt.Println("B")
	case grade >= 70:
		fmt.Println("C")
	case grade >= 60:
		fmt.Println("D")
	default:
		fmt.Println("F")
	}



	// ### Exercise 2: FizzBuzz

for i:=0 ; i<=20; i++ {
	if i%3 == 0 && i%5 == 0 {
		fmt.Println("FizzBuzz")
	} else if i%3 == 0 {
		fmt.Println("Fizz")
	} else if i%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(i)
	}
}

//### Exercise 3: Find a Number

for i:=0 ; i<=10; i++ {

	if i==7 {
		fmt.Println("Found it!")
	}
}



}
