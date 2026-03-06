package main
import (
	"fmt"
	"strconv"


)


func main() {
	// Start writing your code here
	// ابدأ اكتب الكود بتاعك هنا

	// Run: docker compose run --rm go go run 01-basics/variables/main.go
var name string = "mohamed elsherif"
var age int = 28
var height float64 = 1.80
var isMarried bool = false
isWorking := true

//convert the length to float64 and assign it to a variable
var length float64 = 19.99
// iota example
const (
	january   = iota + 1
	february
	march
	april
	may
	june
	july
)



fmt.Println("Hello, world!")
fmt.Println("My name is " + name)
fmt.Println("I am " + strconv.Itoa(age) + " years old")
fmt.Println("My height is " + strconv.FormatFloat(height, 'f', 2, 64) + " meters")
fmt.Println("Am I married? " + strconv.FormatBool(isMarried))
fmt.Println(isWorking)
fmt.Println("the length of my name is " + strconv.Itoa(len(name)) + " characters")
fmt.Println("converting length to int: " + strconv.Itoa(int(length)))
// Print the month of April using iota
fmt.Println("The month of April is: " + strconv.Itoa(april))

}
