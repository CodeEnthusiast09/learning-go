// Package advancedfunctions
package advancedfunctions

import "fmt"

// ————————————————————————————————————————————————————————————————————
// Currying is transforming a function that takes multiple arguments
// into a sequence of functions that each take one argument, so you
// can lock in arguments one at a time and reuse the result.
// ————————————————————————————————————————————————————————————————————

// ─────────────────────────────────────────────
// STEP 1: Normal function — all args at once
// ─────────────────────────────────────────────
func normalAdd(x, y int) int {
	return x + y
}

// ─────────────────────────────────────────────
// STEP 2: Curried version — one arg at a time
//
// Read the return type carefully:
//   - curriedAdd takes ONE int (x)
//   - it returns a func that takes ONE int (y)
//   - that inner func returns the final int
//
// So the chain is: int → func(int) → int
// ─────────────────────────────────────────────
func curriedAdd(x int) func(int) int {
	return func(y int) int {
		return x + y // x is locked in from the outer call
	}
}

// ─────────────────────────────────────────────
// STEP 3: Real world example
//
// Imagine you run an e-commerce site.
// You have different tax rates for different countries.
// Instead of passing the tax rate every single time,
// you lock it in once and reuse it.
// ─────────────────────────────────────────────
func applyTax(taxRate float64) func(float64) float64 {
	return func(price float64) float64 {
		return price + (price * taxRate / 100)
	}
}

// ─────────────────────────────────────────────
// STEP 4: Three levels deep
//
// You can curry as many arguments as you want.
// Each call locks in one more argument.
//
// Chain: string → func(int) → func(int) → string
// ─────────────────────────────────────────────
func buildMessage(greeting string) func(int) func(string) string {
	return func(count int) func(string) string {
		return func(name string) string {
			return fmt.Sprintf("%s %s! You have %d messages.", greeting, name, count)
		}
	}
}

func RunCurrying() {
	// ── Normal vs Curried ──────────────────────
	fmt.Println("─── Normal vs Curried ───")
	fmt.Println("normalAdd(2, 3):", normalAdd(2, 3))   // 5
	fmt.Println("curriedAdd(2)(3):", curriedAdd(2)(3)) // 5 — same result

	// ── Partial application (the real power) ──
	fmt.Println("\n─── Partial Application ───")
	// Stop halfway — lock in 10, get back a reusable function
	addTen := curriedAdd(10)
	// addTen is now a ready-made tool that always adds 10
	fmt.Println("addTen(1):", addTen(1))   // 11
	fmt.Println("addTen(5):", addTen(5))   // 15
	fmt.Println("addTen(90):", addTen(90)) // 100

	// ── Real world: tax calculator ─────────────
	fmt.Println("\n─── Tax Calculator ───")
	// Lock in the tax rate per country once
	nigerianTax := applyTax(7.5)
	ukTax := applyTax(20.0)
	usTax := applyTax(8.25)

	// Now reuse each one freely without repeating the rate
	fmt.Printf("NGN price 5000 after tax: %.2f\n", nigerianTax(5000)) // 5375.00
	fmt.Printf("GBP price 100 after tax:  %.2f\n", ukTax(100))        // 120.00
	fmt.Printf("USD price 80 after tax:   %.2f\n", usTax(80))         // 86.60

	// ── Three levels deep ─────────────────────
	fmt.Println("\n─── Three Levels Deep ───")

	// All at once — chained
	fmt.Println(buildMessage("Hello")(3)("Alice"))

	// Step by step — lock in one at a time
	sayGoodMorning := buildMessage("Good morning") // locked: greeting
	morningWith5 := sayGoodMorning(5)              // locked: count = 5

	// Now reuse morningWith5 for different names
	fmt.Println(morningWith5("Alice")) // Good morning Alice! You have 5 messages.
	fmt.Println(morningWith5("Bob"))   // Good morning Bob! You have 5 messages.
	fmt.Println(morningWith5("Carol")) // Good morning Carol! You have 5 messages.
}
