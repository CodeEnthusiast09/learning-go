// Package mutexes
package mutexes

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Run demonstrates all mutex and sync primitives.
func Run() {
	fmt.Println("--- 1. Race Condition (the problem) ---")
	raceConditionDemo()

	fmt.Println("\n--- 2. Mutex Fix ---")
	mutexFix()

	fmt.Println("\n--- 3. RWMutex — multiple readers ---")
	rwMutexDemo()

	fmt.Println("\n--- 4. WaitGroup ---")
	waitGroupDemo()

	fmt.Println("\n--- 5. sync.Once ---")
	onceDemo()

	fmt.Println("\n--- 6. Atomic Operations ---")
	atomicDemo()

	fmt.Println("\n--- 7. Mutex embedded in a struct ---")
	safeCounterDemo()
}

// ==============================================================
// 1. RACE CONDITION
// ==============================================================
//
// A race condition happens when multiple goroutines read AND write
// the same variable at the same time without coordination.
//
// Goroutine A reads counter: 5
// Goroutine B reads counter: 5   (before A wrote back)
// Goroutine A writes:        6
// Goroutine B writes:        6   ← A's write was LOST
// Should be 7, got 6.
//
// Run with: go run -race .
// The -race flag makes Go detect race conditions at runtime.

func raceConditionDemo() {
	counter := 0
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // NOT safe — multiple goroutines can read/write at the same time
		}()
	}

	wg.Wait()
	// Should be 1000 but you'll often see less — the exact result is unpredictable.
	fmt.Println("  unsafe counter (should be 1000):", counter)
}

// ==============================================================
// 2. MUTEX FIX
// ==============================================================
//
// sync.Mutex has two methods:
//   Lock()   → "I'm in the critical section — everyone else wait"
//   Unlock() → "I'm done — next one can go"
//
// ALWAYS defer Unlock() right after Lock() — that way it runs even if the
// function panics, so the mutex is never left permanently locked (deadlock).

func mutexFix() {
	counter := 0
	var mu sync.Mutex // zero value is an unlocked mutex — ready to use immediately
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mu.Lock()         // only ONE goroutine gets past this at a time
			defer mu.Unlock() // released when this function returns
			counter++         // safe — we're guaranteed to be the only one here
		}()
	}

	wg.Wait()
	fmt.Println("  safe counter (always 1000):", counter)
}

// ==============================================================
// 3. RWMUTEX — MULTIPLE CONCURRENT READERS
// ==============================================================
//
// sync.RWMutex is smarter than a plain Mutex:
//   Lock()   / Unlock()   → write lock (exclusive — blocks everyone)
//   RLock()  / RUnlock()  → read lock  (shared — multiple readers at the same time)
//
// Use RWMutex when reads are frequent and writes are rare.
// With a plain Mutex, readers would needlessly block each other.

type cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func newCache() *cache {
	return &cache{data: make(map[string]string)}
}

func (c *cache) set(key, value string) {
	c.mu.Lock() // exclusive write lock — blocks all readers and writers
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *cache) get(key string) (string, bool) {
	c.mu.RLock() // shared read lock — multiple goroutines can RLock at the same time
	defer c.mu.RUnlock()
	val, ok := c.data[key]
	return val, ok
}

func rwMutexDemo() {
	c := newCache()
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		c.set("name", "Alice")
		c.set("city", "Lagos")
	}()
	wg.Wait() // ensure writes finish before readers start

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			if val, ok := c.get("name"); ok {
				fmt.Printf("  reader %d got: %s\n", n, val)
			}
		}(i)
	}
	wg.Wait()
}

// ==============================================================
// 4. WAITGROUP
// ==============================================================
//
//   Add(n)  → "I'm launching n goroutines, count them"
//   Done()  → "one goroutine is done" (decrements count by 1)
//   Wait()  → "block here until the count reaches zero"
//
// CRITICAL RULE: call Add() BEFORE starting the goroutine.
// If you call Add() inside the goroutine, Wait() might fire before Add() does.

func waitGroupDemo() {
	var wg sync.WaitGroup
	tasks := []string{"emails", "reports", "invoices"}

	for _, task := range tasks {
		wg.Add(1) // add BEFORE go func, not inside it
		go func(t string) {
			defer wg.Done()
			fmt.Println("  processing:", t)
		}(task) // pass task as argument — avoids the closure-capture bug
	}

	wg.Wait()
	fmt.Println("  all tasks done")
}

// ==============================================================
// 5. SYNC.ONCE
// ==============================================================
//
// sync.Once guarantees a function runs EXACTLY ONCE no matter how many
// goroutines call it. Perfect for initializing singletons (DB connection, config).

type appConfig struct {
	DSN string
}

var (
	cfg     *appConfig
	cfgOnce sync.Once
)

func getConfig() *appConfig {
	cfgOnce.Do(func() {
		// This block runs exactly once, even if 100 goroutines call getConfig()
		fmt.Println("  initializing config (runs only once)")
		cfg = &appConfig{DSN: "postgres://localhost:5432/mydb"}
	})
	return cfg
}

func onceDemo() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = getConfig()
		}()
	}

	wg.Wait()
	fmt.Println("  config DSN:", getConfig().DSN)
}

// ==============================================================
// 6. ATOMIC OPERATIONS
// ==============================================================
//
// For simple counters and flags, sync/atomic is FASTER than a Mutex.
// Atomic operations are at the hardware level — indivisible, no race possible.
//
// Use atomic for: counters, flags, single values
// Use Mutex for: anything more complex (structs, maps, slices)

func atomicDemo() {
	var counter int64 // must be int32 or int64 for atomic operations
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1) // indivisible increment — no mutex needed
		}()
	}

	wg.Wait()
	fmt.Println("  atomic counter (always 1000):", atomic.LoadInt64(&counter))
}

// ==============================================================
// 7. MUTEX EMBEDDED IN A STRUCT (real-world pattern)
// ==============================================================
//
// In real code, embed the mutex INSIDE the struct it protects.
// This keeps the lock and data together — harder to forget one without the other.

type safeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *safeCounter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *safeCounter) decrement() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value--
}

func (c *safeCounter) get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func safeCounterDemo() {
	sc := &safeCounter{}
	var wg sync.WaitGroup

	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sc.increment()
		}()
	}

	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sc.decrement()
		}()
	}

	wg.Wait()
	fmt.Println("  final counter (should be 300):", sc.get())
}
