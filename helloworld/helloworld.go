// Package helloworld contains topics from variable to loops
package helloworld

import "fmt"

// Run demonstrates Go basics: variables, constants, fmt verbs,
// arrays, slices, operators, conditions, switch statements, and loops.
func Run() {
	variablesAndConstants()
	fmtVerbs()
	arraysAndSlices()
	operators()
	conditionsAndSwitch()
	loops()
}

// ==============================================================
// VARIABLES & CONSTANTS
// ==============================================================
//
// Go has several ways to declare variables:
//
//   var name string = "John"   → explicit type + value
//   var name = "John"          → type inferred from value
//   name := "John"             → shorthand (only inside functions)
//
// Constants use const and can never be changed after declaration.
// Constants can be typed (const G int = 10) or untyped (const PI = 3.14).
// Untyped constants are more flexible — Go infers the type at use time.

func variablesAndConstants() {
	fmt.Println("--- Variables & Constants ---")

	var student1 string = "John" // explicit type
	student2 := "Jane"           // inferred type
	x := 2                       // shorthand declaration (most common in practice)

	fmt.Println(student1, student2, x)

	// Zero values — variables declared without a value get a safe default
	// string → "", int → 0, bool → false
	var a string
	var b int
	var c bool
	fmt.Printf("zero values: %q %d %t\n", a, b, c)

	// Block declaration — declare multiple variables at once, cleaner than one per line
	var (
		country  = "Nigeria"
		language = "English"
		year     = 2024
	)
	fmt.Println(country, language, year)

	// Constants — typed and untyped
	const PI = 3.142 // untyped — no explicit type, Go decides at use time
	const G int = 10 // typed — locked to int
	fmt.Println("PI:", PI, "G:", G)
}

// ==============================================================
// FMT VERBS
// ==============================================================
//
// Printf uses "verbs" as placeholders in a format string:
//   %v  → value in default format
//   %T  → type of the value
//   %s  → string
//   %d  → integer (base 10)
//   %f  → float  (%.2f = 2 decimal places)
//   %t  → boolean
//   %b  → integer in binary

func fmtVerbs() {
	fmt.Println("\n--- fmt Verbs ---")

	name := "Tolu"
	age := 25
	balance := 1500.75
	active := true

	fmt.Printf("%s is %d years old\n", name, age)
	fmt.Printf("balance: %.2f\n", balance) // 2 decimal places
	fmt.Printf("account active: %t\n", active)
	fmt.Printf("type of balance: %T\n", balance) // prints "float64"
	fmt.Printf("age in binary: %b\n", age)       // 11001
	fmt.Printf("hex: %x  octal: %o\n", age, age)
}

// ==============================================================
// ARRAYS & SLICES
// ==============================================================
//
// Array  → fixed length, [3]int — the size is part of the type
// Slice  → dynamic length, []int — can grow and shrink
//
// In practice, you almost always use slices.
// Arrays are mostly used when you know EXACTLY how many items you'll ever have.
//
// Key slice operations:
//   append(slice, val)      → adds an element, returns new slice
//   copy(dest, src)         → copies elements into a separate slice
//   slice[start:end]        → sub-slice (start inclusive, end exclusive)
//   len(slice)              → number of elements currently in the slice
//   cap(slice)              → how many elements it can hold before reallocation

func arraysAndSlices() {
	fmt.Println("\n--- Arrays & Slices ---")

	// Arrays — fixed size, declared with the size in the type
	arr := [3]int{10, 20, 30}
	fmt.Println("array:", arr, "len:", len(arr))

	// [...]int lets Go count the elements for you — still a fixed array
	arr2 := [...]string{"Go", "is", "fun"}
	fmt.Println("auto-count array:", arr2)

	// Slices — no size in the type, can grow
	fruits := []string{"apple", "banana", "mango"}
	fruits = append(fruits, "pawpaw")
	fmt.Println("slice after append:", fruits)

	// Sub-slice — [start:end], end is exclusive
	// fruits[1:3] means index 1 and 2 (not 3)
	fmt.Println("sub-slice [1:3]:", fruits[1:3])

	// make([]T, length, capacity)
	// length  = how many elements exist right now (all zero values)
	// capacity = how big the underlying array is before Go has to reallocate
	nums := make([]int, 3, 6)
	fmt.Printf("make slice: %v  len:%d  cap:%d\n", nums, len(nums), cap(nums))

	// copy — creates an independent slice, changes won't affect the original
	original := []int{1, 2, 3, 4, 5}
	small := original[:3] // shares memory with original!
	independent := make([]int, 3)
	copy(independent, original[:3]) // truly separate copy
	fmt.Println("original:", original, "small (shared):", small, "independent:", independent)
}

// ==============================================================
// OPERATORS
// ==============================================================

func operators() {
	fmt.Println("\n--- Operators ---")

	// Arithmetic
	fmt.Println("10 + 3 =", 10+3)
	fmt.Println("10 % 3 =", 10%3) // remainder

	// Assignment shorthand
	b := 20
	b += 5                                    // same as b = b + 5
	b *= 2                                    // same as b = b * 2
	fmt.Println("b after += 5 then *= 2:", b) // (20+5)*2 = 50

	// Comparison — always returns bool
	x, y := 10, 20
	fmt.Printf("x=%d y=%d  x>y:%t  x==y:%t  x!=y:%t\n", x, y, x > y, x == y, x != y)

	// Logical
	hasID := true
	isBanned := false
	fmt.Println("hasID && !isBanned:", hasID && !isBanned) // true
	fmt.Println("isBanned || hasID:", isBanned || hasID)   // true
}

// ==============================================================
// CONDITIONS & SWITCH
// ==============================================================
//
// Go's if/else works like most languages.
// One Go-specific feature: you can run a short statement INSIDE the if:
//   if x := getValue(); x > 0 { ... }
// x only exists within the if/else block — not outside.
//
// Switch in Go:
//   - No need for break — Go stops at the first matching case automatically
//   - fallthrough forces Go to also run the NEXT case
//   - Switch with no variable works like a chain of if/else if

func conditionsAndSwitch() {
	fmt.Println("\n--- Conditions & Switch ---")

	score := 75
	if score >= 90 {
		fmt.Println("grade: A")
	} else if score >= 80 {
		fmt.Println("grade: B")
	} else if score >= 70 {
		fmt.Println("grade: C")
	} else {
		fmt.Println("grade: F")
	}

	// Short statement in if — entryAge only lives inside this block
	if entryAge := 20; entryAge >= 18 {
		fmt.Println("entryAge", entryAge, "→ adult")
	}

	// Basic switch
	day := "Friday"
	switch day {
	case "Monday":
		fmt.Println("start of work week")
	case "Friday":
		fmt.Println("end of work week") // matches here, stops
	case "Saturday", "Sunday":
		fmt.Println("weekend!")
	default:
		fmt.Println("midweek")
	}

	// Switch with no variable — each case is its own full condition
	temperature := 35
	switch {
	case temperature >= 40:
		fmt.Println("extremely hot")
	case temperature >= 30:
		fmt.Println("hot") // 35 >= 30 → this runs
	default:
		fmt.Println("cold")
	}

	// fallthrough — forces the next case to ALSO run, ignoring its condition
	rating := 3
	switch rating {
	case 3:
		fmt.Println("average")
		fallthrough // next case runs too, even though rating != 2
	case 2:
		fmt.Println("below average (via fallthrough)")
	case 1:
		fmt.Println("poor")
	}
}

// ==============================================================
// LOOPS
// ==============================================================
//
// Go has ONLY one loop keyword: for
// But it can behave in three ways:
//
//   for i := 0; i < 5; i++ {}      → classic C-style loop
//   for i < 5 {}                    → acts like a while loop
//   for {}                          → infinite loop (needs a break)
//
// range loops over slices, arrays, maps, strings, and channels.
// It gives you two values each iteration: index and value.
// Use _ to ignore the one you don't need.

func loops() {
	fmt.Println("\n--- Loops ---")

	// Classic for loop
	for i := 0; i <= 3; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// range over integer (Go 1.22+) — like Python's range()
	for i := range 5 {
		fmt.Print(i, " ") // 0 1 2 3 4
	}
	fmt.Println()

	// for as a while loop
	j := 0
	for j < 3 {
		fmt.Print(j, " ")
		j++
	}
	fmt.Println()

	// range over slice — index and value
	fruits := []string{"apple", "banana", "mango"}
	for index, value := range fruits {
		fmt.Printf("  [%d] %s\n", index, value)
	}

	// ignore index with _
	for _, fruit := range fruits {
		fmt.Print(fruit, " ")
	}
	fmt.Println()

	// break and continue
	for k := range 10 {
		if k == 3 {
			continue // skip 3
		}
		if k == 7 {
			break // stop at 7
		}
		fmt.Print(k, " ") // 0 1 2 4 5 6
	}
	fmt.Println()

	// Nested loops — inner runs fully for every outer iteration
	for outer := 1; outer <= 3; outer++ {
		for inner := 1; inner <= 3; inner++ {
			fmt.Printf("%d×%d=%d  ", outer, inner, outer*inner)
		}
		fmt.Println()
	}
}
