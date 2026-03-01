// Package generics
package generics

import "fmt"

// ==============================================================
// GENERICS IN GO (Go 1.18+, updated through Go 1.26)
// ==============================================================
//
// WHAT PROBLEM DO GENERICS SOLVE?
//
// Without generics, if you want a function that works on both
// int and float64, you'd have to write it TWICE:
//
//   func SumInts(nums []int) int { ... }
//   func SumFloats(nums []float64) float64 { ... }
//
// That's copy-paste code. Fix a bug in one, must fix it in both.
//
// Generics let you write it ONCE and say:
//   "this function works for any type that fits this constraint"
// ==============================================================

// ==============================================================
// PART 1: YOUR FIRST GENERIC FUNCTION
// ==============================================================
//
// Syntax: func FunctionName[T constraint](param T) T
//                             ^^^^^^^^^^^
//                             type parameter — T is a placeholder for a real type
//
// `any` means "literally any type" — it's an alias for interface{}

func printValue[T any](value T) {
	fmt.Printf("  value: %v (type: %T)\n", value, value)
}

// first returns the first element of any slice.
func first[T any](slice []T) (T, bool) {
	if len(slice) == 0 {
		var zero T // zero value: 0 for int, "" for string, nil for pointer...
		return zero, false
	}
	return slice[0], true
}

// ==============================================================
// PART 2: CONSTRAINTS
// ==============================================================
//
// `any` is too broad — you can't do math on `any` type.
// Constraints define WHAT TYPES ARE ALLOWED.
// A constraint is just an interface listing permitted types.

// Number allows int, float32, or float64 — nothing else.
type Number interface {
	int | float32 | float64
}

func sum[T Number](nums []T) T {
	var total T
	for _, n := range nums {
		total += n // safe — we know T supports + because it's int or float
	}
	return total
}

// ==============================================================
// PART 3: THE ~ OPERATOR (underlying type)
// ==============================================================
//
// `~int` means: "int, AND any type whose UNDERLYING type is int"
//
// If someone defines: type UserID int
// Then UserID's underlying type is int.
// `int` alone would NOT accept UserID.
// `~int` WOULD accept UserID.

type (
	UserID    int
	ProductID int
)

type idConstraint interface {
	~int
}

func printID[T idConstraint](id T) {
	fmt.Printf("  ID: %v\n", int(id))
}

// ==============================================================
// PART 4: COMPARABLE
// ==============================================================
//
// `comparable` is a built-in constraint meaning the type supports == and !=
// Required for using the type as a map key or for equality comparisons.

func contains[T comparable](slice []T, target T) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

func getOrDefault[K comparable, V any](m map[K]V, key K, defaultVal V) V {
	if val, ok := m[key]; ok {
		return val
	}
	return defaultVal
}

// ==============================================================
// PART 5: MULTIPLE TYPE PARAMETERS
// ==============================================================

func mapSlice[T any, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

func filter[T any](slice []T, fn func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

func reduce[T any, U any](slice []T, initial U, fn func(U, T) U) U {
	acc := initial
	for _, v := range slice {
		acc = fn(acc, v)
	}
	return acc
}

// ==============================================================
// PART 6: GENERIC TYPES (structs)
// ==============================================================

// Stack is a generic LIFO (last in, first out) data structure.
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

func (s *Stack[T]) Len() int { return len(s.items) }

// Pair holds two values of potentially different types.
type Pair[A any, B any] struct {
	First  A
	Second B
}

func newPair[A any, B any](first A, second B) Pair[A, B] {
	return Pair[A, B]{First: first, Second: second}
}

// ==============================================================
// PART 7: GENERIC TYPE ALIASES (Go 1.24+)
// ==============================================================
//
// type StringStack = Stack[string]  → alias (same type, just a different name)
// type StringStack Stack[string]    → new type (different from Stack[string])

type (
	StringStack = Stack[string]  // alias — Go 1.24+
	IntPair     = Pair[int, int] // alias — Go 1.24+
)

// ==============================================================
// ORDERED CONSTRAINT
// (replaces golang.org/x/exp/constraints which isn't in stdlib)
// ==============================================================

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~string
}

func minOf[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func maxOf[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// ==============================================================
// RUN
// ==============================================================

func Run() {
	fmt.Println("--- Basic Generic Functions ---")
	printValue(42)
	printValue("hello")
	printValue(3.14)

	nums := []int{10, 20, 30}
	if f, ok := first(nums); ok {
		fmt.Println("  first:", f)
	}

	fmt.Println("\n--- Sum with Number constraint ---")
	fmt.Println("  sum ints:", sum([]int{1, 2, 3, 4, 5}))
	fmt.Println("  sum floats:", sum([]float64{1.1, 2.2, 3.3}))

	fmt.Println("\n--- ~ underlying type ---")
	var uid UserID = 42
	var pid ProductID = 99
	printID(uid)
	printID(pid)
	printID(100) // plain int also works

	fmt.Println("\n--- comparable constraint ---")
	fmt.Println("  contains 3:", contains([]int{1, 2, 3, 4}, 3))
	fmt.Println("  contains 9:", contains([]int{1, 2, 3, 4}, 9))

	ages := map[string]int{"alice": 28, "bob": 34}
	fmt.Println("  alice age:", getOrDefault(ages, "alice", 0))
	fmt.Println("  carol age (default):", getOrDefault(ages, "carol", -1))

	fmt.Println("\n--- map / filter / reduce ---")
	doubled := mapSlice([]int{1, 2, 3, 4, 5}, func(n int) int { return n * 2 })
	fmt.Println("  doubled:", doubled)

	evens := filter([]int{1, 2, 3, 4, 5, 6}, func(n int) bool { return n%2 == 0 })
	fmt.Println("  evens:", evens)

	total := reduce([]int{1, 2, 3, 4, 5}, 0, func(acc, n int) int { return acc + n })
	fmt.Println("  sum via reduce:", total)

	names := mapSlice([]int{1, 2, 3}, func(n int) string {
		return fmt.Sprintf("item-%d", n)
	})
	fmt.Println("  int→string:", names)

	fmt.Println("\n--- Generic Stack ---")
	var s Stack[int]
	s.Push(10)
	s.Push(20)
	s.Push(30)
	if val, ok := s.Pop(); ok {
		fmt.Println("  popped:", val) // 30
	}
	fmt.Println("  stack size:", s.Len())

	var ss StringStack
	ss.Push("hello")
	ss.Push("world")
	if val, ok := ss.Pop(); ok {
		fmt.Println("  string stack popped:", val)
	}

	fmt.Println("\n--- Pair ---")
	p := newPair("Alice", 28)
	fmt.Printf("  %s is %d years old\n", p.First, p.Second)

	ip := IntPair{First: 1, Second: 2}
	fmt.Println("  IntPair:", ip)

	fmt.Println("\n--- Min / Max with Ordered constraint ---")
	fmt.Println("  min(3,7):", minOf(3, 7))
	fmt.Println("  max(3.14,2.71):", maxOf(3.14, 2.71))
	fmt.Println("  min strings:", minOf("apple", "banana"))
}
