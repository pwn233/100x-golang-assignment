package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
	// "sync"
)

// unable to edit
func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	size := 10000
	x, y := make([]int, size), make([]int, size)
	for i := 0; i < size; i++ {
		x[i] = r.Intn(100)
		y[i] = r.Intn(100)
	}
	start := time.Now()
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)

	_ = sum(x, y)
	runtime.ReadMemStats(&m2)
	latency := time.Since(start)
	mem := m2.TotalAlloc - m1.TotalAlloc

	fmt.Println("Latency:", latency)
	fmt.Println("Memory:", mem, "bytes")
}

// TODO edit here only
func sum(x, y []int) []int {
	// fyi 1: I tried both ways, a lot of workers -> it got low latency but high memory usage (my computer automatically restart once I tried kub hah), so I picked the few sizes of workers instead kub (a bit more latency, but low memory usage which also computer not crashed kub)
	// fyi 2: Calculates time which it will perform as On, (workers)O^((len(x)/workers)*len(y))
	// fyi 3: I had tried once without waiting group and it run correctly (which if there was a hidden test data, It should not survived a large amount of data), which I assume that the workers  complete(go routine) before the main func end (even the sum func return to main).
	//        (I need to make sure for all workers done their works before return to main, so added waiting group for safety purpose -> a bit more latency increased (1 - 10 to 50 - 80))
	// create large container due to all of elements x will multiply with all of elements y. (the length must be len(x) * len(y) kub)
	var z []int = make([]int, len(x)*len(y))
	// declare a waiting group for waiting action that will be performed in go func (go routine) kub
	var wg sync.WaitGroup
	// initial workers amount (epocs or laps) for told waiting group to wait for go func (go routine) kub (in this case, I use 100 workers (go func) kub)
	var workers int = 100
	wg.Add(workers)
	// batch size per round
	var batchSize int = len(x) / workers
	// perform action of go func (go routine) in all workers here kub
	var xStart, xEnd int
	for i := 0; i < workers; i++ {
		// pick start and end for the batch (ex : 0 -> 1000, 1000 -> 2000) [reminder : when inside loop working, I used < to prevent duplicate likes 0 -> 999, 1000 -> 1999]
		xStart = i * batchSize
		xEnd = xStart + batchSize
		// start worker
		go func(start, xEnd int) {
			// defer means when go func ended performing action, it will trigger the waiting group to be "done", which will trigger the outside to continue action from "wait" state kub
			defer wg.Done()
			// index for order in z, starts like round*batchSize (ex : 0*1000 -> start 0, 1*1000 -> start 1000)
			var zIndex = start * len(x)
			// start 1st loop for running x value (which are only x values in the batch)
			for xIndex := start; xIndex < xEnd; xIndex++ {
				// start 2nd loop for running y value (which are all y values)
				for _, yValue := range y {
					// sum up values and store in z
					z[zIndex] = x[xIndex] + yValue
					// moving to next index
					zIndex++
				}
			}
		}(xStart, xEnd)
	}
	// told waiting group to wait for all the workers to done there job before return to main (from the original x and y size maybe not worried, but for some hidden test data with large amount maybe needed)
	wg.Wait()
	return z
}
