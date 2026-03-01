// Package functions
package functions

import "fmt"

// Run demonstrates all function patterns in Go.
func Run() {
	fmt.Println("--- Basic Functions ---")
	myMessage()
	familyName("Liam")
	familyName("Jenny")

	fmt.Println("\n--- Multiple Parameters ---")
	familyNameAndAge("Tolu", 25)

	fmt.Println("\n--- Return Values ---")
	fmt.Println("sum:", myFunction(1, 2))
	fmt.Println("named return:", myFunction1(3, 4))

	fmt.Println("\n--- Multiple Return Values ---")
	result, text := myFunction2(5, "Hello")
	fmt.Println(result, text)

	// _ ignores a return value you don't need
	// Go REQUIRES you to use every declared variable, so _ is the escape hatch
	_, text2 := myFunction2(5, "World")
	fmt.Println("ignored first return, text2:", text2)

	fmt.Println("\n--- Recursion ---")
	testcount(1)

	fmt.Println("\n--- Factorial (recursion) ---")
	fmt.Println("10! =", factorial(10))
}

// myMessage — no parameters, no return value
func myMessage() {
	fmt.Println("I just got executed!")
}

// familyName — one parameter
// fname is the parameter name, string is its type
func familyName(fname string) {
	fmt.Println("Hello, I am", fname, "Neeson")
}

// familyNameAndAge — multiple parameters
// each parameter needs its own name and type
func familyNameAndAge(fname string, age int) {
	fmt.Println("Hello, I am", age, "years old and my name is", fname)
}

// myFunction — returns a single value
// The (int) after the params declares the return type
func myFunction(x int, y int) int {
	return x + y
}

// myFunction1 — named return value
//
// result is declared as part of the signature itself.
// The bare return at the end returns whatever result holds at that point.
// Named returns are useful for documenting what the return value represents.
func myFunction1(x int, y int) (result int) {
	result = x + y
	return // same as: return result
}

// myFunction2 — multiple return values
//
// Go lets you return more than one value from a function.
// The caller can either use both, one (with _), or store them in separate variables.
func myFunction2(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return // returns both result and txt1
}

// testcount — recursion example
//
// A function is recursive when it calls itself.
// It MUST have a stop condition (base case), otherwise it calls itself forever → stack overflow.
// Here the base case is: when x equals 11, stop.
func testcount(x int) int {
	if x == 11 {
		return 0 // base case — stop here
	}
	fmt.Println(x)
	return testcount(x + 1) // calls itself, moving x closer to 11 each time
}

// factorial — classic recursion
//
// factorial(5) = 5 × factorial(4)
//
//	= 5 × 4 × factorial(3)
//	= 5 × 4 × 3 × factorial(2)
//	= 5 × 4 × 3 × 2 × factorial(1)
//	= 5 × 4 × 3 × 2 × 1 × factorial(0)
//	= 5 × 4 × 3 × 2 × 1 × 1  ← base case
//	= 120
func factorial(n int) int {
	if n == 0 {
		return 1 // base case — factorial(0) is defined as 1
	}
	return n * factorial(n-1)
}
