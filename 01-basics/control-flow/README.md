# Control Flow | if, switch, for

---

## What is Control Flow?

Control flow is how you tell your program to make decisions and repeat actions.

Without control flow, your code runs top to bottom, every line, every time. With control flow, your code can:
- **Choose** which code to run (if/switch)
- **Repeat** code multiple times (for)
- **Skip** or **stop** repeating (break/continue)

Think of it like driving: you don't just go straight forever - you turn left, turn right, go around roundabouts, and sometimes make U-turns.

---

## Real World Usage

| Pattern | Where companies use it |
|---------|----------------------|
| `if/else` | Check if user is authenticated, validate input, check permissions |
| `switch` | Route HTTP methods (GET/POST/PUT), handle different error types |
| `for` loop | Process all items in a database query, read lines from a file |
| `for range` | Iterate over API responses, process JSON arrays |
| `break` | Stop searching when you find the result |
| `continue` | Skip invalid records while processing data |

---

## Part 1: if / else

### Basic if

```go
age := 28

if age >= 18 {
    fmt.Println("You are an adult")
}
```

- **What**: Runs the code inside `{}` ONLY if the condition is `true`
- **Note**: No parentheses `()` around the condition! This is different from C, Java, JavaScript
- **Note**: The `{` MUST be on the same line as `if`. Go enforces this.

### if / else

```go
age := 15

if age >= 18 {
    fmt.Println("You are an adult")
} else {
    fmt.Println("You are a minor")
}
```

- `else` runs when the condition is `false`
- `else` MUST be on the same line as the closing `}`

### if / else if / else

```go
score := 85

if score >= 90 {
    fmt.Println("Grade: A")
} else if score >= 80 {
    fmt.Println("Grade: B")
} else if score >= 70 {
    fmt.Println("Grade: C")
} else {
    fmt.Println("Grade: F")
}
```

- Go checks conditions from top to bottom
- It runs the FIRST one that is `true` and skips the rest
- `else` at the end is the "catch all" - runs if nothing else matched

### Comparison Operators

```go
a == b    // equal
a != b    // not equal
a > b     // greater than
a < b     // less than
a >= b    // greater than or equal
a <= b    // less than or equal
```

### Logical Operators

```go
&&    // AND - both must be true
||    // OR  - at least one must be true
!     // NOT - flips true to false
```

```go
age := 25
hasID := true

if age >= 18 && hasID {
    fmt.Println("You can enter")
}

if age < 13 || age > 65 {
    fmt.Println("You get a discount")
}

if !hasID {
    fmt.Println("No ID found")
}
```

---

## Part 2: if with Short Statement (Go Special!)

This is a feature unique to Go - you can declare a variable INSIDE the if statement:

```go
if age := 28; age >= 18 {
    fmt.Println("Adult, age:", age)
}
// age does NOT exist here - it only lives inside the if block
```

- The variable `age` is declared and checked in one line
- `age` ONLY exists inside the `if/else` block (not outside)
- This is very common in Go for error checking

**Real world example - error checking:**

```go
if err := doSomething(); err != nil {
    fmt.Println("Error:", err)
    return
}
// if no error, continue normally
```

- `err != nil` means "if there IS an error"
- `nil` in Go means "nothing" or "no value" (like `null` in other languages)
- You will see this pattern EVERYWHERE in Go code

**Why is this useful?**

It keeps the variable scoped to where it's needed. The `err` variable doesn't leak into the rest of your function.

---

## Part 3: switch

Switch is a cleaner way to write multiple if/else statements:

### Basic switch

```go
day := "Monday"

switch day {
case "Monday":
    fmt.Println("Start of work week")
case "Friday":
    fmt.Println("Almost weekend!")
case "Saturday", "Sunday":
    fmt.Println("Weekend!")
default:
    fmt.Println("Regular day")
}
```

- **No `break` needed!** Unlike C/Java/JavaScript, Go automatically breaks after each case
- You can have multiple values in one case: `"Saturday", "Sunday"`
- `default` runs if no case matches (like `else`)

### Switch with no condition (like if/else chain)

```go
score := 85

switch {
case score >= 90:
    fmt.Println("Grade: A")
case score >= 80:
    fmt.Println("Grade: B")
case score >= 70:
    fmt.Println("Grade: C")
default:
    fmt.Println("Grade: F")
}
```

- When you write `switch` with no value, each case is a boolean condition
- This is cleaner than writing many `if/else if`

### Switch with short statement

```go
import "runtime"  // You MUST import this package first!

switch os := runtime.GOOS; os {
case "linux":
    fmt.Println("Linux")
case "darwin":
    fmt.Println("macOS")
case "windows":
    fmt.Println("Windows")
default:
    fmt.Println("Unknown:", os)
}
```

- Just like `if`, you can declare a variable in the switch statement
- **IMPORTANT**: You must `import "runtime"` at the top of your file to use `runtime.GOOS`

### What is runtime.GOOS?

`runtime` is a built-in Go package that gives you information about the system your code is running on.

`GOOS` is a variable inside `runtime` that tells you the **operating system**.

```go
import "runtime"

fmt.Println(runtime.GOOS)    // "linux", "darwin", "windows", etc.
```

**Other useful things in runtime:**

```go
runtime.GOOS       // Operating system: "linux", "darwin", "windows"
runtime.GOARCH     // CPU architecture: "amd64", "arm64"
runtime.NumCPU()   // Number of CPUs: 8
runtime.Version()  // Go version: "go1.22.0"
```

**Why is this useful?** When your program runs on different operating systems and needs to behave differently:

```go
import "runtime"

switch runtime.GOOS {
case "linux":
    configPath = "/etc/myapp/config.yml"
case "windows":
    configPath = "C:\\ProgramData\\myapp\\config.yml"
case "darwin":
    configPath = "/Library/Application Support/myapp/config.yml"
}
```

If you run it in our Docker setup, it will print `"linux"` because the container runs Linux.

### fallthrough

```go
switch 3 {
case 3:
    fmt.Println("Three")
    fallthrough
case 4:
    fmt.Println("Four (runs because of fallthrough)")
case 5:
    fmt.Println("Five (does NOT run)")
}
// Output:
// Three
// Four (runs because of fallthrough)
```

- `fallthrough` forces the next case to run regardless of its condition
- Rarely used in practice, but good to know it exists

**Real world - HTTP method routing:**

```go
switch r.Method {
case "GET":
    handleGet(w, r)
case "POST":
    handlePost(w, r)
case "PUT":
    handlePut(w, r)
case "DELETE":
    handleDelete(w, r)
default:
    http.Error(w, "Method not allowed", 405)
}
```

### Understanding the HTTP Routing Example

This example shows how Go handles **HTTP requests** in a web server. Let's break it down:

**What imports do you need?**

```go
import "net/http"
```

`net/http` is Go's built-in package for building web servers and making HTTP requests. You do NOT need to install anything external.

**What are `r`, `w`, and `r.Method`?**

```go
func handleRequest(w http.ResponseWriter, r *http.Request) {
    // w = ResponseWriter - you write your response (HTML, JSON, etc.) to this
    // r = Request - contains everything about the incoming request
    // r.Method = the HTTP method: "GET", "POST", "PUT", "DELETE"
}
```

- `r` is `*http.Request` — it holds all the information about the request (URL, headers, body, method)
- `w` is `http.ResponseWriter` — it's how you send a response back to the client (browser, API consumer)
- `r.Method` is a string that tells you what kind of request it is

**What are HTTP methods?**

| Method | What it does | Example |
|--------|-------------|---------|
| `GET` | Read/fetch data | Get a list of users |
| `POST` | Create new data | Create a new user |
| `PUT` | Update existing data | Update user's email |
| `DELETE` | Delete data | Delete a user |

**What are `handleGet`, `handlePost`, etc.?**

These are functions YOU define. For example:

```go
func handleGet(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "You sent a GET request")
}

func handlePost(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "You sent a POST request")
}
```

**Full working example:**

```go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        fmt.Fprintf(w, "Hello! You sent a GET request")
    case "POST":
        fmt.Fprintf(w, "You sent a POST request")
    default:
        http.Error(w, "Method not allowed", 405)
    }
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
```

- `http.HandleFunc("/", handler)` — when someone visits `/`, call the `handler` function
- `http.ListenAndServe(":8080", nil)` — start the server on port 8080
- `fmt.Fprintf(w, ...)` — write the response back to the client
- `http.Error(w, message, statusCode)` — send an error response

**You don't need to understand all of this now.** We'll cover `net/http` in detail in the standard library section (06-standard-library). This is just a preview of how `switch` is used in real Go applications.

---

## Part 4: for Loop (The Only Loop in Go!)

Go has ONLY one looping keyword: `for`. But it can do everything that `while`, `do-while`, and `for` do in other languages.

### Classic for loop (like C/Java)

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
// Output: 0, 1, 2, 3, 4
```

- `i := 0` - initialization (runs once before the loop starts)
- `i < 5` - condition (checked before each iteration)
- `i++` - post statement (runs after each iteration)
- **No parentheses** around the three parts

**How to think about it:**

```
for START; CONDITION; AFTER_EACH {
    // code to repeat
}

for i := 0; i < 5; i++ {
    // runs 5 times: i=0, i=1, i=2, i=3, i=4
    // stops when i=5 because 5 < 5 is false
}
```

### Counting backwards

```go
for i := 10; i > 0; i-- {
    fmt.Println(i)
}
fmt.Println("Go!")
// Output: 10, 9, 8, ..., 1, Go!
```

### Stepping by more than 1

```go
for i := 0; i < 20; i += 5 {
    fmt.Println(i)
}
// Output: 0, 5, 10, 15
```

---

## Part 5: for as while

If you drop the init and post statements, `for` becomes a `while` loop:

```go
count := 0
for count < 5 {
    fmt.Println(count)
    count++
}
// Output: 0, 1, 2, 3, 4
```

- Only the condition remains
- This is exactly how `while` works in other languages
- Go just uses `for` for everything

### Infinite loop

```go
for {
    fmt.Println("This runs forever!")
    // use break to exit
}
```

- No condition = runs forever
- You MUST use `break` to exit, or the program never stops
- Common in servers that listen for connections

**Real world - server main loop:**

```go
for {
    conn, err := listener.Accept()
    if err != nil {
        log.Println("Error:", err)
        continue
    }
    go handleConnection(conn)
}
```

---

## Part 6: for range

`for range` iterates over collections (strings, arrays, slices, maps). We'll use it more when we learn about arrays and slices, but here's a preview:

### Iterating over a string

```go
name := "Go"
for index, char := range name {
    fmt.Printf("Index: %d, Char: %c\n", index, char)
}
// Output:
// Index: 0, Char: G
// Index: 1, Char: o
```

- `range` gives you two values: the **index** and the **value**
- If you don't need the index, use `_`:

```go
for _, char := range "Hello" {
    fmt.Printf("%c ", char)
}
// Output: H e l l o
```

### Iterating over numbers (Go 1.22+)

```go
for i := range 5 {
    fmt.Println(i)
}
// Output: 0, 1, 2, 3, 4
```

---

## Part 7: break and continue

### break - exit the loop immediately

```go
for i := 0; i < 100; i++ {
    if i == 5 {
        break
    }
    fmt.Println(i)
}
// Output: 0, 1, 2, 3, 4
// Loop stops at 5, never reaches 6-99
```

**Real world - search:**

```go
users := []string{"Ali", "Mohamed", "Sherif", "Ahmed"}
for _, user := range users {
    if user == "Sherif" {
        fmt.Println("Found Sherif!")
        break  // no need to keep looking
    }
}
```

### continue - skip this iteration, go to next

```go
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue  // skip even numbers
    }
    fmt.Println(i)
}
// Output: 1, 3, 5, 7, 9
```

- `continue` skips the rest of the current iteration
- The loop continues with the next value

**Real world - skip invalid data:**

```go
for _, record := range records {
    if record.Email == "" {
        continue  // skip records with no email
    }
    sendEmail(record.Email)
}
```

---

## Part 8: Nested Loops and Labels

### Nested loops

```go
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        fmt.Printf("(%d, %d) ", i, j)
    }
    fmt.Println()
}
// Output:
// (0, 0) (0, 1) (0, 2)
// (1, 0) (1, 1) (1, 2)
// (2, 0) (2, 1) (2, 2)
```

### Labels - break out of nested loops

```go
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i == 1 && j == 1 {
            break outer  // breaks BOTH loops
        }
        fmt.Printf("(%d, %d) ", i, j)
    }
}
// Output: (0, 0) (0, 1) (0, 2) (1, 0)
```

- Without `outer`, `break` would only exit the inner loop
- Labels let you specify WHICH loop to break out of
- Use sparingly - if you need labels often, your code might need restructuring

---

## How to Think About Control Flow

1. **Use `if`** for simple yes/no decisions (1-2 conditions)
2. **Use `switch`** when comparing one value against many options (3+ cases)
3. **Use `for i`** when you know how many times to loop
4. **Use `for condition`** (while) when you don't know how many times
5. **Use `for range`** when iterating over collections
6. **Use `break`** when you found what you're looking for
7. **Use `continue`** when you want to skip bad data

### Common Mistakes to Avoid

```go
// WRONG - parentheses around condition
if (age >= 18) {  // Don't do this! Go doesn't need ()
}

// WRONG - { on new line
if age >= 18
{  // ERROR! { must be on the same line as if
}

// WRONG - using while
while count < 5 {  // ERROR! Go has no "while" keyword
}

// CORRECT
for count < 5 {  // Use "for" instead of "while"
}
```

---

## Exercises

Open `main.go` in this folder and try these:

### Exercise 1: Grade Calculator
Ask for a score (hardcode it as a variable). Print the grade:
- 90+ = A
- 80+ = B
- 70+ = C
- Below 70 = F

### Exercise 2: FizzBuzz
Loop from 1 to 20. For each number:
- If divisible by 3, print "Fizz"
- If divisible by 5, print "Buzz"
- If divisible by both 3 and 5, print "FizzBuzz"
- Otherwise, print the number

### Exercise 3: Find a Number
Create a loop that searches for the number 7 in numbers 1-20. When found, print "Found!" and stop.

### Exercise 4: Day of Week
Use a switch to print what you do on each day of the week.

### How to run:

```bash
docker compose run --rm go go run 01-basics/control-flow/main.go
```
