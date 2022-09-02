package defer_keyword

import (
	"fmt"
	"testing"
)

func Test3(t *testing.T) {
	defer func() {
		if ok := recover(); ok != nil {
			fmt.Println("recover")
		}
	}()

	panic("error")
}
