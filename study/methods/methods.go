package methods

import "fmt"

type Age uint

func (a Age) String() string {
    return fmt.Sprintf("Age: %d", a)
}

