# Variables, Types & Constants

---

## Before You Start - What is Go and Why Companies Use It

### What is Go?

Go (or Golang) is a programming language created by Google in 2009.

The people who created it are legends: Robert Griesemer, Rob Pike, and Ken Thompson (the same person who created Unix and C language).

### Why did they create it?

Google had a problem: they had MILLIONS of lines of C++ and Java code that was:

- Slow to compile (takes 45+ minutes)
- Hard to read and maintain
- Complex concurrency (running multiple things at once)

So they created Go to be: **simple, fast, and great at concurrency**.

### Who uses Go?

| Company       | What they use it for       |
| ------------- | -------------------------- |
| Google        | Kubernetes, internal tools |
| Uber          | Microservices, geofence    |
| Netflix       | Server architecture        |
| Docker        | The entire Docker platform |
| Kubernetes    | The entire K8s system      |
| Twitch        | Video streaming backend    |
| Cloudflare    | Network services           |
| Mercado Libre | E-commerce backend         |

### Why companies love Go

1. **Fast compilation** - seconds not minutes
2. **Single binary** - no dependencies to install on servers
3. **Built-in concurrency** - handles millions of requests
4. **Simple language** - easy to hire and train developers
5. **Great standard library** - less need for external packages

---

## Go & Memory - How It Manages Memory

### What is Memory?

Think of your computer's RAM as a giant shelf with numbered slots.

When you create a variable, Go reserves a slot on that shelf for your data.

### Stack vs Heap

Go uses two areas in memory:

```
STACK (fast, automatic)              HEAP (slower, managed by GC)
+---------------------------+        +---------------------------+
| func main() {             |        |                           |
|   age := 25         <-- here       |  data that lives longer   |
|   name := "Sherif"  <-- here       |  or is shared between     |
| }                         |        |  functions   <-- here     |
| automatically cleaned up  |        |  cleaned by Garbage       |
| when function ends        |        |  Collector                |
+---------------------------+        +---------------------------+
```

**Stack** (the fast one):

- Small, fast memory
- Variables inside functions go here
- Automatically freed when function ends
- Like a stack of plates - last in, first out

**Heap** (the slower one):

- Bigger, slower memory
- Data that needs to live longer goes here
- Cleaned by the **Garbage Collector (GC)**

### Garbage Collector

In C/C++, YOU have to free memory manually. Forget to do it = memory leak.

In Go, the **Garbage Collector** does it for you automatically.

```
You:    "I don't need this variable anymore"
Go GC:  "I'll clean it up for you, don't worry"
```

**This is why Go is popular** - you get the speed close to C, but the safety of languages like Java/Python.

### How much memory does each type use?

```
int8    = 1 byte   (8 bits)    - stores: -128 to 127
int16   = 2 bytes  (16 bits)   - stores: -32,768 to 32,767
int32   = 4 bytes  (32 bits)   - stores: ~2 billion
int64   = 8 bytes  (64 bits)   - stores: very large numbers
float32 = 4 bytes  (32 bits)
float64 = 8 bytes  (64 bits)
bool    = 1 byte   (8 bits)    - stores: true or false
string  = 16 bytes (header)    - pointer + length
```

**Why does this matter?**

If you're storing 1 million ages (0-120), using `int8` instead of `int64` saves:

```
int64: 1,000,000 x 8 bytes = 8 MB
int8:  1,000,000 x 1 byte  = 1 MB
Saved: 7 MB
```

For most cases, just use `int` and `float64`. Only optimize when you have a LOT of data.

---

## Project Structure - How to Organize a Go Project

### Simplest structure (what we're doing now)

```
my-project/
  go.mod          <-- project definition
  main.go         <-- entry point
```

### Small project

```
my-api/
  go.mod
  go.sum           <-- dependency lock file
  main.go
  handlers/
    user.go
    product.go
  models/
    user.go
    product.go
```

### Production project (what companies use)

```
my-service/
  cmd/
    api/
      main.go           <-- entry point for the API server
    worker/
      main.go           <-- entry point for background worker
  internal/
    handler/             <-- HTTP handlers
      user.go
      product.go
    service/             <-- business logic
      user.go
      product.go
    repository/          <-- database access
      user.go
      product.go
    model/               <-- data structures
      user.go
      product.go
  pkg/                   <-- shared code other projects can use
    logger/
      logger.go
  config/
    config.go
  go.mod
  go.sum
  Dockerfile
  README.md
```

**Key folders explained:**

- `cmd/` - entry points. Each subfolder = one executable
- `internal/` - private code. Go PREVENTS other projects from importing this
- `pkg/` - public shared code
- `handler/` - receives HTTP requests
- `service/` - business logic (the real work)
- `repository/` - talks to the database

---

## How to Start Any Go Project

### Step 1: Initialize the module

```bash
mkdir my-project
cd my-project
go mod init my-project
```

This creates `go.mod` which tells Go: "this is a project called my-project".

### Step 2: Create main.go

Every Go program MUST have:

```go
package main        // 1. this is the main package

import "fmt"        // 2. import what you need

func main() {       // 3. the entry point
    fmt.Println("Hello, World!")
}
```

### Step 3: Run it

```bash
go run main.go
```

Or with Docker in our setup:

```bash
docker compose run --rm go go run main.go
```

---

## Variables - Line by Line

### The Full Program Structure

```go
package main
```

- **What**: Declares this file belongs to the `main` package
- **Why**: Go requires every file to belong to a package. The `main` package is special - it's the one that gets executed

```go
import (
    "fmt"
    "strconv"
)
```

- **What**: Import packages we need
- **Why**: `fmt` for printing, `strconv` for string-to-number conversion
- **Note**: The parentheses `()` let you import multiple packages at once

```go
func main() {
```

- **What**: The main function - where execution starts
- **Why**: Go looks for this function to start the program. No `main()` = program won't run

---

### Part 1: Declaring Variables

```go
var name string = "mohamed Elsherif"
```

- `var` = keyword to declare a variable
- `name` = the variable name
- `string` = the type (text)
- `= "mohamed Elsherif"` = the initial value
- **Memory**: Go allocates 16 bytes (string header) on the stack, pointing to " "mohamed Elsherif" data

```go
var age int = 25
```

- Same pattern but with `int` type (whole number)
- **Memory**: 8 bytes on the stack (on 64-bit system)

```go
var city = "Cairo"
```

- **No type specified!** Go infers it's a `string` from the value `"Cairo"`
- This is called **type inference**

```go
country := "Egypt"
isStudent := true
```

- **Short declaration** with `:=` - the most common way in Go
- `:=` = declare + assign in one step
- `country` is inferred as `string`, `isStudent` is inferred as `bool`
- **RULE**: `:=` only works inside functions, NOT at package level

```go
x, y, z := 1, 2, 3
```

- **Multiple variables in one line**
- `x=1`, `y=2`, `z=3` - all are `int` type
- Useful when you want related variables together

```go
var (
    firstName string = "Elsherif"
    lastName  string = "mohamed"
    score     int    = 95
)
```

- **var block** - group multiple declarations together
- Cleaner than writing `var` on each line
- Common at the **package level** for global variables

---

### Part 2: Printing

```go
fmt.Println("Name:", name)
```

- `fmt` = the format package
- `Println` = Print Line (prints and adds a new line at the end)
- `"Name:"` = literal text
- `name` = the variable value
- **Output**: `Name: Sherif`

```go
fmt.Printf("int: %d\n", defaultInt)
```

- `Printf` = Print Formatted
- `%d` = placeholder for integer
- `%f` = placeholder for float
- `%s` = placeholder for string
- `%t` = placeholder for boolean
- `%q` = placeholder for quoted string
- `%c` = placeholder for character
- `\n` = new line (Printf doesn't add one automatically)

---

### Part 3: Zero Values

```go
var defaultInt int
var defaultFloat float64
var defaultBool bool
var defaultString string
```

- **No value assigned!** But Go NEVER leaves a variable uninitialized
- Every type has a **zero value**:
  - `int` -> `0`
  - `float64` -> `0.0`
  - `bool` -> `false`
  - `string` -> `""` (empty string)

**Why is this important?**

In C, an uninitialized variable contains garbage data (whatever was in that memory slot before).
In Go, you are GUARANTEED it starts at zero. This prevents entire categories of bugs.

---

### Part 4: Data Types in Action

```go
var small int8 = 127
```

- `int8` = 1 byte = values from -128 to 127
- `127` is the MAXIMUM value for int8. If you try `128`, Go will give an error at compile time

```go
var big int64 = 9999999999
```

- `int64` = 8 bytes = very large numbers
- Use when `int` isn't big enough

```go
var positive uint = 42
```

- `uint` = **unsigned** integer = positive numbers only (0 and above)
- Cannot store negative numbers
- Used for things that can't be negative: array index, count, size

```go
var price float64 = 19.99
fmt.Printf("Price: %.2f\n", price)
```

- `float64` = decimal number with high precision
- `%.2f` = show only 2 decimal places
- **Always use `float64`** unless you have a specific reason for `float32`

```go
greeting := "Hello, " + name
```

- `+` operator concatenates (joins) strings
- `"Hello, " + "Sherif"` = `"Hello, Sherif"`

```go
fmt.Println("Name length:", len(name))
```

- `len()` returns the number of **bytes** in a string (not characters!)
- For English text, bytes = characters. For Arabic, one character = 2-4 bytes

```go
var b byte = 'A'
var r rune = 'G'
```

- `byte` = alias for `uint8`. Stores one ASCII character (English, numbers, symbols)
- `rune` = alias for `int32`. Stores one Unicode character (any language, emoji)
- Single quotes `'A'` = one character. Double quotes `"A"` = a string

---

### Part 5: Type Conversion

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

- Go does **NOT** convert types automatically. This is intentional.
- You MUST explicitly convert: `float64(i)` converts `int` to `float64`

**Why no automatic conversion?**

```
// In JavaScript (automatic - dangerous):
"5" + 3 = "53"     // string! not 8
"5" - 3 = 2        // number! inconsistent

// In Go (explicit - safe):
// "5" + 3          // COMPILE ERROR. You must convert first.
```

```go
numStr := strconv.Itoa(42)          // int to string: 42 -> "42"
numInt, _ := strconv.Atoi("42")     // string to int: "42" -> 42
```

- `strconv` = string conversion package
- `Itoa` = Integer to ASCII (int -> string)
- `Atoi` = ASCII to Integer (string -> int)
- The `_` in `numInt, _` means "ignore the second return value" (which is an error)
- We'll learn about error handling later

---

### Part 6: Constants

```go
const Pi = 3.14159
const AppName = "MyApp"
```

- `const` = a value that CANNOT change after declaration
- If you try `Pi = 3.14`, Go gives a compile error
- **Use for**: mathematical constants, app config, status codes

```go
const (
    Sunday = iota    // 0
    Monday           // 1
    Tuesday          // 2
    Wednesday        // 3
    Thursday         // 4
    Friday           // 5
    Saturday         // 6
)
```

- `iota` = a special Go counter that starts at 0 and increments by 1
- Only works inside `const ()` blocks
- You only write `= iota` on the first line. The rest auto-increment

**Real world use:**

```go
// HTTP Status codes
const (
    StatusOK       = 200
    StatusNotFound = 404
    StatusError    = 500
)

// User roles using iota
const (
    RoleGuest = iota  // 0
    RoleUser          // 1
    RoleAdmin         // 2
    RoleSuperAdmin    // 3
)
```

---

## Exercises

Open `main.go` in this folder and try writing the code yourself:

### Exercise 1: Personal Info

Create variables for your name, age, city, and whether you're a student. Print them all.

### Exercise 2: Type Conversion

Create a float price (19.99), convert it to int, and see what happens to the decimal part.

### Exercise 3: Constants

Create an `iota` block for the months of the year (January=1). Print March.

### Exercise 4: Zero Values

Declare variables of each type without values and print them to see the zero values.

### How to run:

```bash
docker compose run --rm go go run 01-basics/variables/main.go
```
