// Package pointers
package pointers

import "fmt"

// ==============================================================
// POINTERS IN GO
// ==============================================================
//
// A pointer is a variable that stores the MEMORY ADDRESS of another variable.
//
// Think of memory like a street with houses (addresses).
//   A normal variable is a house that holds a value.
//   A pointer is a piece of paper with that house's ADDRESS written on it.
//
// Two key operators:
//   &  →  "give me the ADDRESS of this variable"     (address-of)
//   *  →  "go to the address and get what's there"   (dereference)
// ==============================================================

// person is a local type used for pointer-with-struct demos.
type person struct {
	name string
	age  int
}

// rectangle is a local type used for value vs pointer receiver demos.
type rectangle struct {
	width  float64
	height float64
}

// area uses a VALUE receiver — only reads, no modification needed.
func (r rectangle) area() float64 {
	return r.width * r.height
}

// scale uses a POINTER receiver — modifies the original rectangle.
func (r *rectangle) scale(factor float64) {
	r.width *= factor
	r.height *= factor
}

// Run demonstrates all pointer concepts.
func Run() {
	basicPointer()
	whyPointers()
	pointerToStruct()
	nilPointer()
	newFunction()
	valueVsPointerReceiver()
}

// ==============================================================
// 1. BASIC POINTER
// ==============================================================

func basicPointer() {
	fmt.Println("--- Basic Pointer ---")

	age := 25
	ptr := &age // ptr holds the MEMORY ADDRESS of age

	fmt.Println("value of age:", age)      // 25
	fmt.Println("address of age:", ptr)    // e.g. 0xc000018090
	fmt.Println("value at address:", *ptr) // 25 ← dereference: "go get what's at that address"

	// Changing the value THROUGH the pointer changes the original variable
	*ptr = 30
	fmt.Println("age after *ptr = 30:", age) // 30 — the original changed!
}

// ==============================================================
// 2. WHY DO WE NEED POINTERS?
// ==============================================================
//
// When you pass a variable to a function, Go makes a COPY.
// The function works on the copy — the original is untouched.
// Pointers let you say: "don't copy, work on the REAL thing."

func whyPointers() {
	fmt.Println("\n--- Why Pointers? (copy vs reference) ---")

	score := 10

	tryToDouble(score)
	fmt.Println("after tryToDouble (by value):", score) // still 10 — copy was doubled

	actuallyDouble(&score)                                   // pass the ADDRESS
	fmt.Println("after actuallyDouble (by pointer):", score) // 20 — original changed!
}

// tryToDouble works on a COPY of n — original untouched.
func tryToDouble(n int) {
	n = n * 2 // only changes the local copy
}

// actuallyDouble dereferences the pointer to reach and modify the original.
func actuallyDouble(n *int) {
	*n = *n * 2 // *n = "go to the address n points at and update the value there"
}

// ==============================================================
// 3. POINTER TO A STRUCT
// ==============================================================
//
// The most common use of pointers in Go.
// Instead of copying an entire struct every time, pass a pointer.
// Go lets you access struct fields through a pointer WITHOUT writing (*p).Name —
// it auto-dereferences: p.Name is the same as (*p).Name.

func pointerToStruct() {
	fmt.Println("\n--- Pointer to Struct ---")

	user := person{name: "Alice", age: 28}

	fmt.Println("before birthday:", user.age) // 28
	haveBirthday(&user)
	fmt.Println("after birthday:", user.age) // 29

	p := &user
	fmt.Println("explicit dereference (*p).age:", (*p).age) // 29
	fmt.Println("auto-dereference p.age:", p.age)           // 29 — same, Go handles it
}

func haveBirthday(u *person) {
	u.age++ // Go auto-dereferences: same as (*u).age++
}

// ==============================================================
// 4. NIL POINTER
// ==============================================================
//
// A pointer declared but not assigned points to NOTHING — it's nil.
// Dereferencing nil crashes the program (nil pointer dereference panic).
// Always check for nil before using a pointer.

func nilPointer() {
	fmt.Println("\n--- Nil Pointer ---")

	var nilPtr *int                      // declared but not pointing at anything
	fmt.Println("nilPtr value:", nilPtr) // <nil>

	if nilPtr == nil {
		fmt.Println("pointer is nil — safe, we didn't dereference it")
	}
	// *nilPtr would PANIC here — never dereference nil
}

// ==============================================================
// 5. new() FUNCTION
// ==============================================================
//
// new(T) allocates memory for a T, sets it to its zero value,
// and returns a *T (pointer to T).
// Useful when you want a pointer but don't have an existing variable.

func newFunction() {
	fmt.Println("\n--- new() function ---")

	n := new(int)                 // allocates an int, zero value = 0, returns *int
	fmt.Println("*n before:", *n) // 0
	*n = 99
	fmt.Println("*n after:", *n) // 99
}

// ==============================================================
// 6. VALUE RECEIVER vs POINTER RECEIVER
// ==============================================================
//
// Value receiver (r rectangle) → works on a COPY — cannot modify the original
// Pointer receiver (*rectangle)  → works on the REAL struct — CAN modify it
//
// Rule of thumb:
//   Use pointer receiver if the method needs to MODIFY the struct,
//   or if the struct is large (avoids expensive copying).
//   Use value receiver for small, read-only operations.

func valueVsPointerReceiver() {
	fmt.Println("\n--- Value vs Pointer Receiver ---")

	rect := rectangle{width: 10, height: 5}
	fmt.Println("area (value receiver):", rect.area()) // reads, no copy issue

	fmt.Println("before scale:", rect)
	rect.scale(3)                        // pointer receiver — modifies the REAL rect
	fmt.Println("after scale(3):", rect) // width: 30, height: 15
}
