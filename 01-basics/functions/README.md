# Functions

---

## What is a Function?

A function is a reusable block of code that does a specific job. Instead of writing the same code over and over, you write it once in a function and call it whenever you need it.

Think of it like a recipe: you write the recipe once, and you can cook the dish anytime by following it.

---

## Real World Usage

| Pattern | Where companies use it |
|---------|----------------------|
| Basic functions | Break code into manageable pieces |
| Multiple return values | Return result + error (Go's error handling pattern) |
| Named return values | Make function signatures self-documenting |
| Variadic functions | `fmt.Println` accepts any number of arguments |
| Anonymous functions | HTTP handlers, goroutine callbacks |
| Closures | Middleware in web servers, counter generators |
| defer | Close files, database connections, unlock mutexes |
| First-class functions | Strategy pattern, middleware chains |

---

## Part 1: Basic Functions

### Function with no parameters and no return value

```go
func sayHello() {
    fmt.Println("Hello!")
}

func main() {
    sayHello()  // Output: Hello!
    sayHello()  // Output: Hello!
}
```

- `func` = keyword to declare a function
- `sayHello` = function name
- `()` = no parameters
- No return type = doesn't return anything
- You **call** the function by writing its name with `()`

### Function with parameters

```go
func greet(name string) {
    fmt.Println("Hello, " + name + "!")
}

func main() {
    greet("Sherif")   // Output: Hello, Sherif!
    greet("Mohamed")  // Output: Hello, Mohamed!
}
```

- `name string` = the function takes one parameter called `name` of type `string`
- When you call `greet("Sherif")`, the value `"Sherif"` gets assigned to `name`

### Function with multiple parameters

```go
func add(a int, b int) {
    fmt.Println(a + b)
}

func main() {
    add(5, 3)   // Output: 8
    add(10, 20) // Output: 30
}
```

**Shorthand** - if parameters have the same type, you can write it once:

```go
func add(a, b int) {
    fmt.Println(a + b)
}
```

`a, b int` means both `a` and `b` are `int`.

---

## Part 2: Return Values

### Function with a return value

```go
func add(a, b int) int {
    return a + b
}

func main() {
    result := add(5, 3)
    fmt.Println(result)  // Output: 8

    // Or use it directly
    fmt.Println(add(10, 20))  // Output: 30
}
```

- `int` after the `()` = the return type
- `return a + b` = sends the result back to whoever called the function
- The caller can store the result in a variable with `:=`

### Multiple return values (Go Special!)

This is one of Go's most important features. Most Go functions return **two values**: the result and an error.

```go
func divide(a, b float64) (float64, string) {
    if b == 0 {
        return 0, "cannot divide by zero"
    }
    return a / b, ""
}

func main() {
    result, err := divide(10, 3)
    fmt.Println(result, err)  // Output: 3.3333... (empty string)

    result, err = divide(10, 0)
    fmt.Println(result, err)  // Output: 0 cannot divide by zero
}
```

- `(float64, string)` = returns TWO values
- The caller MUST receive both: `result, err := divide(10, 3)`
- If you don't need one, use `_`: `result, _ := divide(10, 3)`

**Real world pattern - error handling:**

```go
import (
    "fmt"
    "strconv"
)

func main() {
    num, err := strconv.Atoi("42")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Number:", num)
}
```

- `strconv.Atoi` returns `(int, error)` — two values
- You check `err != nil` to see if something went wrong
- This pattern is EVERYWHERE in Go

---

## Part 3: Named Return Values

You can name the return values in the function signature:

```go
func divide(a, b float64) (result float64, err string) {
    if b == 0 {
        err = "cannot divide by zero"
        return
    }
    result = a / b
    return
}
```

- `(result float64, err string)` = named return values
- `result` and `err` are created as variables automatically (with zero values)
- `return` without values = returns whatever `result` and `err` currently hold
- This is called a **naked return**

**When to use named returns:**
- When the function has multiple return values and the meaning isn't obvious
- Makes the function signature self-documenting
- Don't overuse them - for simple functions, regular returns are clearer

---

## Part 4: Variadic Functions (Variable number of arguments)

A variadic function accepts **any number** of arguments:

```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2))           // Output: 3
    fmt.Println(sum(1, 2, 3, 4, 5)) // Output: 15
    fmt.Println(sum())               // Output: 0
}
```

- `numbers ...int` = "accept zero or more int values"
- Inside the function, `numbers` is a **slice** (like an array) of `int`
- `...` is called the **ellipsis** and must be the LAST parameter

**You already use variadic functions!**

```go
fmt.Println("Hello", "World", 42, true)
// Println accepts any number of arguments of any type
```

### Passing a slice to a variadic function

```go
nums := []int{1, 2, 3, 4, 5}
fmt.Println(sum(nums...))  // Output: 15
```

- `nums...` "unpacks" the slice into individual arguments

---

## Part 5: Anonymous Functions and Closures

### Anonymous function (function with no name)

```go
func main() {
    // Define and call immediately
    func() {
        fmt.Println("I'm anonymous!")
    }()

    // Store in a variable
    greet := func(name string) {
        fmt.Println("Hello, " + name)
    }

    greet("Sherif")  // Output: Hello, Sherif
}
```

- A function without a name
- Can be called immediately with `()`
- Can be stored in a variable and called later

### Closures (function that remembers its environment)

```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    c := counter()
    fmt.Println(c())  // Output: 1
    fmt.Println(c())  // Output: 2
    fmt.Println(c())  // Output: 3

    c2 := counter()   // New counter, starts from 0
    fmt.Println(c2()) // Output: 1
}
```

- `counter()` returns a **function**
- The returned function "remembers" the `count` variable even after `counter()` finishes
- Each call to `counter()` creates a **new** independent counter
- This is called a **closure** - a function that "closes over" its surrounding variables

**Why are closures useful?**

```go
func multiplier(factor int) func(int) int {
    return func(n int) int {
        return n * factor
    }
}

func main() {
    double := multiplier(2)
    triple := multiplier(3)

    fmt.Println(double(5))  // Output: 10
    fmt.Println(triple(5))  // Output: 15
}
```

You create specialized functions from a general one.

---

## Part 6: defer

`defer` schedules a function call to run **just before the current function returns**:

```go
func main() {
    fmt.Println("Start")
    defer fmt.Println("This runs last")
    fmt.Println("Middle")
    fmt.Println("End")
}
// Output:
// Start
// Middle
// End
// This runs last
```

- `defer` doesn't run the function immediately
- It waits until the surrounding function is about to return
- Then it runs the deferred call

### Multiple defers (LIFO - Last In First Out)

```go
func main() {
    defer fmt.Println("First defer")
    defer fmt.Println("Second defer")
    defer fmt.Println("Third defer")
}
// Output:
// Third defer
// Second defer
// First defer
```

- Multiple defers run in **reverse order** (like a stack of plates)
- Last one deferred = first one to run

### Why is defer useful?

**Closing files:**

```go
func readFile(path string) {
    file, err := os.Open(path)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()  // Will close no matter how the function exits

    // Read from file...
    // Even if an error happens here, the file will be closed
}
```

**Real world uses of defer:**
- Close files: `defer file.Close()`
- Close database connections: `defer db.Close()`
- Unlock mutexes: `defer mutex.Unlock()`
- Close HTTP response body: `defer resp.Body.Close()`

The beauty of `defer` is that you put the cleanup code RIGHT NEXT TO the creation code, so you never forget to clean up.

---

## Part 7: Functions as Values (First-Class Functions)

In Go, functions are **first-class citizens** - they can be stored in variables, passed as arguments, and returned from other functions.

### Passing a function as a parameter

```go
func apply(a, b int, operation func(int, int) int) int {
    return operation(a, b)
}

func main() {
    add := func(a, b int) int { return a + b }
    subtract := func(a, b int) int { return a - b }

    fmt.Println(apply(10, 5, add))      // Output: 15
    fmt.Println(apply(10, 5, subtract)) // Output: 5
}
```

- `operation func(int, int) int` = a parameter that accepts a function
- The function must take two `int` and return one `int`

---

## Part 8: init() Function

Go has a special function called `init()` that runs automatically **before** `main()`:

```go
var config string

func init() {
    config = "loaded"
    fmt.Println("init() ran first")
}

func main() {
    fmt.Println("main() ran second")
    fmt.Println("Config:", config)
}
// Output:
// init() ran first
// main() ran second
// Config: loaded
```

- `init()` runs automatically - you never call it yourself
- It runs BEFORE `main()`
- Used for setup: loading config, checking environment, initializing variables
- A file can have multiple `init()` functions (but this is rare)

---

## How to Think About Functions

1. **Keep functions small** - each function should do ONE thing
2. **Name them clearly** - `calculateTax` not `calc` or `doStuff`
3. **Always check errors** - if a function returns an error, check it
4. **Use defer for cleanup** - files, connections, locks
5. **Return early** - check for errors first, then do the main work

### Common Mistakes to Avoid

```go
// WRONG - ignoring errors
result, _ := strconv.Atoi(userInput)
// If userInput is "abc", result is 0 and you'd never know why!

// CORRECT - check the error
result, err := strconv.Atoi(userInput)
if err != nil {
    fmt.Println("Invalid input:", err)
    return
}

// WRONG - function does too many things
func handleEverything() { /* 200 lines of code */ }

// CORRECT - break it into smaller functions
func validateInput() error { ... }
func processData() (Result, error) { ... }
func saveResult(r Result) error { ... }
```

---

## Exercises

Open `main.go` in this folder and try these:

### Exercise 1: Calculator Functions
Create functions `add`, `subtract`, `multiply`, `divide` that take two `float64` and return `float64`. Call them all from main.

### Exercise 2: Multiple Returns
Create a function `checkAge(age int)` that returns `(string, bool)`:
- If age >= 18: return "adult", true
- If age < 18: return "minor", false

### Exercise 3: Variadic Sum
Create a function `sum(numbers ...int) int` that returns the sum of all numbers passed to it. Test it with different numbers of arguments.

### Exercise 4: Closure Counter
Create a `counter()` function that returns a function. Each time you call the returned function, it should return the next number (1, 2, 3...).

### Exercise 5: defer Practice
Write a function that prints "Start", defers "End", and prints "Middle". Predict the output before running it.

### How to run:

```bash
docker compose run --rm go go run 01-basics/functions/main.go
```
