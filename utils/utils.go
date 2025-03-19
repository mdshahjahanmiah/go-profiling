package utils

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

// PrintTable - Print the comparison table: AI generated code
func PrintTable(arrayBeforeAlloc, arrayBeforeHeapSys float64, arrayBeforeGCFraction float64,
	arrayAfterAlloc, arrayAfterHeapSys float64, arrayAfterGCFraction float64,
	linkedBeforeAlloc, linkedBeforeHeapSys float64, linkedBeforeGCFraction float64,
	linkedAfterAlloc, linkedAfterHeapSys float64, linkedAfterGCFraction float64,
	afterGCAlloc, afterGCHheapSys float64, afterGCGCFraction float64,
	arrayPush, arrayPop, linkedPush, linkedPop time.Duration) {

	fmt.Println("\nStack Comparison Table")
	fmt.Println(strings.Repeat("-", 105))
	fmt.Printf("| %-20s | %-20s | %-20s | %-20s | %-20s |\n",
		"Metric", "ArrayStack (Before)", "ArrayStack (After)", "LinkedListStack (Before)", "LinkedListStack (After)")
	fmt.Println(strings.Repeat("-", 105))
	fmt.Printf("| %-20s | %20v KB | %20v KB | %20v KB | %20v KB |\n",
		"Alloc", arrayBeforeAlloc, arrayAfterAlloc, linkedBeforeAlloc, linkedAfterAlloc)
	fmt.Printf("| %-20s | %20v KB | %20v KB | %20v KB | %20v KB |\n",
		"HeapSys", arrayBeforeHeapSys, arrayAfterHeapSys, linkedBeforeHeapSys, linkedAfterHeapSys)
	fmt.Printf("| %-20s | %20.2f %% | %20.2f %% | %20.2f %% | %20.2f %% |\n",
		"GC CPU Fraction", arrayBeforeGCFraction, arrayAfterGCFraction, linkedBeforeGCFraction, linkedAfterGCFraction)
	fmt.Printf("| %-20s | %20s | %20v | %20s | %20v |\n",
		"Push Time (10,000)", "-", arrayPush, "-", linkedPush)
	fmt.Printf("| %-20s | %20s | %20v | %20s | %20v |\n",
		"Pop Time (5,000)", "-", arrayPop, "-", linkedPop)
	fmt.Printf("| %-20s | %20s | %20s | %20s | %20v KB |\n",
		"After GC Alloc", "-", "-", "-", afterGCAlloc)
	fmt.Printf("| %-20s | %20s | %20s | %20s | %20v KB |\n",
		"After GC HeapSys", "-", "-", "-", afterGCHheapSys)
	fmt.Printf("| %-20s | %20s | %20s | %20s | %20.2f %% |\n",
		"After GC CPU Fraction", "-", "-", "-", afterGCGCFraction)
	fmt.Println(strings.Repeat("-", 105))
}

func PrintMemoryStats(prefix string) (alloc, heapSys uint64, gcFraction float64) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	alloc = m.Alloc / 1024             // Convert to KB
	heapSys = m.HeapSys / 1024         // Convert to KB
	gcFraction = m.GCCPUFraction * 100 // Convert to percentage
	return alloc, heapSys, gcFraction
}
