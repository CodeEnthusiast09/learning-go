// Package pkgmodules contains topics - Packages and Modules
package pkgmodules

import (
	"fmt"
	"math"
	"strings"
)

// ==============================================================
// PACKAGES AND MODULES IN GO
// ==============================================================
//
// PACKAGE  → a folder of .go files that belong together.
//            Every .go file starts with: package <name>
//
// MODULE   → a collection of packages — your entire project.
//            Defined by a go.mod file at the root.
//
// Think of it like:
//   Module  = the whole bookshelf
//   Package = one book on the shelf
//   File    = a chapter inside the book
//
// ==============================================================
// EXPORTED vs UNEXPORTED (Go's visibility system)
// ==============================================================
//
//   Capitalized  = Exported   = visible OUTSIDE the package (public)
//   lowercase    = Unexported = visible ONLY inside the package (private)
//
// There is no `public` or `private` keyword in Go.
// Capitalization IS the access modifier.
// ==============================================================

// BankAccount has exported and unexported fields.
type BankAccount struct {
	Balance float64 // Exported — other packages can read/set this
	ownerID string  // unexported — only code inside this package can access it
}

// NewBankAccount is a constructor (exported).
// This is the ONLY way to set ownerID from outside the package.
// This is intentional encapsulation.
func NewBankAccount(owner string, balance float64) BankAccount {
	return BankAccount{
		Balance: balance,
		ownerID: owner, // we can set it because we're in the same package
	}
}

func (b BankAccount) getOwner() string { // unexported method
	return b.ownerID
}

// Owner is an exported method that exposes the unexported ownerID safely.
func (b BankAccount) Owner() string {
	return b.getOwner()
}

// ==============================================================
// INIT FUNCTION
// ==============================================================
//
// Each package can have an init() function.
// It runs automatically when the package is loaded — before main().
// You cannot call it manually. Used for setup like registering drivers.
// ==============================================================

var appName string

func init() {
	// This runs automatically when this package is first imported.
	// You'll see its effect in Run() below.
	appName = "Learning Go"
}

// Run demonstrates package/module concepts.
func Run() {
	fmt.Println("--- init() was called automatically ---")
	fmt.Println("  appName set by init():", appName)

	fmt.Println("\n--- Standard library packages ---")
	fmt.Println("  math.Pi:", math.Pi)
	fmt.Println("  math.Sqrt(144):", math.Sqrt(144))
	fmt.Println("  strings.ToUpper:", strings.ToUpper("hello from strings package"))
	fmt.Println("  strings.Contains:", strings.Contains("Lagos Nigeria", "Nigeria"))

	fmt.Println("\n--- Exported vs Unexported ---")
	account := NewBankAccount("alice-123", 500.00)
	fmt.Println("  account.Balance (exported field):", account.Balance)
	fmt.Println("  account.Owner() (exported method):", account.Owner())
	// account.ownerID → COMPILE ERROR: unexported field — try it to see Go enforce it

	fmt.Println("\n--- How module imports work ---")
	fmt.Println(`
  Import paths:
    "fmt"                          → stdlib (Go's built-in packages)
    "github.com/you/app/utils"     → your own package (module-relative)
    "github.com/gin-gonic/gin"     → external dependency

  go.mod commands:
    go mod init github.com/you/project   → create go.mod, start a module
    go get github.com/some/package       → download a dependency
    go mod tidy                          → remove unused, add missing deps

  Import aliases (when two packages have the same name):
    import (
        mrand "math/rand"
        crand "crypto/rand"
    )

  Side-effect imports (runs init() but you don't call it directly):
    import _ "github.com/lib/pq"   → registers the postgres driver
	`)
}
