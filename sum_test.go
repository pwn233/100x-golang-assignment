package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSumFunctionMatchOriginal(t *testing.T) {
	// initial x and y for testing both optimized and original func
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	size := 10000
	x, y := make([]int, size), make([]int, size)
	for i := 0; i < size; i++ {
		x[i] = r.Intn(100)
		y[i] = r.Intn(100)
	}
	// optimized method for finding sum
	testResult := sum(x, y)
	// original method for finding sum from the assignment
	var original []int
	for _, i := range x {
		for _, j := range y {
			original = append(original, i+j)
		}
	}
	// check length for both result (suppose to match)
	if len(original) != len(testResult) {
		t.Error("original and z lengths are not matched\t: [", len(original), "!= ", len(testResult), "]")
		return
	} else {
		fmt.Println("original and z lengths are matched\t: [", len(original), "=", len(testResult), "]")
	}
	// check values for both result (suppose to match), in case any mistake or misvalue will count as falseCount to calculate percentage
	var falseCount int
	for i := range original {
		if original[i] != testResult[i] {
			falseCount++
		}
	}
	if falseCount > 0 {
		t.Error("original and z values are matched only\t: [", (float64(len(original)-falseCount)/float64(len(original)))*100, "% ]")
		return
	} else {
		fmt.Println("original and z values are matched\t: [", (float64(len(original)-falseCount)/float64(len(original)))*100, "% ]")
	}
}
