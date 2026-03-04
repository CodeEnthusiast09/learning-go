// Package advancedfunctions
package advancedfunctions

import "fmt"

// ─────────────────────────────────────────────
// BASIC CLOSURE
//
// outer() creates a variable `message`
// inner() is defined inside outer() and uses `message`
// even after outer() finishes, inner() still knows `message`
// ─────────────────────────────────────────────
func outer() func() {
	message := "I was created inside outer()"

	// This inner function captures `message` from above
	// That makes it a closure
	inner := func() {
		fmt.Println(message) // still accessible even after outer() returns
	}

	return inner
}

// ─────────────────────────────────────────────
// MUTABLE STATE IN A CLOSURE
//
// The captured variable is NOT a copy.
// It is a reference to the SAME variable.
// So changes made inside the closure STICK.
// ─────────────────────────────────────────────
func makeStepCounter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

// ─────────────────────────────────────────────
// INDEPENDENT CLOSURES
//
// Each call to makeStepCounter() creates a BRAND NEW notebook.
// Two counters do not share state — they are completely isolated.
// ─────────────────────────────────────────────

// ─────────────────────────────────────────────
// CLOSURE OVER A LOOP VARIABLE — classic gotcha
//
// This is the most common mistake people make with closures.
// When you close over a loop variable, ALL closures
// end up pointing to the SAME variable — not copies of it.
// By the time they run, the loop is done and the variable
// holds its final value.
// ─────────────────────────────────────────────
func loopGotcha() {
	funcs := []func(){}
	i := 0 // OUTSIDE the loop — all closures share this ONE variable

	for ; i < 3; i++ {
		funcs = append(funcs, func() {
			fmt.Println("gotcha i:", i)
		})
	}

	for _, f := range funcs {
		f() // prints 3, 3, 3
	}
}

func loopFixed() {
	funcs := []func(){}

	for i := 0; i < 3; i++ {
		// FIX — create a NEW variable per iteration
		// each closure now captures its OWN copy
		i := i // shadows the outer i with a brand new variable
		funcs = append(funcs, func() {
			fmt.Println("fixed i:", i)
		})
	}

	for _, f := range funcs {
		f() // prints 0, 1, 2 — correct
	}
}

// ─────────────────────────────────────────────
// REAL WORLD USE CASE — a reusable adder factory
//
// Instead of writing addFive(), addTen(), addHundred()
// as separate functions, you build ONE factory function
// that generates them on demand.
// ─────────────────────────────────────────────
func makeAdder(amount int) func(int) int {
	// `amount` is captured in the closure's notebook
	return func(x int) int {
		return x + amount
	}
}

// ─────────────────────────────────────────────
// REAL WORLD USE CASE — access control
//
// The balance variable is completely private.
// Nobody outside can read or write it directly.
// The only way to interact with it is through
// the functions we hand back.
// This is basically encapsulation — like in OOP.
// ─────────────────────────────────────────────
type bankAccount struct {
	deposit  func(float64)
	withdraw func(float64)
	balance  func() float64
}

func newBankAccount(initial float64) bankAccount {
	balance := initial // private — nobody outside can touch this directly

	return bankAccount{
		deposit: func(amount float64) {
			balance += amount
			fmt.Printf("  deposited %.2f — new balance: %.2f\n", amount, balance)
		},
		withdraw: func(amount float64) {
			if amount > balance {
				fmt.Println("  insufficient funds")
				return
			}
			balance -= amount
			fmt.Printf("  withdrew %.2f — new balance: %.2f\n", amount, balance)
		},
		balance: func() float64 {
			return balance
		},
	}
}

func RunClosures() {
	fmt.Println("─── Basic Closure ───")
	// outer() finishes running here
	// but the function it returned still knows `message`
	rememberMe := outer()
	rememberMe() // still prints the message

	fmt.Println("\n─── Mutable State ───")
	counter := makeStepCounter()
	fmt.Println("counter():", counter()) // 1
	fmt.Println("counter():", counter()) // 2
	fmt.Println("counter():", counter()) // 3

	fmt.Println("\n─── Independent Closures ───")
	counterA := makeStepCounter() // brand new notebook
	counterB := makeStepCounter() // completely separate notebook
	counterA()
	counterA()
	counterB()
	fmt.Println("counterA:", counterA()) // 3 — called 3 times
	fmt.Println("counterB:", counterB()) // 2 — called 2 times, independent

	fmt.Println("\n─── Loop Gotcha ───")
	loopGotcha() // 3, 3, 3

	fmt.Println("\n─── Loop Fixed ───")
	loopFixed() // 0, 1, 2

	fmt.Println("\n─── Adder Factory ───")
	addFive := makeAdder(5)
	addHundred := makeAdder(100)
	fmt.Println("addFive(3):", addFive(3))       // 8
	fmt.Println("addFive(10):", addFive(10))     // 15
	fmt.Println("addHundred(3):", addHundred(3)) // 103

	fmt.Println("\n─── Bank Account (encapsulation) ───")
	account := newBankAccount(1000.00)
	fmt.Printf("opening balance: %.2f\n", account.balance())
	account.deposit(500)
	account.withdraw(200)
	account.withdraw(2000) // should fail
	fmt.Printf("final balance: %.2f\n", account.balance())
}
