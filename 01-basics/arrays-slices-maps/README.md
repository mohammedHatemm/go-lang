# Arrays, Slices & Maps

---

## What are Collections?

So far, we've worked with single values: one name, one age, one score. But in real applications, you need to store **groups** of related data: a list of users, a set of scores, a dictionary of settings.

Go gives you three main collection types:

| Type | What it is | Size | Example |
|------|-----------|------|---------|
| **Array** | Fixed-size list | Cannot change | `[5]int` — always 5 elements |
| **Slice** | Dynamic list | Can grow/shrink | `[]int` — any number of elements |
| **Map** | Key-value pairs | Can grow/shrink | `map[string]int` — like a dictionary |

**In practice**: You'll use **slices** 95% of the time, **maps** 4%, and **arrays** 1%. But you need to understand arrays first because slices are built on top of them.

---

## Real World Usage

| Pattern | Where companies use it |
|---------|----------------------|
| Slices | Store list of users, products, orders from database |
| Maps | Store config values, cache data, count occurrences |
| Arrays | Fixed-size buffers, cryptographic hashes (rare) |
| Slice of structs | API responses: `[]User`, `[]Product` |
| Map of slices | Group data: `map[string][]Order` (orders by customer) |
| Nested maps | JSON parsing: `map[string]interface{}` |

---

## Part 1: Arrays

### Declaring arrays

```go
// Method 1: Declare with var
var numbers [5]int
fmt.Println(numbers)  // [0 0 0 0 0] — zero values!

// Method 2: Declare and initialize
scores := [5]int{90, 85, 78, 92, 88}
fmt.Println(scores)   // [90 85 78 92 88]

// Method 3: Let Go count the size
colors := [...]string{"red", "green", "blue"}
fmt.Println(colors)   // [red green blue]
```

- `[5]int` = an array of exactly 5 integers
- `[...]` = "Go, count how many elements I gave you"
- Arrays have **zero values** — an uninitialized `[5]int` is `[0 0 0 0 0]`

### Accessing and modifying elements

```go
scores := [5]int{90, 85, 78, 92, 88}

// Access by index (starts at 0)
fmt.Println(scores[0])  // 90 (first element)
fmt.Println(scores[4])  // 88 (last element)

// Modify an element
scores[1] = 95
fmt.Println(scores)  // [90 95 78 92 88]

// Get the length
fmt.Println(len(scores))  // 5
```

### Iterating over an array

```go
scores := [5]int{90, 85, 78, 92, 88}

// Method 1: Classic for loop
for i := 0; i < len(scores); i++ {
    fmt.Println(i, scores[i])
}

// Method 2: for range (preferred)
for index, value := range scores {
    fmt.Println(index, value)
}

// Method 3: Only values
for _, value := range scores {
    fmt.Println(value)
}
```

### Why arrays are rarely used

```go
a := [3]int{1, 2, 3}
b := a       // COPIES the entire array!
b[0] = 999
fmt.Println(a)  // [1 2 3]   — a is unchanged
fmt.Println(b)  // [999 2 3] — b is a separate copy
```

- Arrays are **copied** when assigned or passed to functions
- You can't change the size after creation
- That's why we use **slices** instead

---

## Part 2: Slices (The One You'll Actually Use)

A slice is like an array but **dynamic** — it can grow and shrink.

### Creating slices

```go
// Method 1: From an array literal (most common)
names := []string{"Ali", "Mohamed", "Sherif"}
fmt.Println(names)  // [Ali Mohamed Sherif]

// Method 2: Empty slice
var numbers []int
fmt.Println(numbers)       // []
fmt.Println(len(numbers))  // 0

// Method 3: Using make (when you know the size you'll need)
scores := make([]int, 5)       // length 5, capacity 5
fmt.Println(scores)             // [0 0 0 0 0]

grades := make([]int, 0, 10)   // length 0, capacity 10
fmt.Println(len(grades))       // 0
fmt.Println(cap(grades))       // 10
```

- `[]string` (no size) = slice. `[3]string` (with size) = array
- `make([]int, length, capacity)` — capacity is optional
- **length** = how many elements it has now
- **capacity** = how many it can hold before needing to grow

### append — Adding elements

```go
names := []string{"Ali", "Mohamed"}
fmt.Println(names)  // [Ali Mohamed]

// Add one element
names = append(names, "Sherif")
fmt.Println(names)  // [Ali Mohamed Sherif]

// Add multiple elements
names = append(names, "Ahmed", "Omar")
fmt.Println(names)  // [Ali Mohamed Sherif Ahmed Omar]
```

- `append` returns a **new** slice — you MUST reassign it: `names = append(names, ...)`
- If you just write `append(names, "Sherif")` without `names =`, the new element is lost!

### Slicing — Getting a part of a slice

```go
numbers := []int{10, 20, 30, 40, 50}

// slice[start:end] — includes start, excludes end
fmt.Println(numbers[1:3])   // [20 30]     — index 1 and 2
fmt.Println(numbers[:3])    // [10 20 30]  — from start to index 2
fmt.Println(numbers[2:])    // [30 40 50]  — from index 2 to end
fmt.Println(numbers[:])     // [10 20 30 40 50] — copy of all
```

Think of it as: `[start:end]` means "from start up to (but not including) end".

**Important**: Slicing creates a **reference**, not a copy:

```go
original := []int{1, 2, 3, 4, 5}
slice := original[1:3]  // [2, 3]

slice[0] = 999
fmt.Println(original)  // [1 999 3 4 5] — original changed!
```

### Removing elements from a slice

```go
numbers := []int{10, 20, 30, 40, 50}

// Remove element at index 2 (value 30)
numbers = append(numbers[:2], numbers[3:]...)
fmt.Println(numbers)  // [10 20 40 50]
```

- `numbers[:2]` = `[10, 20]` (before the element)
- `numbers[3:]` = `[40, 50]` (after the element)
- `...` unpacks the second slice into individual arguments

### Copying slices

```go
original := []int{1, 2, 3}

// The safe way to copy
copied := make([]int, len(original))
copy(copied, original)

copied[0] = 999
fmt.Println(original)  // [1 2 3] — safe!
fmt.Println(copied)    // [999 2 3]
```

### Iterating over slices

Same as arrays — use `for range`:

```go
users := []string{"Ali", "Mohamed", "Sherif"}

for i, user := range users {
    fmt.Printf("%d: %s\n", i, user)
}
// 0: Ali
// 1: Mohamed
// 2: Sherif
```

---

## Part 3: How Slices Work Under the Hood

A slice is actually a small struct with three fields:

```
Slice Header (24 bytes):
+----------+--------+----------+
| Pointer  | Length  | Capacity |
| (to data)|   (3)  |    (5)   |
+----------+--------+----------+
     |
     v
+---+---+---+---+---+
| 1 | 2 | 3 | _ | _ |
+---+---+---+---+---+
     Underlying Array
```

- **Pointer**: points to the first element in the underlying array
- **Length**: how many elements the slice currently has (`len()`)
- **Capacity**: how many elements fit before needing a new array (`cap()`)

### What happens when you append?

```go
s := make([]int, 3, 5)  // len=3, cap=5
// Room for 2 more elements without growing

s = append(s, 4)  // len=4, cap=5 — fits!
s = append(s, 5)  // len=5, cap=5 — fits!
s = append(s, 6)  // len=6, cap=10 — HAD TO GROW!
// Go created a new array with double capacity
```

When capacity is full, Go:
1. Creates a new, larger underlying array (usually 2x)
2. Copies all elements to the new array
3. Returns a slice pointing to the new array

**Tip**: If you know how many elements you'll need, use `make([]int, 0, expectedSize)` to avoid multiple growths.

---

## Part 4: Maps

A map stores **key-value pairs** — like a dictionary.

### Creating maps

```go
// Method 1: Using make
ages := make(map[string]int)
ages["Ali"] = 25
ages["Mohamed"] = 28
ages["Sherif"] = 30
fmt.Println(ages)  // map[Ali:25 Mohamed:28 Sherif:30]

// Method 2: Map literal (most common)
ages := map[string]int{
    "Ali":     25,
    "Mohamed": 28,
    "Sherif":  30,
}
```

- `map[string]int` = keys are `string`, values are `int`
- `map[KeyType]ValueType`

### Accessing values

```go
ages := map[string]int{
    "Ali":     25,
    "Mohamed": 28,
}

// Get a value
fmt.Println(ages["Ali"])  // 25

// Key doesn't exist — returns zero value
fmt.Println(ages["Omar"])  // 0

// Check if key exists
age, exists := ages["Omar"]
fmt.Println(age, exists)  // 0 false

age, exists = ages["Ali"]
fmt.Println(age, exists)  // 25 true

// Common pattern
if age, ok := ages["Ali"]; ok {
    fmt.Println("Ali's age:", age)
} else {
    fmt.Println("Ali not found")
}
```

- When a key doesn't exist, you get the **zero value** (0 for int, "" for string)
- Always use the **two-value form** to check if the key exists: `value, ok := myMap[key]`
- `ok` is `true` if key exists, `false` if not

### Adding, updating, deleting

```go
colors := map[string]string{
    "r": "red",
    "g": "green",
}

// Add new key
colors["b"] = "blue"

// Update existing key
colors["r"] = "dark red"

// Delete a key
delete(colors, "g")

fmt.Println(colors)  // map[b:blue r:dark red]
```

### Iterating over a map

```go
ages := map[string]int{
    "Ali":     25,
    "Mohamed": 28,
    "Sherif":  30,
}

for name, age := range ages {
    fmt.Printf("%s is %d years old\n", name, age)
}
```

**Important**: Map iteration order is **random**! Go randomizes it on purpose. Don't depend on any specific order.

### Getting the length

```go
ages := map[string]int{"Ali": 25, "Mohamed": 28}
fmt.Println(len(ages))  // 2
```

---

## Part 5: Common Patterns

### Counting occurrences

```go
words := []string{"go", "is", "go", "fun", "go", "is", "great"}

count := make(map[string]int)
for _, word := range words {
    count[word]++
}

fmt.Println(count)  // map[fun:1 go:3 great:1 is:2]
```

- `count[word]++` works because missing keys return 0, so `0++` = 1

### Checking for duplicates

```go
func hasDuplicates(numbers []int) bool {
    seen := make(map[int]bool)
    for _, num := range numbers {
        if seen[num] {
            return true
        }
        seen[num] = true
    }
    return false
}
```

### Grouping data

```go
students := map[string][]string{
    "classA": {"Ali", "Mohamed"},
    "classB": {"Sherif", "Ahmed"},
}

// Add a student to classA
students["classA"] = append(students["classA"], "Omar")
```

### Filtering a slice

```go
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

var evens []int
for _, n := range numbers {
    if n%2 == 0 {
        evens = append(evens, n)
    }
}
fmt.Println(evens)  // [2 4 6 8 10]
```

---

## Part 6: nil Slices and Maps

### nil slice

```go
var s []int
fmt.Println(s == nil)   // true
fmt.Println(len(s))     // 0
s = append(s, 1, 2, 3)  // Works fine!
fmt.Println(s)           // [1 2 3]
```

- A `nil` slice has length 0 and capacity 0
- You can still `append` to it — Go handles it

### nil map

```go
var m map[string]int
fmt.Println(m == nil)  // true
fmt.Println(len(m))    // 0

// Reading from nil map is OK
fmt.Println(m["key"])  // 0

// WRITING to nil map PANICS!
// m["key"] = 1  // RUNTIME ERROR: assignment to entry in nil map

// Always initialize before writing
m = make(map[string]int)
m["key"] = 1  // Now it works
```

**Rule**: Always use `make()` or a map literal before writing to a map.

---

## How to Think About Collections

1. **Need a fixed-size list?** Use an array (rare)
2. **Need a dynamic list?** Use a slice (most common)
3. **Need key-value lookup?** Use a map
4. **Know the size in advance?** Use `make([]T, 0, size)` for better performance
5. **Need to check if element exists?** Use `map[T]bool`

### Common Mistakes to Avoid

```go
// WRONG — appending without reassigning
names := []string{"Ali"}
append(names, "Mohamed")  // Result is lost!
// CORRECT
names = append(names, "Mohamed")

// WRONG — writing to nil map
var m map[string]int
m["key"] = 1  // PANIC!
// CORRECT
m := make(map[string]int)
m["key"] = 1

// WRONG — assuming map order
// Maps are UNORDERED. Don't rely on iteration order.

// WRONG — modifying slice while iterating
// Use index-based loop or build a new slice instead
```

---

## Exercises

Open `main.go` in this folder and try these:

### Exercise 1: Slice Operations
Create a slice of your favorite foods. Add 2 more with `append`. Remove the second one. Print the result.

### Exercise 2: Sum and Average
Create a function `average(numbers []float64) float64` that returns the average of a slice of numbers.

### Exercise 3: Word Counter
Given a slice of words, use a map to count how many times each word appears. Print the counts.

### Exercise 4: Student Grades
Create a `map[string]int` for student names and grades. Write a function that:
- Returns the student with the highest grade
- Returns the average grade

### Exercise 5: Remove Duplicates
Write a function `unique(numbers []int) []int` that returns a new slice with duplicates removed.

### How to run:

```bash
docker compose run --rm go go run 01-basics/arrays-slices-maps/main.go
```
