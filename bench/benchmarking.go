package bench

import (
	"github.com/mdshahjahanmiah/go-profiling/stack"
	"time"
)

func BenchmarkStack(stackName string, stackInterface interface{}, pushCount, popCount int) (pushTime, popTime time.Duration) {
	startTime := time.Now()
	// Push operations
	for i := 1; i <= pushCount; i++ {
		switch s := stackInterface.(type) {
		case *stack.Array:
			s.Push(i)
		case *stack.LinkedList:
			s.Push(i)
		}
	}
	pushTime = time.Since(startTime)

	startTime = time.Now()
	// Pop operations
	for i := 0; i < popCount; i++ {
		switch s := stackInterface.(type) {
		case *stack.Array:
			if _, ok := s.Pop(); !ok {
				break
			}
		case *stack.LinkedList:
			if _, ok := s.Pop(); !ok {
				break
			}
		}
	}
	popTime = time.Since(startTime)
	return pushTime, popTime
}
