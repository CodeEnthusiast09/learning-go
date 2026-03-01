// Package enums
package enums

import (
	"fmt"
)

// ==============================================================
// ENUMS IN GO
// ==============================================================
//
// Go does NOT have a native `enum` keyword like Java, TypeScript, or C.
//
// Instead, Go uses a combination of:
//   1. A custom type (based on int or string)
//   2. const block with iota
//   3. A String() method to make it human-readable
//
// This pattern gives you everything an enum does — and it's idiomatic Go.
// ==============================================================

// ==============================================================
// PART 1: THE BASIC PATTERN — iota
// ==============================================================
//
// iota is a special constant that auto-increments inside a const block.
// It starts at 0 and increases by 1 for each constant.
//
// Step 1: Create a named type (so your constants have a real type, not just int)
// Step 2: Define constants using iota

type Direction int

const (
	North Direction = iota // 0
	East                   // 1 (iota increments automatically)
	South                  // 2
	West                   // 3
)

// String method implemented by Direction
// Step 3: Add a String() method so it prints "North" instead of "0"
// This implements the fmt.Stringer interface automatically.
func (d Direction) String() string {
	switch d {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	default:
		return "Unknown Direction"
	}
}

// ==============================================================
// PART 2: SKIPPING VALUES WITH iota
// ==============================================================
//
// You can do math with iota to skip or transform values.

type ByteSize float64

const (
	_           = iota             // blank identifier — skip 0 (iota = 0)
	KB ByteSize = 1 << (10 * iota) // 1 << 10 = 1024   (iota = 1)
	MB                             // 1 << 20           (iota = 2)
	GB                             // 1 << 30           (iota = 3)
	TB                             // 1 << 40           (iota = 4)
)

func (b ByteSize) String() string {
	switch {
	case b >= TB:
		return fmt.Sprintf("%.2f TB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2f GB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2f MB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2f KB", b/KB)
	}
	return fmt.Sprintf("%.2f B", b)
}

// ==============================================================
// PART 3: STRING-BASED ENUM
// ==============================================================
//
// When you want the values to be meaningful strings, not numbers.
// Useful for JSON serialization, configs, and readable logs.

type Status string

const (
	StatusPending   Status = "pending"
	StatusActive    Status = "active"
	StatusSuspended Status = "suspended"
	StatusClosed    Status = "closed"
)

// String() is optional here since Status IS already a string,
// but you can still add methods:

func (s Status) IsTerminal() bool {
	return s == StatusClosed
}

func (s Status) CanTransitionTo(next Status) bool {
	switch s {
	case StatusPending:
		return next == StatusActive
	case StatusActive:
		return next == StatusSuspended || next == StatusClosed
	case StatusSuspended:
		return next == StatusActive || next == StatusClosed
	case StatusClosed:
		return false // terminal — no transitions out
	}
	return false
}

// ==============================================================
// PART 4: USING ENUMS IN FUNCTIONS AND SWITCHES
// ==============================================================
//
// This is where the typed constant pays off.
// The compiler will warn you if you pass a wrong type.

func Describe(d Direction) string {
	switch d {
	case North:
		return "Heading up"
	case South:
		return "Heading down"
	case East:
		return "Heading right"
	case West:
		return "Heading left"
	default:
		return "Unknown direction"
	}
}

// ==============================================================
// PART 5: BITMASK / FLAGS ENUM
// ==============================================================
//
// A useful pattern where each constant is a power of 2 (a single bit).
// You can COMBINE flags using | (bitwise OR)
// and CHECK flags using &  (bitwise AND)
//
// Common in permissions, feature flags, file modes.

type Permission int

const (
	PermRead    Permission = 1 << iota // 1  (binary: 001)
	PermWrite                          // 2  (binary: 010)
	PermExecute                        // 4  (binary: 100)
)

func (p Permission) String() string {
	var result string
	if p&PermRead != 0 {
		result += "r"
	} else {
		result += "-"
	}
	if p&PermWrite != 0 {
		result += "w"
	} else {
		result += "-"
	}
	if p&PermExecute != 0 {
		result += "x"
	} else {
		result += "-"
	}
	return result
}

func (p Permission) Has(flag Permission) bool {
	return p&flag != 0
}

// ==============================================================
// PART 6: EXHAUSTIVE SWITCH CHECK
// ==============================================================
//
// Go doesn't force you to handle all enum values in a switch.
// But you should — unhandled cases are bugs waiting to happen.
//
// Pattern: use a default that panics during development.
// In production, default should probably log + return an error.

type OrderState int

const (
	OrderCreated OrderState = iota
	OrderPaid
	OrderShipped
	OrderDelivered
	OrderCancelled
)

func (o OrderState) String() string {
	names := [...]string{"Created", "Paid", "Shipped", "Delivered", "Cancelled"}
	if int(o) >= len(names) {
		return fmt.Sprintf("OrderState(%d)", o)
	}
	return names[o]
}

func ProcessOrder(state OrderState) {
	switch state {
	case OrderCreated:
		fmt.Println("Waiting for payment...")
	case OrderPaid:
		fmt.Println("Preparing shipment...")
	case OrderShipped:
		fmt.Println("In transit...")
	case OrderDelivered:
		fmt.Println("Delivered!")
	case OrderCancelled:
		fmt.Println("Order cancelled.")
	default:
		// This fires if someone adds a new OrderState and forgets to handle it here
		panic(fmt.Sprintf("unhandled OrderState: %v", state))
	}
}

// ==============================================================
// PART 7: VALIDATION
// ==============================================================
//
// Since anyone could do  Direction(99), it's good practice to add
// an IsValid() method to verify the value is a known enum value.

func (d Direction) IsValid() bool {
	switch d {
	case North, East, South, West:
		return true
	}
	return false
}

// ==============================================================
// MAIN
// ==============================================================

func Run() {
	fmt.Println("=== DIRECTION ENUM ===")
	d := North
	fmt.Println("Direction:", d) // prints "North" not 0
	fmt.Println("Describe:", Describe(d))
	fmt.Println("East is valid:", East.IsValid())
	fmt.Println("Direction(99) is valid:", Direction(99).IsValid())

	fmt.Println("\n=== BYTE SIZES ===")
	fmt.Println("KB:", KB)
	fmt.Println("MB:", MB)
	fmt.Println("GB:", GB)
	fileSize := ByteSize(1536) * KB
	fmt.Println("File size:", fileSize)

	fmt.Println("\n=== STRING STATUS ENUM ===")
	account := StatusPending
	fmt.Println("Status:", account)
	fmt.Println("Can go Active:", account.CanTransitionTo(StatusActive)) // true
	fmt.Println("Can go Closed:", account.CanTransitionTo(StatusClosed)) // false (pending → active only)
	account = StatusActive
	fmt.Println("After activate:", account)
	fmt.Println("Is terminal:", account.IsTerminal())             // false
	fmt.Println("Closed is terminal:", StatusClosed.IsTerminal()) // true

	fmt.Println("\n=== BITMASK PERMISSIONS ===")
	// Combine permissions with |
	readWrite := PermRead | PermWrite
	fmt.Println("readWrite:", readWrite)                    // "rw-"
	fmt.Println("Has read:", readWrite.Has(PermRead))       // true
	fmt.Println("Has execute:", readWrite.Has(PermExecute)) // false

	all := PermRead | PermWrite | PermExecute
	fmt.Println("All permissions:", all) // "rwx"

	readonly := PermRead
	fmt.Println("Readonly:", readonly) // "r--"

	fmt.Println("\n=== ORDER STATE ===")
	states := []OrderState{OrderCreated, OrderPaid, OrderShipped, OrderDelivered}
	for _, s := range states {
		fmt.Printf("  State %d (%s): ", s, s)
		ProcessOrder(s)
	}

	fmt.Println("\n=== NAMING CONVENTION REMINDER ===")
	// Go convention: name your constants like TypeValue, e.g.:
	//   Direction + North = DirectionNorth? No.
	//   Actually Go convention is plain:  North, East (if type name is clear)
	//   OR prefixed:  DirectionNorth, DirectionEast (if context is ambiguous)
	//
	// For Status:  StatusPending, StatusActive (prefix helps because "Pending" alone is vague)
	// For Direction: North, East is fine (Direction type makes it obvious)
	fmt.Println("Convention: use prefix when the constant name is ambiguous alone")
}
