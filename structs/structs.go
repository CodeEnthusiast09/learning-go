// Package structs
package structs

import "fmt"

// ==============================================================
// STRUCTS IN GO
// ==============================================================
//
// A struct (short for "structure") groups related data of different types
// into a single named unit — like a custom type you design yourself.
//
// Think of it like a form:
//   A Person form has Name (string), Age (int), Job (string), Salary (int).
//   Every person you create fills in that same form.
// ==============================================================

// Person represents a human with basic personal and professional details.
type Person struct {
	name   string // lowercase = unexported (only accessible inside this package)
	age    int
	job    string
	salary int
}

// Address represents a physical location.
type Address struct {
	city    string
	country string
}

// PersonWithAddress is a Person that also carries an Address.
// Embedding a struct inside another struct is called COMPOSITION.
// You access nested fields by chaining dots: p.address.city
type PersonWithAddress struct {
	name    string
	age     int
	address Address // nested struct — Address is a field whose type is another struct
}

// ==============================================================
// METHODS
// ==============================================================
//
// A method is a function attached to a type.
// The (p Person) before the function name is called the RECEIVER.
// It says: "this Greet function belongs to Person; inside it, that person is called p."

// Greet returns a greeting string using a VALUE receiver.
// Value receiver = Go passes a COPY of the Person. Greet cannot modify the original.
// Use value receivers when you're only reading data.
func (p Person) Greet() string {
	return "Hi, I am " + p.name + " and I work as a " + p.job
}

// getRaise uses a POINTER receiver (*Person).
// Pointer receiver = Go works on the REAL struct in memory, not a copy.
// Use pointer receivers when you need to MODIFY the struct.
func (p *Person) getRaise(amount int) {
	p.salary += amount // changes the ACTUAL salary, not a copy
}

// NewPerson is a constructor function — the Go convention for creating
// a struct with validation or setup logic.
// Returns a pointer so callers work on the same instance in memory.
func NewPerson(name, job string, age, salary int) *Person {
	return &Person{name: name, age: age, job: job, salary: salary}
}

// ==============================================================
// RUN
// ==============================================================

func Run() {
	fmt.Println("--- Basic Struct: var then assign ---")
	var pers1 Person
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000
	printPerson(pers1)

	fmt.Println("\n--- Struct Literal (most common way) ---")
	// All fields assigned at once. Named fields — order doesn't matter.
	pers2 := Person{
		name:   "John",
		age:    32,
		job:    "Engineer",
		salary: 8000,
	}
	printPerson(pers2)

	fmt.Println("\n--- Constructor function ---")
	pers3 := NewPerson("Amara", "Developer", 28, 7000)
	printPerson(*pers3)

	fmt.Println("\n--- Method: Greet() ---")
	fmt.Println(pers2.Greet())

	fmt.Println("\n--- Pointer Receiver: getRaise() ---")
	fmt.Println("salary before raise:", pers2.salary)
	pers2.getRaise(1500)
	fmt.Println("salary after raise:", pers2.salary)

	fmt.Println("\n--- Nested Struct ---")
	pers4 := PersonWithAddress{
		name: "Bola",
		age:  30,
		address: Address{
			city:    "Lagos",
			country: "Nigeria",
		},
	}
	fmt.Printf("%s lives in %s. %s. she is %v years old.\n", pers4.name, pers4.address.city, pers4.address.country, pers4.age)

	fmt.Println("\n--- Anonymous Struct ---")
	// Defined inline — no reusable type name. Good for one-off temporary groupings.
	product := struct {
		name  string
		price float64
	}{
		name:  "Keyboard",
		price: 49.99,
	}
	fmt.Printf("%s costs $%.2f\n", product.name, product.price)

	fmt.Println("\n--- Struct Comparison ---")
	// Two struct values are equal if ALL their fields are equal.
	a := Person{name: "Hege", age: 45, job: "Teacher", salary: 6000}
	b := Person{name: "Hege", age: 45, job: "Teacher", salary: 6000}
	c := Person{name: "John", age: 32, job: "Engineer", salary: 8000}
	fmt.Println("a == b:", a == b) // true
	fmt.Println("a == c:", a == c) // false
}

func printPerson(p Person) {
	fmt.Printf("  Name: %s | Age: %d | Job: %s | Salary: %d\n",
		p.name, p.age, p.job, p.salary)
}
