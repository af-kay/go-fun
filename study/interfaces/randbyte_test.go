package interfaces

import (
	"fmt"
	"testing"
	"time"
)

func TestRandomBytes(t *testing.T) {
	// создаём генератор случайных чисел
	generator := New(time.Now().UnixNano()) // в качестве затравки передаём ему текущее время, и при каждом запуске оно будет разным.

	buf := make([]byte, 8)

	for i := 0; i < 5; i++ {
		n, _ := generator.Read(buf) // единственный доступный метод, но он нам и нужен.
		fmt.Printf("Generate bytes: %v size(%d)\n", buf, n)
	}
}
