// Package errordemo
package errordemo

// NOTE: This package is named "errordemo" (not "errors") because Go's standard
// library already has a package called "errors". Naming ours the same would
// shadow the stdlib one inside this package, making it impossible to use errors.New().

import "fmt"

// Run demonstrates basic error patterns in Go.
func Run() {
	fmt.Println("--- SMS Cost Errors ---")
	test(1.4, "+2348114957554")
	test(2.1, "+2348055069859")
	test(32.1, "+2348085029118")

	fmt.Println("--- sendSMS with error return ---")
	demonstrateSendSMS()
}

// getSMSErrorString builds a descriptive error message string.
// This is the simplest form of error handling — just format a message.
func getSMSErrorString(cost float64, recipient string) string {
	return fmt.Sprintf("SMS that costs ₦%.2f to be sent to %s cannot be sent", cost, recipient)
}

func test(cost float64, recipient string) {
	s := getSMSErrorString(cost, recipient)
	fmt.Println(s)
}

// ==============================================================
// ERRORS AS VALUES
// ==============================================================
//
// In Go, errors are just VALUES — a function returns (result, error).
// If error is nil → success.
// If error is non-nil → something went wrong, inspect the error.
//
// This forces you to think about failure at every step.
// There are no exceptions flying around invisibly — the error is right there.

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2

	if len(message) > maxTextLen {
		// fmt.Errorf creates a formatted error message — like Sprintf but returns an error
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}

	return costPerChar * len(message), nil // nil error means "all good"
}

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	customerCost, err := sendSMS(msgToCustomer)
	if err != nil {
		// Early return pattern — the moment we hit an error, we stop and bubble it up
		return 0, err
	}

	spouseCost, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0, err
	}

	return customerCost + spouseCost, nil
}

func demonstrateSendSMS() {
	// Success case
	cost, err := sendSMSToCouple("Hi, I love you", "Hey love, miss you")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Total SMS cost:", cost, "units")
	}

	// Failure case — one message too long
	cost2, err2 := sendSMSToCouple("Hi babyy, I love you so much today", "Hey")
	if err2 != nil {
		fmt.Println("Error:", err2) // triggered because first message > 25 chars
	} else {
		fmt.Println("Total SMS cost:", cost2, "units")
	}
}
