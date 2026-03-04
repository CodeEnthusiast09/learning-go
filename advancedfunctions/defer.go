// Package advancedfunctions
package advancedfunctions

import "fmt"

// ──────────────────────────────────────────────────────────────────────────────
// defer schedules a function call to run at the end of the surrounding function,
// no matter how that function exits — making it perfect for cleanup work you
// never want to forget.
// ──────────────────────────────────────────────────────────────────────────────

// ─────────────────────────────────────────────
// BASIC DEFER
// The deferred line runs last, even though it's
// written first in the function body
// ─────────────────────────────────────────────
func basicDefer() {
	defer fmt.Println("3. I run last (deferred)")
	fmt.Println("1. I run first")
	fmt.Println("2. I run second")
}

// ─────────────────────────────────────────────
// MULTIPLE DEFERS — LIFO ORDER
// Last deferred = first to run
// Think of a stack of plates
// ─────────────────────────────────────────────
func multipleDefers() {
	defer fmt.Println("I was deferred first — I run LAST")
	defer fmt.Println("I was deferred second — I run MIDDLE")
	defer fmt.Println("I was deferred third — I run FIRST")

	fmt.Println("Normal code runs before any deferred calls")
}

// ────────────────────────────────────────────────
// RULE 2 — Arguments are captured immediately
//
// Even though the Println is deferred,
// the VALUE of `x` is read right now — not later.
// So even if x changes after the defer line,
// the deferred call still uses the original value.
// ────────────────────────────────────────────────
func argumentCapture() {
	x := 10
	defer fmt.Println("deferred x:", x) // captures x=10 RIGHT NOW

	x = 99                       // x changes AFTER defer — but defer already captured 10
	fmt.Println("current x:", x) // prints 99
	// deferred line will still print 10
}

// ─────────────────────────────────────────────
// REAL WORLD USE CASE — cleanup
//
// The most common use of defer is making sure
// you clean up after yourself.
//
// Imagine opening a file or a database connection.
// You want to GUARANTEE it closes, no matter what
// happens in the middle of the function.
// ─────────────────────────────────────────────
func processFile(filename string) {
	fmt.Printf("Opening file: %s\n", filename)

	// Defer the close RIGHT AFTER you open.
	// You write it here so you never forget,
	// and Go guarantees it runs when the function exits.
	defer fmt.Printf("Closing file: %s\n", filename)

	// ... imagine doing work with the file here ...
	fmt.Println("Reading from file...")
	fmt.Println("Processing data...")
	// No matter what happens above, the file WILL be closed
}

// ─────────────────────────────────────────────
// DEFER WITH EARLY RETURN
//
// Even if the function returns early,
// deferred calls still run.
// ─────────────────────────────────────────────
func earlyReturn(fail bool) {
	defer fmt.Println("deferred: I ALWAYS run")

	if fail {
		fmt.Println("returning early!")
		return // defer still runs even here
	}

	fmt.Println("completing normally")
}

func RunDefer() {
	fmt.Println("─── Basic Defer ───")
	basicDefer()

	fmt.Println("\n─── Multiple Defers (LIFO) ───")
	multipleDefers()

	fmt.Println("\n─── Argument Capture ───")
	argumentCapture()

	fmt.Println("\n─── Real World: File Cleanup ───")
	processFile("data.txt")

	fmt.Println("\n─── Early Return ───")
	fmt.Println("with fail=true:")
	earlyReturn(true)
	fmt.Println("\nwith fail=false:")
	earlyReturn(false)
}
