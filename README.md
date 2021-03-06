# Learn Go

Learn how to use Go (Golang), an open-source programming language supported by Google!

## Why Learn Go?

Go (or Golang) is an open source programming language designed to build fast, reliable, and efficient software at scale.

Google uses Go specifically for its large networks of servers, and Go also powers much of Google’s own cloud platform. Developers use Go in application development, web development, in operations and infrastructure teams, and much more. It is the language of Cloud Native infrastructure and software development.

As Go’s popularity and adoption continue to increase, we’ll see even greater usage and more creative implementations of the language!

## Take-Away Skills:

In this course, we will cover:

- The setup of a Go environment and how to create a Go file.
- Go’s data types and variables.
- Go’s conditional statements.
- Using functions in Go.


## LESSON 1 Learn Go: Introduction

Get go ing by learning the basics of Google’s very own programming language, Go!

Exercises

- [x] 1. From the Get Go

- [x] 2. Compiling

- [x] 3. Running Files

- [x] 4. Basic Go Structure: Packages

- [x] 5. Basic Go Structure: main

- [x] 6. Importing Multiple Packages

- [x] 7. Comments

- [x] 8. Go Resources


## LESSON 2 Learn Go: Variables and Types

Learn all about the different types and variables in Go by building fundamental programming knowledge.

Exercises

- [x] 1. Values and Variables

- [x] 2. Literals

- [x] 3. Constants

- [x] 4. What is a Data Type?

![DataType](./src/DataType.png)

- [x] 5. Basic Numeric Types in Go

![NumericTypes](./src/NumericTypes.png)

- [x] 6. What is a Variable?

- [x] 7. Reading Go Errors

- [x] 8. Assigning Variables

- [x] 9. Strings

- [x] 10. Zero Values

- [x] 11. Inferring Type

- [x] 12. Default int Type

- [x] 13. Updating Variables

- [x] 14. Multiple Variable Declaration


## LESSON 3 Learn Go: fmt Package

Explore more of Go’s fmt package methods!

Exercises

- [x] 1. The fmt Package

- [x] 2. The Print Method

- [x] 3. The Printf Method

- [x] 4. Different Verbs

The format 'verbs' are derived from C's but are simpler. [their documentation](https://golang.org/pkg/fmt/#hdr-Printing)

- [x] 5. Sprint and Sprintln

- [x] 6. The Sprintf Method

- [x] 7. Getting User Input

## LESSON 4 Learn Go: Conditionals

If you want to learn how to include conditionals into your Go program, take this lesson! Else, you can use this lesson to review conditionals.

Exercises

- [x] 1. What are Conditionals?

Conditionals check on values and, depending on their status (if the values turn out to be true or false), executes an appropriate block of code in response.

- [x] 2. The if Statement

- [x] 3. The else Statement

- [x] 4. Comparison Operators

| Operator   | Meanig:     |
| ---------- | ----------  |
| ==         | Is equal to |
| !=         | Is NOT equal to   |

- [x] 5. Comparison Operators Continued

| Operator   | Meanig:     |
| ---------- | ----------  |
| <          | Less than   |
| >          | Greater than   |
| <=         | Less than or equal to   |
| >=         | Greater than or equal to   |


- [x] 6. Logical Operators (And, Or)

| Operator   | Meanig:     |
| ---------- | ----------  |
| &&         | And         |
| ||         | Or          |
| !          | Not         |

- [x] 7. Logical Operators (Not)

- [x] 8. The else if Statement

- [x] 9. The switch Statement

- [x] 10. Scoped Short Declaration Statement

- [x] 11. Randomizing

- [x] 12. Seeding

```go

package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  fmt.Println(rand.Intn(100))
}

```

## LESSON 5 Learn Go: Functions

Learn about how to create functions, reusable code that takes certain parameters and uses them to add modularity to your programs!

Exercises

- [x] 1. What is a Function?

In programming, a function is a block of code designed to be reused. As programmers, we want to find solutions to problems, but we also don’t want to do additional work when it’s not necessary. 

```go

func doubleNum(num int) int {
  return num * 2
}

fmt.Println(doubleNum(x)) // Prints: 10
fmt.Println(doubleNum(y)) // Prints: 6

```

- [x] 2. Using Functions

- [x] 3. Scope

- [x] 4. Returning Values from Functions

- [x] 5. Using Function Parameters

```go

func multiplier(x, y int32) int32 {
  return x * y
}

func main() {
  var product int32
  product = multiplier(25, 4)
  fmt.Println(product) // Prints: 100
}

```

- [x] 6. Reusing Code with Functions

- [x] 7. Multiple Return Values

- [x] 8. Deferring Resolution

```go
func calculateTaxes(revenue, deductions, credits float64) float64 {
  defer fmt.Println("Taxes Calculated!")
  taxRate := .06143
  fmt.Println("Calculating Taxes")

  if deductions == 0 || credits == 0 {
    return revenue * taxRate
  }


  taxValue := (revenue - (deductions * credits)) * taxRate
  if taxValue >= 0 {
    return taxValue
  } else {
    return 0
  }
}
```

## LESSON 6 Learn Go: Addresses and Pointers

Dive deeper into how Go stores values in memory using addresses and pointers!

Exercises

- [x] 1. The Point of Pointers and Addresses

Go is a pass-by-value language. In other words, we’re passing functions the value of an argument. In a technical sense, when we’re calling a function with an argument, the Go compiler is strictly using the <b>value</b> of the argument rather than the argument itself. Because of this feature (pass-by-value), the changes that take place in our function, stay within that function

But, we do have the ability to change values from different scopes. To do so, we need to make use of:

- addresses
- pointers
- dereferencing

- [ ] 2. Addresses

```go
x := "My very first address"
fmt.Println(&x) // Prints 0x414020

```

- [x] 3. Pointers

![addresses](./src/addresses.png)

- [x] 4. Dereferencing

- [x] 5. Changing Values in Different Scopes

---

[Documentation The Go programming language](https://golang.org/doc/)

[Guide for developing Go locally](https://www.codecademy.com/articles/setting-up-go-locally)
