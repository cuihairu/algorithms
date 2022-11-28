package array

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	l := make([]int, 0, 3)
	fmt.Println(cap(l))
}
