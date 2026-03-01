// Package errorifaces contains topics about error interfaces
package errorifaces

import (
	"errors"
	"fmt"
)

// ==============================================================
// ERRORS WITH INTERFACES & CUSTOM ERROR TYPES
// ==============================================================
//
// The plain errors.New() from errorplain only gives you a message.
// But what if the caller needs MORE information?
//
//   "insufficient funds" → ok but HOW short are we?
//   "account is locked" → ok but WHICH account?
//
// Solution: create custom error types (structs) that carry data.
// To be used as an error, a type just needs one method:
//
//   Error() string
//
// That satisfies the built-in `error` interface.
// Then the caller can use errors.As() to unwrap the specific type
// and access those extra fields.
// ==============================================================

// InsufficientFundsError carries the amounts involved in the failure.
type InsufficientFundsError struct {
	Requested float64
	Available float64
}

// Error satisfies the built-in error interface.
// Once a type has this method, it IS an error — no explicit declaration needed.
func (e *InsufficientFundsError) Error() string {
	return fmt.Sprintf(
		"insufficient funds: tried %.2f but only %.2f available",
		e.Requested, e.Available,
	)
}

// AccountLockedError carries the account ID that's locked.
type AccountLockedError struct {
	AccountID string
}

func (e *AccountLockedError) Error() string {
	return fmt.Sprintf("account %s is locked", e.AccountID)
}

// ==============================================================
// THE INTERFACE
// ==============================================================
//
// BankAccount defines the contract. Any type that has these three
// methods automatically satisfies BankAccount — no declaration needed.

type BankAccount interface {
	Withdraw(amount float64) error
	Balance() float64
	ID() string
}

// ==============================================================
// CONCRETE IMPLEMENTATION
// ==============================================================

// SavingsAccount implements BankAccount.
type SavingsAccount struct {
	id      string
	balance float64
	locked  bool
}

// NewSavingsAccount creates a new account.
func NewSavingsAccount(id string, initialBalance float64) *SavingsAccount {
	return &SavingsAccount{id: id, balance: initialBalance}
}

func (a *SavingsAccount) ID() string       { return a.id }
func (a *SavingsAccount) Balance() float64 { return a.balance }

// Withdraw returns a TYPED error — the caller can inspect the specific failure.
func (a *SavingsAccount) Withdraw(amount float64) error {
	if a.locked {
		return &AccountLockedError{AccountID: a.id}
	}
	if amount > a.balance {
		return &InsufficientFundsError{Requested: amount, Available: a.balance}
	}
	a.balance -= amount
	return nil
}

// ==============================================================
// PROCESSING WITH RICH ERROR INSPECTION
// ==============================================================
//
// errors.As(err, &target) checks if err (or anything it wraps) matches the
// type of target. If yes, it fills target with that specific error value
// so you can read its fields.
//
// This is better than a plain type switch for two reasons:
//   1. It handles WRAPPED errors (errors wrapped with fmt.Errorf("%w", err))
//   2. It's more idiomatic Go

func processWithdrawal(account BankAccount, amount float64) {
	fmt.Printf("[Account: %s] withdrawing %.2f...\n", account.ID(), amount)

	err := account.Withdraw(amount)
	if err == nil {
		fmt.Printf("  ✓ Success! New balance: %.2f\n", account.Balance())
		return
	}

	var insufficientErr *InsufficientFundsError
	var lockedErr *AccountLockedError

	switch {
	case errors.As(err, &insufficientErr):
		// We can access Requested and Available because errors.As filled insufficientErr
		shortfall := insufficientErr.Requested - insufficientErr.Available
		fmt.Printf("  ✗ Short by %.2f\n", shortfall)

	case errors.As(err, &lockedErr):
		fmt.Printf("  ✗ Account %s is locked — contact support\n", lockedErr.AccountID)

	default:
		fmt.Printf("  ✗ Unexpected error: %v\n", err)
	}
}

// Run demonstrates typed errors with rich inspection.
func Run() {
	fmt.Println("--- Errors With Interfaces (typed errors) ---")

	alice := NewSavingsAccount("ACC-001", 500.00)
	bob := NewSavingsAccount("ACC-002", 100.00)
	bob.locked = true

	processWithdrawal(alice, 200.00) // success
	processWithdrawal(alice, 400.00) // insufficient funds (300 left)
	processWithdrawal(bob, 50.00)    // locked
}
