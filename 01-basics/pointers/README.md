# Pointers

---

## What is a Pointer?

A pointer is a variable that stores the **memory address** of another variable. Instead of holding a value like `42` or `"hello"`, a pointer holds the **location** where that value lives in memory.

Think of it like a home address: the address isn't the house itself, but it tells you exactly where to find the house.

```
Regular variable:         Pointer variable:
+--------+                +--------+
|   42   |                | 0xc000 |----> points to where 42 lives
+--------+                +--------+
  age                       ptr
```

**Why do pointers matter?**

1. **Efficiency** — Instead of copying large data around, you pass a pointer (just an address)
2. **Shared access** — Multiple parts of your code can modify the same data
3. **Mutability** — Functions can modify the caller's variables (not just their own copy)

---

## Real World Usage

| Pattern | Where companies use it |
|---------|----------------------|
| Pointer receivers | Methods that modify struct fields (`func (u *User) UpdateName()`) |
| Passing large structs | Avoid copying big objects to functions |
| Optional values | `*string` can be `nil` (no value) or point to a string |
| Linked data structures | Linked lists, trees, graphs (`type Node struct { Next *Node }`) |
| Interface satisfaction | Many interfaces require pointer receivers |
| Database scanning | `sql.Scan` needs pointers to write values into your variables |
| JSON unmarshaling | `json.Unmarshal(data, &myStruct)` needs a pointer to fill your struct |
| Error handling | Functions modify caller's data through pointers |

---

## Part 1: Memory Addresses and the `&` Operator

Every variable lives somewhere in your computer's memory. The `&` operator gives you that address.

```go
name := "Sherif"
age := 25

fmt.Println("Value of name:", name)       // Sherif
fmt.Println("Address of name:", &name)    // 0xc0000101e0 (some hex address)

fmt.Println("Value of age:", age)         // 25
fmt.Println("Address of age:", &age)      // 0xc0000b2008 (different address)
```

- `&name` = "give me the address where `name` is stored"
- The address is a hexadecimal number like `0xc0000101e0`
- Every variable has a **unique** address in memory

```
Memory (simplified):
Address:     0xc000   0xc004   0xc008   0xc00c   0xc010
           +--------+--------+--------+--------+--------+
           |   25   |  ....  |  ....  | "Sher" | "if"   |
           +--------+--------+--------+--------+--------+
              age                         name
```

---

## Part 2: Creating Pointers and Dereferencing with `*`

### Declaring a pointer

```go
age := 25
var ptr *int    // ptr is a "pointer to int"
ptr = &age      // ptr now holds the address of age

fmt.Println("age value:", age)      // 25
fmt.Println("age address:", &age)   // 0xc0000b2008
fmt.Println("ptr value:", ptr)      // 0xc0000b2008 (same address!)
fmt.Println("ptr points to:", *ptr) // 25 (the value at that address)
```

- `*int` = "this variable will hold the address of an int"
- `&age` = "give me the address of age"
- `*ptr` = "give me the value at the address ptr holds" (this is called **dereferencing**)

### The two meanings of `*`

This is the most confusing part for beginners. The `*` symbol has **two different meanings**:

```go
// Meaning 1: In a TYPE — declares a pointer type
var ptr *int          // "ptr is a pointer to an int"

// Meaning 2: In an EXPRESSION — dereferences (follows the pointer)
value := *ptr         // "give me what ptr points to"
*ptr = 100            // "set the value at ptr's address to 100"
```

**Rule of thumb:**
- `*` before a **type** (like `*int`, `*string`) = "pointer to"
- `*` before a **variable** (like `*ptr`) = "value at"

### Full example with ASCII diagram

```go
age := 25
ptr := &age

fmt.Println(age)   // 25
fmt.Println(ptr)   // 0xc0000b2008
fmt.Println(*ptr)  // 25

*ptr = 30          // Change the value THROUGH the pointer

fmt.Println(age)   // 30  <-- age changed!
fmt.Println(*ptr)  // 30
```

```
BEFORE *ptr = 30:

  age             ptr
+------+        +----------+
|  25  |  <---- | 0xc00008 |
+------+        +----------+
0xc00008


AFTER *ptr = 30:

  age             ptr
+------+        +----------+
|  30  |  <---- | 0xc00008 |
+------+        +----------+
0xc00008
```

The pointer didn't change — it still points to the same address. But the **value at that address** changed from 25 to 30. Since `age` lives at that address, `age` is now 30.

### Shorthand declaration

```go
// Long way
var ptr *int
age := 25
ptr = &age

// Short way (most common)
age := 25
ptr := &age

// Create with new() — allocates memory and returns a pointer
ptr2 := new(int)      // *int pointing to 0 (zero value)
*ptr2 = 42
fmt.Println(*ptr2)    // 42
```

---

## Part 3: Pointer to Pointer

Yes, you can have a pointer that points to another pointer!

```go
value := 42
ptr := &value      // pointer to value
pptr := &ptr       // pointer to pointer

fmt.Println("value:", value)     // 42
fmt.Println("*ptr:", *ptr)       // 42
fmt.Println("**pptr:", **pptr)   // 42

**pptr = 100
fmt.Println("value:", value)     // 100
```

```
  value           ptr              pptr
+------+       +----------+     +----------+
|  42  |  <--- | 0xc00008 | <-- | 0xc00010 |
+------+       +----------+     +----------+
0xc00008        0xc00010         0xc00018

**pptr means: follow pptr to ptr, then follow ptr to value
```

- `*int` = pointer to int
- `**int` = pointer to pointer to int
- `**pptr` = follow two levels of pointers to get the final value

**In practice**, you rarely need pointer to pointer. But it's good to understand the concept.

---

## Part 4: Pointers with Functions (Pass by Value vs Pass by Reference)

This is the **most important** reason pointers exist in Go.

### Without pointers — pass by value (copy)

```go
func doubleValue(n int) {
    n = n * 2
    fmt.Println("Inside function:", n)  // 50
}

func main() {
    num := 25
    doubleValue(num)
    fmt.Println("After function:", num)  // 25 — UNCHANGED!
}
```

```
main():          doubleValue(n):
  num              n (COPY of num)
+------+         +------+
|  25  |         |  25  | --> becomes 50
+------+         +------+
  (unchanged)      (separate variable, thrown away when function ends)
```

Go is **pass by value** — when you pass `num` to a function, Go makes a **copy**. The function works on the copy, not the original.

### With pointers — pass by reference

```go
func doubleValue(n *int) {
    *n = *n * 2
    fmt.Println("Inside function:", *n)  // 50
}

func main() {
    num := 25
    doubleValue(&num)
    fmt.Println("After function:", num)  // 50 — CHANGED!
}
```

```
main():          doubleValue(n):
  num              n (pointer to num)
+------+         +----------+
|  25  | <------ | 0xc00008 |
+------+         +----------+
  |
  v (after *n = *n * 2)
+------+
|  50  |
+------+
```

- `func doubleValue(n *int)` = "I accept a pointer to an int"
- `doubleValue(&num)` = "here's the address of num"
- `*n = *n * 2` = "take the value at n's address, double it, store it back"
- Now the **original** `num` is modified!

### Practical example — swapping two values

```go
// WITHOUT pointers — doesn't work
func swapBroken(a, b int) {
    a, b = b, a
    // Only swaps the COPIES, original values unchanged
}

// WITH pointers — works!
func swap(a, b *int) {
    *a, *b = *b, *a
}

func main() {
    x, y := 10, 20
    swap(&x, &y)
    fmt.Println(x, y)  // 20 10 — swapped!
}
```

### When to use pointers in functions

| Situation | Use pointer? | Why |
|-----------|-------------|-----|
| Function needs to **modify** the input | Yes | So changes affect the original |
| Passing a **large struct** (many fields) | Yes | Avoid expensive copy |
| Passing a **small value** (int, bool) | No | Copy is cheap, simpler code |
| Function only **reads** the input | Usually no | Unless the struct is very large |
| Return a value from a function | Either | Returning a pointer avoids a copy |

---

## Part 5: Pointers with Structs

Pointers are especially common with structs because structs can be large.

```go
type User struct {
    Name  string
    Email string
    Age   int
}

// Without pointer — works on a COPY
func birthdayValue(u User) {
    u.Age++
    // Original user is NOT modified
}

// With pointer — modifies the ORIGINAL
func birthdayPointer(u *User) {
    u.Age++    // Go automatically dereferences: (*u).Age++
}

func main() {
    user := User{Name: "Sherif", Email: "sherif@example.com", Age: 25}

    birthdayValue(user)
    fmt.Println(user.Age)  // 25 — unchanged

    birthdayPointer(&user)
    fmt.Println(user.Age)  // 26 — modified!
}
```

**Important shorthand**: Go automatically dereferences struct pointers. You write `u.Age` instead of `(*u).Age`. Both work, but `u.Age` is the convention.

```go
user := User{Name: "Sherif", Age: 25}
ptr := &user

// These are the same:
fmt.Println((*ptr).Name)  // Sherif (explicit dereference)
fmt.Println(ptr.Name)     // Sherif (Go does it for you)
```

### Creating struct pointers

```go
// Method 1: Address of existing struct
user := User{Name: "Ali", Age: 20}
ptr := &user

// Method 2: Address of literal (most common)
ptr2 := &User{Name: "Mohamed", Age: 28}

// Method 3: Using new (rare)
ptr3 := new(User)
ptr3.Name = "Omar"
ptr3.Age = 22
```

---

## Part 6: Pointers with Slices and Maps (They're Already References!)

Here's something important: **slices** and **maps** already behave like references. You usually **don't** need to pass them as pointers.

### Slices — already reference-like

```go
func addElement(s []int) {
    s[0] = 999   // This DOES modify the original!
}

func main() {
    numbers := []int{1, 2, 3}
    addElement(numbers)
    fmt.Println(numbers)  // [999 2 3] — modified!
}
```

Why? Because a slice is actually a small struct (pointer + length + capacity). When you pass a slice, the pointer to the underlying array is copied — so both the original and the copy point to the same data.

```
main():              addElement(s):
  numbers              s (copy of slice header)
+----------+         +----------+
| ptr  | 3 | 3 |    | ptr  | 3 | 3 |
+----------+         +----------+
    |                     |
    +------> same <-------+
           +---+---+---+
           | 1 | 2 | 3 |
           +---+---+---+
           Underlying Array
```

**But**: If the function uses `append` and the slice grows, the original won't see the new elements:

```go
func addToSlice(s []int) {
    s = append(s, 4)  // May create new underlying array
    fmt.Println("Inside:", s)  // [1 2 3 4]
}

func main() {
    numbers := []int{1, 2, 3}
    addToSlice(numbers)
    fmt.Println("Outside:", numbers)  // [1 2 3] — no 4!
}
```

To fix this, either return the new slice or pass a pointer to the slice:

```go
// Option 1: Return the new slice (preferred)
func addToSlice(s []int) []int {
    return append(s, 4)
}

// Option 2: Use a pointer to the slice (rare)
func addToSlice(s *[]int) {
    *s = append(*s, 4)
}
```

### Maps — already reference-like

```go
func addKey(m map[string]int) {
    m["new"] = 42   // This DOES modify the original!
}

func main() {
    data := map[string]int{"a": 1}
    addKey(data)
    fmt.Println(data)  // map[a:1 new:42] — modified!
}
```

Maps are internally pointers, so you almost **never** need `*map[...]`.

### Summary table

| Type | Passed by | Need pointer? |
|------|----------|--------------|
| `int`, `float64`, `bool`, `string` | Value (copy) | Yes, if you want to modify |
| `array` (`[5]int`) | Value (copy) | Yes, if you want to modify |
| `slice` (`[]int`) | Slice header (copy), but shares data | Rarely (only for append) |
| `map` | Reference-like | Almost never |
| `struct` | Value (copy) | Yes, for large structs or modification |

---

## Part 7: nil Pointers and Safety

A pointer that doesn't point to anything has the value `nil`.

```go
var ptr *int
fmt.Println(ptr)        // <nil>
fmt.Println(ptr == nil) // true

// DANGER: Dereferencing a nil pointer causes a PANIC!
// fmt.Println(*ptr)    // RUNTIME ERROR: invalid memory address
```

```
  ptr
+------+
| nil  |----> NOWHERE (crash if you follow this!)
+------+
```

### Always check for nil before dereferencing

```go
func printValue(ptr *int) {
    if ptr == nil {
        fmt.Println("Pointer is nil, nothing to print")
        return
    }
    fmt.Println("Value:", *ptr)
}

func main() {
    var ptr *int
    printValue(ptr)    // Pointer is nil, nothing to print

    value := 42
    printValue(&value) // Value: 42
}
```

### nil pointers in real-world code

```go
type Config struct {
    MaxRetries *int    // nil means "use default"
    Timeout    *int    // nil means "use default"
    Debug      *bool   // nil means "use default"
}

func getMaxRetries(c Config) int {
    if c.MaxRetries == nil {
        return 3  // default
    }
    return *c.MaxRetries
}

func main() {
    // Only set what you need — everything else uses defaults
    retries := 5
    config := Config{
        MaxRetries: &retries,
        // Timeout and Debug are nil — will use defaults
    }
    fmt.Println(getMaxRetries(config))  // 5
}
```

This is a common pattern for **optional fields** — `*int` can be `nil` (not set) or point to a value. A plain `int` is always 0, so you can't tell "not set" from "set to 0".

---

## Part 8: Common Patterns and Best Practices

### Pattern 1: Constructor functions returning pointers

```go
type Server struct {
    Host string
    Port int
}

func NewServer(host string, port int) *Server {
    return &Server{
        Host: host,
        Port: port,
    }
}

func main() {
    s := NewServer("localhost", 8080)
    fmt.Println(s.Host, s.Port)  // localhost 8080
}
```

This is Go's convention for constructors — a function named `NewXxx` that returns `*Xxx`.

### Pattern 2: Method with pointer receiver

```go
type Counter struct {
    Value int
}

// Pointer receiver — modifies the original
func (c *Counter) Increment() {
    c.Value++
}

// Value receiver — works on a copy
func (c Counter) GetValue() int {
    return c.Value
}

func main() {
    counter := Counter{Value: 0}
    counter.Increment()
    counter.Increment()
    counter.Increment()
    fmt.Println(counter.GetValue())  // 3
}
```

### Pattern 3: Returning a pointer to indicate "not found"

```go
func findUser(id int) *User {
    users := map[int]User{
        1: {Name: "Ali", Email: "ali@test.com", Age: 25},
        2: {Name: "Mohamed", Email: "mohamed@test.com", Age: 28},
    }

    if user, ok := users[id]; ok {
        return &user
    }
    return nil  // Not found
}

func main() {
    user := findUser(1)
    if user != nil {
        fmt.Println("Found:", user.Name)
    }

    user = findUser(99)
    if user == nil {
        fmt.Println("User not found")
    }
}
```

---

## Part 9: Common Mistakes to Avoid

### Mistake 1: Forgetting to dereference

```go
// WRONG
func double(n *int) {
    n = n * 2   // ERROR: can't multiply an address!
}

// CORRECT
func double(n *int) {
    *n = *n * 2  // Dereference to get/set the value
}
```

### Mistake 2: Dereferencing a nil pointer

```go
// WRONG — will crash
var ptr *int
fmt.Println(*ptr)  // PANIC!

// CORRECT — check first
var ptr *int
if ptr != nil {
    fmt.Println(*ptr)
}
```

### Mistake 3: Returning address of loop variable

```go
// WRONG — all pointers point to the same variable!
func getPointers() []*int {
    var ptrs []*int
    for i := 0; i < 3; i++ {
        ptrs = append(ptrs, &i)  // All point to the same 'i'!
    }
    return ptrs
}
// Result: all three pointers have value 3 (the final value of i)

// CORRECT — create a new variable each iteration
func getPointers() []*int {
    var ptrs []*int
    for i := 0; i < 3; i++ {
        v := i  // New variable each iteration
        ptrs = append(ptrs, &v)
    }
    return ptrs
}
```

### Mistake 4: Unnecessary pointers

```go
// WRONG — overusing pointers for small types
func add(a *int, b *int) *int {
    result := *a + *b
    return &result
}

// CORRECT — just use values for simple operations
func add(a, b int) int {
    return a + b
}
```

### Mistake 5: Confusing `*` in types vs expressions

```go
var p *int     // DECLARATION: p is a pointer to int
x := *p        // EXPRESSION: get the value p points to

// Tip: if * is next to a TYPE name (int, string, MyStruct), it's declaring a pointer
// If * is next to a VARIABLE name, it's dereferencing
```

---

## How to Think About Pointers

```
Decision flowchart:

Does your function need to MODIFY the input?
├── Yes → Use a pointer (*Type)
└── No → Is the data LARGE (big struct)?
    ├── Yes → Use a pointer (avoid expensive copy)
    └── No → Use a value (simpler, safer)

Is it a slice or map?
├── Yes → Don't use a pointer (already reference-like)
└── No → Apply the rules above
```

**Key rules:**
1. `&variable` = "give me the address of this variable"
2. `*pointer` = "give me the value at this address"
3. `*Type` in a declaration = "this is a pointer to Type"
4. Always check for `nil` before dereferencing
5. Slices and maps already act like references — don't over-pointer them
6. When in doubt, start without pointers and add them when needed

---

## Exercises

Open `main.go` in this folder and try these:

### Exercise 1: Basic Pointer Operations
Create an `int` variable with value `10`. Create a pointer to it. Use the pointer to change the value to `20`. Print both the variable and the dereferenced pointer to confirm they both show `20`.

### Exercise 2: Swap Function
Write a function `swap(a, b *int)` that swaps the values of two integers using pointers. Test it from main with two variables.

### Exercise 3: Modify Struct via Pointer
Create a `Student` struct with `Name` (string) and `Grade` (int). Write a function `promote(s *Student)` that increases the grade by 10. Call it from main and verify the original struct changed.

### Exercise 4: Safe Dereference
Write a function `safeDeref(ptr *int) int` that returns the value the pointer points to, or returns `0` if the pointer is `nil`. Test it with both a valid pointer and a nil pointer.

### Exercise 5: Pointer to Slice Append
Write a function `addItem(items *[]string, item string)` that appends an item to a slice using a pointer. Call it from main, add 3 items, and print the result. Then rewrite it to return the slice instead (without a pointer) and compare both approaches.

### How to run:

```bash
docker compose run --rm go go run 01-basics/pointers/main.go
```
