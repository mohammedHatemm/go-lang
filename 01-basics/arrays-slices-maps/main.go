package main

import (
	"fmt"
)

func main() {
	// Start writing your code here
	// ابدأ اكتب الكود بتاعك هنا

	// Run: docker compose run --rm go go run 01-basics/arrays-slices-maps/main.go

	fmt.Println("=== Arrays, Slices & Maps ===")
/*
*
* Arrays: Fixed-size collections of elements of the same type.
* Slices: Dynamic-size, flexible views into arrays.
* Maps: Collections of key-value pairs, where keys are unique.
*
*if not initialized arrayes values will be set to the zero value of the element type (e.g., 0 for int, "" for string).
*/
	var numbers [5]int
	fmt.Println("Array:", numbers)

	// initializing an array with values
	//and wthiout var

	colors := [3]string{"red", "green", "blue"}
	fmt.Println("Array:", colors)

//intializing array without specifying the size
	fruits := [...]string{"apple", "banana", "cherry"}
	fmt.Println("Array:", fruits)

	//changing array values
	numbers[0] = 10
	numbers[1] = 20
	fmt.Println("Updated Array:", numbers)

	// Slices: Dynamic-size, flexible views into arrays.
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", slice)

	//empty slice
	// not like arrays
				//-> arrays if it empty it will be filled with the zero value of the element type
				//-> slices if it empty it will be nil

	var emptySlice []string
	fmt.Println("Empty Slice:", emptySlice)

	scores := make([]string, 5 )
	fmt.Println("Scores Slice:", scores)
	fmt.Println("Length of Scores Slice:", len(scores))
	fmt.Println("Capacity of Scores Slice:", cap(scores))

// append to a slice
// append is a built-in function that adds elements to the end of a slice and returns the updated slice.
	slice = append(slice, 6, 7)
	fmt.Println("Updated Slice:", slice)

	makeMapes := map[string]int{

"apple":  1,
"banana": 2,
"cherry": 3,
	}

	fmt.Println("Map:", makeMapes)


	mapesIntString := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	fmt.Println("Map:", mapesIntString)






fmt.Println(hasDuplcates([]int{1,2,3,4,5}))

//grouping

students := map[int][]string{
	1 : {"mohamed" , "ahmed"},
	2 : {"djfs" , "sdfsd"},
}

students[1] = append(students[1], "hossam")
fmt.Println(students)
students[2] = append(students[2], "ahmed")
fmt.Println(students)



//### Exercise 1: Slice Operations

favFood := []string{"makaronaBashamel" , "mahawy"}

newFavFood := append(favFood,"feraks")
fmt.Println(newFavFood)


theArrayWithoutSacoundElemet := []string{}
for i , newNEwFood := range newFavFood{
	if i== 1 {
		 continue
	}
	fmt.Println(newNEwFood)
	theArrayWithoutSacoundElemet = append(theArrayWithoutSacoundElemet , newNEwFood)
}
fmt.Println(theArrayWithoutSacoundElemet)

//### Exercise 2: Sum and Average
fmt.Println(calcualtAverage([]float64{1.00,2.00,3.00,4.00,5.00}))


//### Exercise 3: Word Counter

words := []string{"apple", "banana", "apple", "cherry", "banana"}
makeCount := make(map[string]int)
for _, word := range words {
	makeCount[word]++
}
fmt.Println(makeCount)


//### Exercise 4: Student Grades
studentss := map[string]int{

	"mohamed": 100,
	"ahmed": 62,
}
for name , grade := range studentss{
	fmt.Printf("%s is %d grade\n" , name , grade)
}

fmt.Println(studentWithHighestGrade(studentss))

//### Exercise 5: Remove Duplicates
numbersss := []int{1, 2, 2, 3, 4, 4, 5}


fmt.Println(uniqueNumbers(numbersss))



}
func hasDuplcates(numbers []int) bool {
	seen := make(map[int]bool)
	for _, num:= range numbers{
		if seen[num] {
			return true
		}
			seen[num] = true
	}
	return false
	}


	func calcualtAverage(numbers []float64) float64{
		sum := 0.0
		for _, num := range numbers{
			sum += num

		}
		Average := sum / float64(len(numbers))
		fmt.Println(Average)

		return Average

	}
func studentWithHighestGrade(students map[string]int) string {
	highestGrade := 0
	var highestGradeStudent string
	for name, grade := range students {
		if grade > highestGrade {
			highestGrade = grade
			highestGradeStudent = name
		}
	}
	return highestGradeStudent
}


func uniqueNumbers(numbers []int) []int {
	unique := []int{}
	seen := make(map[int]bool)
	for _, num := range numbers {
		if !seen[num] {
			unique = append(unique, num)
			seen[num] = true
		}
	}
	return unique
}
