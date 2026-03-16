// Package channels
package channels

import (
	"fmt"
	"sync"
	"time"
)

// Run demonstrates all channel and goroutine patterns.
func Run() {
	fmt.Println("--- 1. Basic Goroutine ---")
	basicGoroutine()

	fmt.Println("\n--- 2. Unbuffered Channel ---")
	unbufferedChannel()

	fmt.Println("\n--- 3. Buffered Channel ---")
	bufferedChannel()

	fmt.Println("\n--- 4. Directional Channels ---")
	directionalChannels()

	fmt.Println("\n--- 5. Range Over Channel ---")
	rangeOverChannel()

	fmt.Println("\n--- 6. Select Statement ---")
	selectStatement()

	fmt.Println("\n--- 7. Done Channel Pattern ---")
	donePattern()

	fmt.Println("\n--- 8. Fan-Out Pattern ---")
	fanOut()
}

// ==============================================================
// 1. BASIC GOROUTINE
// ==============================================================
//
// A goroutine is a lightweight thread managed by Go's runtime.
// Start one with the `go` keyword — it runs concurrently with everything else.
// wg.Go() (Go 1.22+) is a convenient way to launch a goroutine and
// automatically track it in the WaitGroup at the same time.

func basicGoroutine() {
	var wg sync.WaitGroup

	wg.Go(func() {
		fmt.Println("Hello from a goroutine!")
	})

	wg.Wait() // block here until all goroutines tracked by wg are done
	fmt.Println("Main continues after goroutine")
}

// ==============================================================
// 2. UNBUFFERED CHANNEL
// ==============================================================
//
// make(chan T) creates an unbuffered channel.
//
// UNBUFFERED = synchronous hand-off:
//   - The SENDER blocks until someone RECEIVES
//   - The RECEIVER blocks until someone SENDS
//   - Both must "meet" at the same time — like passing a note directly
//
// Syntax:
//   ch <- value     → send a value into the channel
//   value := <-ch   → receive a value from the channel

func unbufferedChannel() {
	ch := make(chan string) // unbuffered channel of strings

	go func() {
		fmt.Println("goroutine: about to send...")
		ch <- "hello from goroutine" // BLOCKS here until main receives
		fmt.Println("goroutine: sent!")
	}()

	time.Sleep(50 * time.Millisecond) // let goroutine start first
	msg := <-ch                       // BLOCKS here until goroutine sends
	fmt.Println("main received:", msg)
}

// ==============================================================
// 3. BUFFERED CHANNEL
// ==============================================================
//
// make(chan T, N) creates a buffered channel with capacity N.
//
// BUFFERED = asynchronous, up to a point:
//   - The sender only blocks when the buffer is FULL
//   - The receiver only blocks when the buffer is EMPTY
//   - Like a mailbox — drop the letter in and walk away

func bufferedChannel() {
	ch := make(chan int, 3) // buffer holds up to 3 ints

	// We can send 3 values without any goroutine receiving — buffer absorbs them
	ch <- 10
	ch <- 20
	ch <- 30
	// ch <- 40 would BLOCK — buffer is full

	fmt.Println("receive 1:", <-ch) // 10
	fmt.Println("receive 2:", <-ch) // 20
	fmt.Println("receive 3:", <-ch) // 30

	ch2 := make(chan int, 5)
	ch2 <- 1
	ch2 <- 2
	fmt.Printf("len (items in buffer): %d  cap (total buffer size): %d\n", len(ch2), cap(ch2))
}

// ==============================================================
// 4. DIRECTIONAL CHANNELS
// ==============================================================
//
// You can restrict a channel to send-only or receive-only.
// This makes your code clearer — functions declare exactly what they do.
//
//   chan<- T   →  send-only  (can only put things in)
//   <-chan T   →  receive-only (can only take things out)

func producer(ch chan<- string) { // ONLY allowed to send
	ch <- "message from producer"
	close(ch) // close signals: no more values will be sent
}

func consumer(ch <-chan string) { // ONLY allowed to receive
	msg := <-ch
	fmt.Println("consumer got:", msg)
}

func directionalChannels() {
	ch := make(chan string, 1)
	producer(ch)
	consumer(ch)
}

// ==============================================================
// 5. RANGE OVER CHANNEL
// ==============================================================
//
// You can range over a channel — it keeps receiving until the channel is CLOSED.
// If you forget close(), range loops forever → deadlock.

func rangeOverChannel() {
	ch := make(chan int, 5)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i * i // send squares: 1, 4, 9, 16, 25
		}
		close(ch) // MUST close — tells range "nothing more is coming"
	}()

	for val := range ch {
		fmt.Print(val, " ")
	}
	fmt.Println("\n  channel closed, loop done")
}

// ==============================================================
// 6. SELECT STATEMENT
// ==============================================================
//
// select is like a switch but for channels.
// It waits on multiple channels and picks whichever is ready first.
// If multiple are ready at the same time, it picks randomly.
// A default case makes it non-blocking — runs immediately if nothing is ready.

func selectStatement() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "from ch1"
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch2 <- "from ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("select picked:", msg)
		case msg := <-ch2:
			fmt.Println("select picked:", msg)
		}
	}

	// Non-blocking select — default runs immediately if no channel is ready
	ch3 := make(chan int)
	select {
	case val := <-ch3:
		fmt.Println("got:", val)
	default:
		fmt.Println("nothing ready — default ran")
	}
}

// ==============================================================
// 7. DONE CHANNEL PATTERN
// ==============================================================
//
// A common way to signal a goroutine to stop.
// We send struct{} because it takes ZERO bytes — it's just a signal, no data.
// close(done) broadcasts the signal to ALL goroutines listening on it.

func donePattern() {
	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("worker: received stop signal, exiting")
				return
			default:
				time.Sleep(20 * time.Millisecond) // simulated work
			}
		}
	}()

	time.Sleep(60 * time.Millisecond)
	close(done)                       // broadcast stop to all listeners
	time.Sleep(30 * time.Millisecond) // give goroutine time to print
}

// ==============================================================
// 8. FAN-OUT PATTERN
// ==============================================================
//
// One producer sends jobs. Multiple workers pick them up concurrently.
// This is how you parallelize work in Go.

func fanOut() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	var wg sync.WaitGroup

	// Start 3 worker goroutines
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs { // each worker picks up jobs as they arrive
				result := job * job
				fmt.Printf("worker %d: %d² = %d\n", workerID, job, result)
				results <- result
			}
		}(w)
	}

	// Send 9 jobs
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs) // no more jobs — workers will exit their range loop

	// Wait for all workers to finish, then close results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect all results
	total := 0
	for r := range results {
		total += r
	}
	fmt.Println("total of all squares:", total) // 1+4+9+16+25+36+49+64+81 = 285
}
