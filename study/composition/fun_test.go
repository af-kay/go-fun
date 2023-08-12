package composition

import (
	"fmt"
	"testing"
)

func TestComposition(t *testing.T) {
	s := NewStudent("John Doe", 1980, "701")
	s.Print()
    s.Debug()
	// вызовется метод String() для Student
	fmt.Println(s)
	fmt.Println(s.Name, s.Year, s.Group)
}
