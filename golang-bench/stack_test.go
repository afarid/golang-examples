package stack

import (
	"fmt"
	"testing"
)

func BenchmarkMutableStackPush(b *testing.B) {
	newStack := NewMutableStack()
	numGoroutines := []int{1, 5, 100, 1000, 5000}
	for _, curNum := range numGoroutines {
		b.Run(fmt.Sprintf("push with %d g", curNum), func(b *testing.B) {
			b.SetParallelism(curNum)
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					newStack.Push(1)
					newStack.Top()
				}
			})
		})
	}
}