// Package sprint
package sprint

import "fmt"

// Run demonstrates fmt.Sprint, fmt.Sprintln, and fmt.Sprintf.
func Run() {
	sprintDemo()
	sprintlnDemo()
	sprintfDemo()
}

// ==============================================================
// fmt.Sprint()
// ==============================================================
//
// Sprint() glues values into a string WITHOUT printing them.
// No newline is added at the end.
//
// Space rules (easy to get wrong):
//   Both args are non-strings → Go adds a space between them automatically
//   At least one arg is a string → NO auto-space, you must add it yourself

func sprintDemo() {
	fmt.Println("--- fmt.Sprint() ---")

	s1 := fmt.Sprint("Hello", " ", "World")
	// Both args are strings → no auto-space → we manually added " " as a third arg
	// s1 = "Hello World"

	s2 := fmt.Sprint(10, 20)
	// Both are ints (non-strings) → auto-space added between them
	// s2 = "10 20"

	s3 := fmt.Sprint("Score:", 100)
	// One is a string → NO auto-space
	// s3 = "Score:100"

	fmt.Println(s1) // Hello World
	fmt.Println(s2) // 10 20
	fmt.Println(s3) // Score:100
}

// ==============================================================
// fmt.Sprintln()
// ==============================================================
//
// Sprintln() is like Sprint but:
//   1. ALWAYS adds a space between EVERY argument (regardless of type)
//   2. ALWAYS adds a newline \n at the very end of the string

func sprintlnDemo() {
	fmt.Println("\n--- fmt.Sprintln() ---")

	sl1 := fmt.Sprintln("Hello", "World", 42)
	// sl1 = "Hello World 42\n"
	// Sprintln adds spaces between all args, plus a newline at the end

	sl2 := fmt.Sprintln("Score:", 100)
	// sl2 = "Score: 100\n"
	// Even string + int → Sprintln still adds the space (unlike Sprint)

	// Use fmt.Print (not Println) here because the \n is already baked into the string.
	// Using Println would give a double newline — an empty blank line.
	fmt.Print(sl1) // Hello World 42
	fmt.Print(sl2) // Score: 100
}

// ==============================================================
// fmt.Sprintf()
// ==============================================================
//
// Sprintf() is the most powerful — a format template with verbs.
// It builds and RETURNS the string without printing it.
// This lets you reuse the string: pass it to a function, store it, build on it.
//
// Common verbs:
//   %s    → string
//   %d    → integer
//   %f    → float  (%.2f = 2 decimal places)
//   %t    → boolean (true/false)
//   %T    → the TYPE of the variable, not the value
//   %v    → any value in its default format

func sprintfDemo() {
	fmt.Println("\n--- fmt.Sprintf() ---")

	name := "Tolu"
	age := 25
	balance := 1500.75
	isActive := true

	greeting := fmt.Sprintf("My name is %s and I am %d years old.", name, age)
	moneyMsg := fmt.Sprintf("Account balance: $%.2f", balance)
	statusMsg := fmt.Sprintf("Account active: %t", isActive)
	typeMsg := fmt.Sprintf("Type of balance: %T", balance)

	fmt.Println(greeting)  // My name is Tolu and I am 25 years old.
	fmt.Println(moneyMsg)  // Account balance: $1500.75
	fmt.Println(statusMsg) // Account active: true
	fmt.Println(typeMsg)   // Type of balance: float64

	// The key advantage: we can now PASS the string somewhere
	fmt.Println("\n--- Passing Sprintf result to a function ---")
	sendMessage(greeting)
}

// sendMessage represents any function that consumes a string —
// send an email, write to a file, call an API, etc.
// The point: Sprintf built the string, and now we can do anything with it.
func sendMessage(msg string) {
	fmt.Println("  Message sent:", msg)
}
