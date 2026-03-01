// Package mapdemo contains the topice maps
package mapdemo

// NOTE: This package is named "mapdemo" (not "maps") because Go's standard
// library already has a package called "maps". If we named ours "maps" too,
// the import path would shadow the stdlib one and cause confusion.

import (
	"fmt"
	"sort"
)

// Run demonstrates all map operations in Go.
func Run() {
	nilMapDemo()
	makeMapDemo()
	mapLiteralDemo()
	keyExistenceDemo()
	addUpdateDeleteDemo()
	loopingDemo()
	orderedIterationDemo()
	mapOfStructsDemo()
	countingDemo()
	nestedMapsDemo()
}

// ==============================================================
// 1. NIL MAP
// ==============================================================
//
// var m map[string]int creates a nil map.
// Reading from a nil map is safe — returns zero value.
// WRITING to a nil map PANICS — always use make() or a literal before writing.

func nilMapDemo() {
	fmt.Println("--- Nil Map ---")
	var nilMap map[string]int
	fmt.Println("nilMap:", nilMap)        // map[]
	fmt.Println("is nil:", nilMap == nil) // true
	fmt.Println("length:", len(nilMap))   // 0
	// nilMap["key"] = 1  ← PANIC: assignment to entry in nil map
}

// ==============================================================
// 2. make()
// ==============================================================
//
// make(map[KeyType]ValueType) creates an empty map ready for use.
// map[string]int means:
//   - keys must be strings
//   - values must be ints

func makeMapDemo() {
	fmt.Println("\n--- make() Map ---")
	personAge := make(map[string]int)

	personAge["Alice"] = 30
	personAge["Bob"] = 25
	personAge["Charlie"] = 35

	fmt.Println("map:", personAge)
	fmt.Println("length:", len(personAge))
	fmt.Println("Alice's age:", personAge["Alice"])
}

// ==============================================================
// 3. MAP LITERAL
// ==============================================================
//
// Define and populate in one shot — the most common pattern.

func mapLiteralDemo() {
	fmt.Println("\n--- Map Literal ---")
	capitals := map[string]string{
		"Nigeria": "Abuja",
		"France":  "Paris",
		"Japan":   "Tokyo",
	}
	fmt.Println("capitals:", capitals)
	fmt.Println("capital of Japan:", capitals["Japan"])
}

// ==============================================================
// 4. CHECKING IF A KEY EXISTS
// ==============================================================
//
// If you access a missing key, Go silently returns the zero value — no error, no crash.
// This can cause silent bugs. Always use the two-value form when unsure:
//
//   value, ok := m[key]
//   ok == true  → key exists
//   ok == false → key does not exist

func keyExistenceDemo() {
	fmt.Println("\n--- Key Existence Check ---")
	capitals := map[string]string{
		"Nigeria": "Abuja",
		"France":  "Paris",
	}

	// UNSAFE — returns "" silently, no indication the key is missing
	unknown := capitals["Brazil"]
	fmt.Println("unsafe lookup for Brazil:", unknown) // empty string — misleading!

	// SAFE — always do this when unsure
	if capital, ok := capitals["Nigeria"]; ok {
		fmt.Println("found:", capital)
	}

	if _, ok := capitals["Brazil"]; !ok {
		fmt.Println("Brazil is not in the map")
	}
}

// ==============================================================
// 5. ADD, UPDATE, DELETE
// ==============================================================

func addUpdateDeleteDemo() {
	fmt.Println("\n--- Add, Update, Delete ---")
	m := map[string]string{
		"Nigeria": "Abuja",
		"France":  "Paris",
	}

	m["Japan"] = "Tokyo" // add new key
	fmt.Println("after add:", m)

	m["France"] = "Lyon" // update existing key
	fmt.Println("after update France:", m["France"])

	delete(m, "Nigeria") // delete a key
	fmt.Println("after delete Nigeria, length:", len(m))

	_, exists := m["Nigeria"]
	fmt.Println("Nigeria still exists:", exists) // false
}

// ==============================================================
// 6. LOOPING WITH range
// ==============================================================
//
// IMPORTANT: Maps are UNORDERED in Go by design.
// The order of range is random on every run — never rely on it.

func loopingDemo() {
	fmt.Println("\n--- Looping (unordered) ---")
	scores := map[string]int{
		"Alice":   95,
		"Bob":     87,
		"Charlie": 91,
	}

	for name, score := range scores {
		fmt.Printf("  %s: %d\n", name, score)
	}

	// Just the keys — ignore value with _
	fmt.Println("names only:")
	for name := range scores {
		fmt.Println(" ", name)
	}
}

// ==============================================================
// 7. ORDERED ITERATION
// ==============================================================
//
// Strategy 1 — manual: define a separate slice with your desired key order,
//   then range the slice and look up each key in the map.
//
// Strategy 2 — sort: extract all keys into a slice, sort it, then range it.

func orderedIterationDemo() {
	fmt.Println("\n--- Ordered Iteration ---")

	data := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	// Strategy 1: manual order guide
	order := []string{"one", "two", "three", "four"}
	fmt.Print("manual order: ")
	for _, key := range order {
		fmt.Printf("%s:%d  ", key, data[key])
	}
	fmt.Println()

	// Strategy 2: sort package
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Print("sorted order: ")
	for _, k := range keys {
		fmt.Printf("%s:%d  ", k, data[k])
	}
	fmt.Println()
}

// ==============================================================
// 8. MAP OF STRUCTS
// ==============================================================
//
// Values don't have to be primitives — you can map to structs.

func mapOfStructsDemo() {
	fmt.Println("\n--- Map of Structs ---")

	type Product struct {
		price   float64
		inStock bool
	}

	inventory := map[string]Product{
		"Laptop":  {price: 999.99, inStock: true},
		"Mouse":   {price: 25.00, inStock: true},
		"Monitor": {price: 399.00, inStock: false},
	}

	for item, details := range inventory {
		fmt.Printf("  %-10s | $%.2f | in stock: %t\n", item, details.price, details.inStock)
	}
}

// ==============================================================
// 9. COUNTING WITH MAPS
// ==============================================================
//
// A very common pattern. Go returns 0 for missing keys,
// so wordCount[word]++ automatically starts each word at 0 → 1.

func countingDemo() {
	fmt.Println("\n--- Counting Occurrences ---")
	words := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}

	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++ // first time: 0 + 1 = 1; next time: 1 + 1 = 2; etc.
	}

	for word, count := range wordCount {
		fmt.Printf("  %s: %d\n", word, count)
	}
}

// ==============================================================
// 10. NESTED MAPS
// ==============================================================
//
// A map whose values are themselves maps.
// Chain brackets to go deeper: countries["Nigeria"]["currency"]

func nestedMapsDemo() {
	fmt.Println("\n--- Nested Maps ---")

	countries := map[string]map[string]string{
		"Nigeria": {
			"capital":  "Abuja",
			"currency": "Naira",
		},
		"Japan": {
			"capital":  "Tokyo",
			"currency": "Yen",
		},
	}

	fmt.Println("Nigeria's currency:", countries["Nigeria"]["currency"])
	fmt.Println("Japan's capital:", countries["Japan"]["capital"])
}
