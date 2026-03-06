// Package advancedfunctions
package advancedfunctions

import (
	"fmt"
	"sort"
)

// ─────────────────────────────────────────────
// 1. ASSIGNED TO A VARIABLE
//
// The function has no name, but the VARIABLE has one.
// You call it through the variable.
// ─────────────────────────────────────────────
func assignedToVariable() {
	// `square` is the variable — the function itself is anonymous
	square := func(n int) int {
		return n * n
	}

	fmt.Println("square(4):", square(4)) // 16
	fmt.Println("square(9):", square(9)) // 81

	// You can even reassign it to a completely different function
	// because `square` is just a variable holding a func(int) int
	square = func(n int) int {
		return n * n * n // cube now
	}

	fmt.Println("square reassigned to cube(4):", square(4)) // 64
}

// ─────────────────────────────────────────────
// 2. PASSED DIRECTLY AS AN ARGUMENT
//
// Instead of defining a function separately and passing it,
// you write the function inline right where it is needed.
// The sort.Slice function from the standard library
// is a perfect real world example of this.
// ─────────────────────────────────────────────
func passedAsArgument() {
	numbers := []int{5, 2, 8, 1, 9, 3}

	fmt.Println("before sort:", numbers)

	// The second argument to sort.Slice IS an anonymous function
	// defined right here inline — no need to name it
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	// slices.Sort(numbers)

	fmt.Println("after sort using the sort package:", numbers)

	// fmt.Println("after sort using the slice package:", numbers)

	// Another example — apply an anonymous operation to each item
	words := []string{"go", "is", "awesome"}
	transform(words, func(s string) string {
		return "[" + s + "]"
	})
}

// helper — applies a transformation function to each string in a slice
func transform(items []string, operation func(string) string) {
	for _, item := range items {
		fmt.Println(operation(item))
	}
}

// ─────────────────────────────────────────────
// 3. IMMEDIATELY INVOKED (IIFE)
// Immediately Invoked Function Expression
//
// You define the function AND call it on the same line.
// The () at the end is what calls it immediately.
// Useful for one-off logic you don't need to reuse.
// ─────────────────────────────────────────────
func immediatelyInvoked() {
	// defined and called right here — never stored, never reused
	result := func(a, b int) int {
		return a * b
	}(6, 7) // ← the (6, 7) calls it immediately with these arguments

	fmt.Println("IIFE result:", result) // 42

	// Useful for initialising something with complex logic
	config := func() map[string]string {
		m := make(map[string]string)
		m["env"] = "production"
		m["version"] = "1.0.0"
		m["region"] = "us-east"
		return m
	}() // called immediately — config holds the map, not the function

	fmt.Println("config:", config)
}

// ─────────────────────────────────────────────
// 4. ANONYMOUS FUNCTIONS + CLOSURES TOGETHER
//
// Anonymous functions become closures the moment
// they reference a variable from their surrounding scope.
// Most closures you write in practice WILL be anonymous.
// ─────────────────────────────────────────────
func withClosure() {
	prefix := "LOG"

	// This anonymous function captures `prefix` — making it a closure
	log := func(message string) {
		fmt.Printf("[%s] %s\n", prefix, message)
	}

	log("server started")   // [LOG] server started
	log("request received") // [LOG] request received

	prefix = "ERROR"            // change the captured variable
	log("something went wrong") // [ERROR] something went wrong — reflects the change
}

func RunAnonymousFunctions() {
	fmt.Println("─── Assigned to Variable ───")
	assignedToVariable()

	fmt.Println("\n─── Passed as Argument ───")
	passedAsArgument()

	fmt.Println("\n─── Immediately Invoked (IIFE) ───")
	immediatelyInvoked()

	fmt.Println("\n─── Anonymous + Closure ───")
	withClosure()
}
