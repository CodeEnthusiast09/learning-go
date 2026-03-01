// Package errorplain
package errorplain

import (
	"errors"
	"fmt"
)

// ==============================================================
// ERRORS WITHOUT INTERFACES
// ==============================================================
//
// Sometimes you don't need fancy custom error types.
// A plain errors.New() or fmt.Errorf() is enough.
//
// The SavingsAccount here returns simple string-based errors.
// The caller gets the error message but can't inspect specific fields.
// That's fine when the message is all you need.
// ==============================================================

// SavingsAccount is a basic bank account.
type SavingsAccount struct {
	id      string
	balance float64
	locked  bool
}

// NewSavingsAccount creates a new account with an initial balance.
func NewSavingsAccount(id string, initialBalance float64) *SavingsAccount {
	return &SavingsAccount{id: id, balance: initialBalance}
}

func (a *SavingsAccount) ID() string       { return a.id }
func (a *SavingsAccount) Balance() float64 { return a.balance }

// Withdraw returns a plain error — just a message, no extra data attached.
func (a *SavingsAccount) Withdraw(amount float64) error {
	if a.locked {
		return errors.New("account is locked")
	}
	if amount > a.balance {
		return errors.New("insufficient balance")
	}
	a.balance -= amount
	return nil
}

// processWithdrawal calls Withdraw and prints the result.
// Notice: all we can do with the error is print its message — we can't
// inspect WHY it failed (locked? insufficient?) without parsing the string.
// That's the limitation this pattern has. See errorifaces for the solution.
func processWithdrawal(account *SavingsAccount, amount float64) {
	fmt.Printf("[Account: %s] withdrawing %.2f...\n", account.ID(), amount)

	if err := account.Withdraw(amount); err != nil {
		fmt.Println("  ✗ Error:", err)
		return
	}

	fmt.Printf("  ✓ Success! New balance: %.2f\n", account.Balance())
}

// Run demonstrates simple error-as-string patterns.
func Run() {
	fmt.Println("--- Errors Without Interfaces ---")

	alice := NewSavingsAccount("ACC-001", 500.00)
	bob := NewSavingsAccount("ACC-002", 100.00)
	bob.locked = true

	processWithdrawal(alice, 200.00) // success
	processWithdrawal(alice, 400.00) // insufficient balance (300 left after first withdrawal)
	processWithdrawal(bob, 50.00)    // account is locked
}
