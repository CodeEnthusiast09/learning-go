// Package pkgmodules
package pkgmodules

import (
	"cmp"
	"fmt"
	"slices"
)

type person struct {
	name string
	age  int
}

func RunSlicesSort() {
	// ── slices.Sort — simple, just works ──────
	fmt.Println("─── slices.Sort ───")
	numbers := []int{5, 2, 8, 1, 9, 3}
	slices.Sort(numbers)
	fmt.Println("sorted ints:", numbers) // [1 2 3 5 8 9]

	words := []string{"banana", "apple", "cherry"}
	slices.Sort(words)
	fmt.Println("sorted strings:", words) // [apple banana cherry]

	// ── slices.SortFunc — custom logic ────────
	// Use this when slices.Sort isn't enough:
	// sorting structs, reverse order, custom rules etc.
	fmt.Println("\n─── slices.SortFunc ───")
	people := []person{
		{"Alice", 30},
		{"Bob", 25},
		{"Carol", 35},
	}

	// Sort by age ascending
	slices.SortFunc(people, func(a, b person) int {
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println("by age asc:", people)

	// Sort by name ascending
	slices.SortFunc(people, func(a, b person) int {
		return cmp.Compare(a.name, b.name)
	})
	fmt.Println("by name asc:", people)

	// Sort by age DESCENDING — just flip a and b
	slices.SortFunc(people, func(a, b person) int {
		return cmp.Compare(b.age, a.age) // b before a = descending
	})
	fmt.Println("by age desc:", people)

	// ── slices.IsSorted — check before sorting ─
	fmt.Println("\n─── slices.IsSorted ───")
	sorted := []int{1, 2, 3, 4, 5}
	unsorted := []int{3, 1, 4, 1, 5}
	fmt.Println("sorted?", slices.IsSorted(sorted))   // true
	fmt.Println("sorted?", slices.IsSorted(unsorted)) // false

	// ── slices.SortStableFunc — preserves order of equal elements ─
	// Regular sort may swap equal elements around.
	// Stable sort guarantees equal elements stay in their original order.
	fmt.Println("\n─── slices.SortStableFunc ───")
	employees := []person{
		{"Alice", 30},
		{"Bob", 30},   // same age as Alice
		{"Carol", 30}, // same age as Alice and Bob
	}

	slices.SortStableFunc(employees, func(a, b person) int {
		return cmp.Compare(a.age, b.age)
	})
	// Alice, Bob, Carol — original order preserved among equals
	fmt.Println("stable sort:", employees)
}
