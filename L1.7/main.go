package main

import (
	"fmt"
	"sync"
)

func main() {
	// There's no record order, whatsoever.
	ordinaryMap()
	syncMap()
}

func ordinaryMap() {
	// ordinaryMap is a function that uses a normal map and a mutex to protect access to the map.
	// It starts 5 goroutines and each goroutine adds an item to the map.
	// The map is then printed out.
	fmt.Println("ordinaryMap")
	card := make(map[int]string)

	// Create a mutex to protect access to the map
	var mu sync.Mutex
	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Start 5 goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			// Each goroutine adds an item to the map.
			defer wg.Done()
			// It locks the mutex before writing to the map.
			// And unlocks the mutex after writing to the map.
			mu.Lock()
			card[i] = fmt.Sprintf("Value: %d", i*10)
			mu.Unlock()
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	// Print out the map
	for i, v := range card {
		fmt.Printf("[%d]: {%s}\n", i, v)
	}
}

func syncMap() {
	// syncMap is a function that uses a sync.Map and does not need a mutex to protect access to the map.
	fmt.Println("syncMap")
	var card sync.Map
	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Start 5 goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			// Each goroutine adds an item to the map
			defer wg.Done()
			card.Store(id, fmt.Sprintf("Value: %d", id*20))
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	// Print out the map
	card.Range(func(k, v interface{}) bool {
		fmt.Printf("[%d]: {%s}\n", k, v)
		return true
	})
}

/*
 - Output: -
ordinaryMap
[5]: {Value: 50}
[1]: {Value: 10}
[2]: {Value: 20}
[3]: {Value: 30}
[4]: {Value: 40}
syncMap
[2]: {Value: 40}
[3]: {Value: 60}
[4]: {Value: 80}
[5]: {Value: 100}
[1]: {Value: 20}
*/