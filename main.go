package main

import (
	"fmt"
	"github.com/mdshahjahanmiah/go-profiling/bench"
	"github.com/mdshahjahanmiah/go-profiling/stack"
	"github.com/mdshahjahanmiah/go-profiling/utils"
	"net/http"
	_ "net/http/pprof" // Enable pprof HTTP endpoints
	"runtime"
	"time"
)

func main() {
	// Start an HTTP server for pprof profiling on port 6060
	go func() {
		fmt.Println("Profiling server running on http://localhost:6060/debug/pprof/")
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
			fmt.Printf("Profiling server error: %v\n", err)
		}
	}()

	const elementCount = 10000

	// Run benchmarks sequentially with reusable variables
	var beforeAlloc, beforeHeapSys, afterAlloc, afterHeapSys uint64
	var beforeGCFraction, afterGCFraction float64
	var pushDuration, popDuration time.Duration

	// Benchmark array-based stack
	fmt.Println("\nBenchmarking Stacks...")
	arrayStack := stack.NewArray(elementCount)

	beforeAlloc, beforeHeapSys, beforeGCFraction = utils.PrintMemoryStats("Before ArrayStack")
	pushDuration, popDuration = bench.BenchmarkStack("ArrayStack", arrayStack, elementCount, elementCount/2)
	afterAlloc, afterHeapSys, afterGCFraction = utils.PrintMemoryStats("After ArrayStack Push/Pop")

	arrayBeforeMem := []float64{float64(beforeAlloc), float64(beforeHeapSys), beforeGCFraction}
	arrayAfterMem := []float64{float64(afterAlloc), float64(afterHeapSys), afterGCFraction}
	arrayPerf := []time.Duration{pushDuration, popDuration}

	// Clear array stack to help GC
	arrayStack = nil
	runtime.GC()
	time.Sleep(time.Millisecond * 50)

	// Benchmark linked list-based stack
	linkedListStack := stack.NewLinkedList()

	beforeAlloc, beforeHeapSys, beforeGCFraction = utils.PrintMemoryStats("Before LinkedListStack")
	pushDuration, popDuration = bench.BenchmarkStack("LinkedListStack", linkedListStack, elementCount, elementCount/2)
	afterAlloc, afterHeapSys, afterGCFraction = utils.PrintMemoryStats("After LinkedListStack Push/Pop")

	linkedBeforeMem := []float64{float64(beforeAlloc), float64(beforeHeapSys), beforeGCFraction}
	linkedAfterMem := []float64{float64(afterAlloc), float64(afterHeapSys), afterGCFraction}
	linkedPerf := []time.Duration{pushDuration, popDuration}

	// Clear linked list stack to help GC
	linkedListStack = nil

	// Trigger final GC
	runtime.GC()
	time.Sleep(time.Millisecond * 50)
	afterGCAlloc, afterGCHeapSys, afterGCGCFraction := utils.PrintMemoryStats("After GC")

	// Print the results table
	utils.PrintTable(
		arrayBeforeMem[0], arrayBeforeMem[1], arrayBeforeMem[2],
		arrayAfterMem[0], arrayAfterMem[1], arrayAfterMem[2],
		linkedBeforeMem[0], linkedBeforeMem[1], linkedBeforeMem[2],
		linkedAfterMem[0], linkedAfterMem[1], linkedAfterMem[2],
		float64(afterGCAlloc), float64(afterGCHeapSys), afterGCGCFraction,
		arrayPerf[0], arrayPerf[1], linkedPerf[0], linkedPerf[1],
	)

	// Keep the main goroutine running for profiling
	fmt.Println("Press Ctrl+C to stop the profiling server...")
	select {} // Block indefinitely to allow profiling
}
