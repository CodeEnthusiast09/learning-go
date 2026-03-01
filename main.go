package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"learning-go/topics/channels"
	"learning-go/topics/enums"
	"learning-go/topics/errordemo"
	"learning-go/topics/errorifaces"
	"learning-go/topics/errorplain"
	"learning-go/topics/functions"
	"learning-go/topics/generics"
	"learning-go/topics/helloworld"
	"learning-go/topics/interfaces"
	"learning-go/topics/mapdemo"
	"learning-go/topics/mutexes"
	"learning-go/topics/pkgmodules"
	"learning-go/topics/pointers"
	"learning-go/topics/sprint"
	"learning-go/topics/structs"
)

// topic pairs a display label with the Run function for that topic.
type topic struct {
	label string
	run   func()
}

func main() {
	topics := []topic{
		{"Hello World (variables, constants, fmt, arrays, slices, loops)", helloworld.Run},
		{"Functions", functions.Run},
		{"Structs", structs.Run},
		{"Pointers", pointers.Run},
		{"Interfaces", interfaces.Run},
		{"Maps", mapdemo.Run},
		{"Errors (basic)", errordemo.Run},
		{"Errors without interfaces", errorplain.Run},
		{"Errors with interfaces", errorifaces.Run},
		{"Channels & Goroutines", channels.Run},
		{"Mutexes & Sync", mutexes.Run},
		{"Generics", generics.Run},
		{"Packages & Modules", pkgmodules.Run},
		{"Sprint — fmt (Sprint / Sprintf / Sprintln)", sprint.Run},
		{"Enums", enums.Run},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		printMenu(topics)

		fmt.Print("\nEnter a number (or 0 to run ALL, q to quit): ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == "q" || input == "Q" {
			fmt.Println("Goodbye!")
			return
		}

		if input == "0" {
			fmt.Println("\n========== RUNNING ALL TOPICS ==========")
			for _, t := range topics {
				printSectionHeader(t.label)
				t.run()
			}
			fmt.Println("\n========== DONE ==========")
			pauseForEnter(scanner)
			continue
		}

		n, err := strconv.Atoi(input)
		if err != nil || n < 1 || n > len(topics) {
			fmt.Printf("  ✗ Invalid choice. Enter a number between 1 and %d, 0, or q.\n", len(topics))
			continue
		}

		selected := topics[n-1]
		printSectionHeader(selected.label)
		selected.run()
		pauseForEnter(scanner)
	}
}

func printMenu(topics []topic) {
	fmt.Println("\n========================================")
	fmt.Println("       GO LEARNING — TOPIC MENU")
	fmt.Println("========================================")
	for i, t := range topics {
		fmt.Printf("  %2d. %s\n", i+1, t.label)
	}
	fmt.Println("----------------------------------------")
	fmt.Println("   0. Run ALL topics")
	fmt.Println("   q. Quit")
	fmt.Println("========================================")
}

func printSectionHeader(label string) {
	fmt.Printf("\n========== %s ==========\n\n", strings.ToUpper(label))
}

func pauseForEnter(scanner *bufio.Scanner) {
	fmt.Print("\nPress Enter to return to the menu...")
	scanner.Scan()
}
