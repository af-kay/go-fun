package funclosures

import "fmt"

type MyStruct struct {
	A   int
	Log func(s string) string
}

func makeMyStruct(value int) (s *MyStruct) {
	s = &MyStruct{
		A: value,
		Log: func(message string) (m string) {
			m = fmt.Sprintf("Struct value is %d, message is %s", s.A, message)

			fmt.Println(m)
			return
		},
	}

	return
}
