// Package interfaces
package interfaces

import "fmt"

// ==============================================================
// INTERFACES IN GO
// ==============================================================
//
// An interface defines a SET OF METHOD SIGNATURES (a contract).
// Any type that has ALL those methods automatically satisfies the interface.
// You don't write "implements Speaker" — Go figures it out for you.
//
// This is called IMPLICIT implementation — very different from Java/C#
// where you have to explicitly declare what you implement.
//
// Why is this powerful?
//   You can write a function that accepts a Speaker, and it will work
//   with Dog, Human, Robot, or ANY future type that has Speak() —
//   even types that didn't exist when you wrote the function.
// ==============================================================

// Speaker is the interface — the "job description".
// Any type that has a Speak() method returning string is a Speaker.
type Speaker interface {
	Speak() string
}

// ==============================================================
// CONCRETE TYPES
// These are the actual "workers" that satisfy the Speaker interface.
// ==============================================================

// Dog has a Speak() method → automatically satisfies Speaker.
type Dog struct{ Name string }

func (d Dog) Speak() string { return d.Name + " says: Woof!" }

// Human has a Speak() method → automatically satisfies Speaker.
type Human struct{ Name string }

func (h Human) Speak() string { return h.Name + " says: Hello!" }

// Robot has a Speak() method → automatically satisfies Speaker.
type Robot struct{ Model string }

func (r Robot) Speak() string { return r.Model + " says: BEEP BOOP." }

// ==============================================================
// FUNCTION USING THE INTERFACE
// ==============================================================
//
// makeItSpeak doesn't care if it receives a Dog, Human, or Robot.
// It only cares: "does this thing have a Speak() method?"
// This is the power of interfaces — write the function ONCE,
// it works with every type that satisfies the interface.

func makeItSpeak(s Speaker) {
	fmt.Println(s.Speak())
}

// Run demonstrates interface basics.
func Run() {
	fmt.Println("--- Calling makeItSpeak with different types ---")
	dog := Dog{Name: "Rex"}
	human := Human{Name: "Alice"}
	robot := Robot{Model: "X-9000"}

	makeItSpeak(dog)
	makeItSpeak(human)
	makeItSpeak(robot)

	fmt.Println("\n--- Interface slice: mixed types, same behavior ---")
	// A slice of Speaker can hold Dog, Human, Robot — all at once.
	// This is called POLYMORPHISM — same method call, different behavior per type.
	speakers := []Speaker{dog, human, robot}
	for _, s := range speakers {
		makeItSpeak(s)
	}

	fmt.Println("\n--- Type assertion: getting the concrete type back ---")
	// Sometimes you have an interface value and need the underlying concrete type.
	// Type assertion: value.(ConcreteType)
	// Two-value form: val, ok := value.(ConcreteType)
	// ok is false if the assertion fails — use this to avoid panics.
	var s Speaker = Dog{Name: "Buddy"}
	if d, ok := s.(Dog); ok {
		fmt.Println("it's a Dog! Name:", d.Name)
	}

	fmt.Println("\n--- Type switch: handle each type differently ---")
	// A type switch is like a switch statement but checks the TYPE of the value.
	things := []Speaker{Dog{Name: "Max"}, Human{Name: "Bob"}, Robot{Model: "R2D2"}}
	for _, t := range things {
		switch v := t.(type) {
		case Dog:
			fmt.Println("Dog named", v.Name)
		case Human:
			fmt.Println("Human named", v.Name)
		case Robot:
			fmt.Println("Robot model", v.Model)
		}
	}
}
