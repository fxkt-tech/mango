package math

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	x := Max(1.2, 2.3)
	fmt.Println(x)
}

func TestMaxMin(t *testing.T) {
	max, min := MaxMin(3, 2, 1)
	fmt.Println(max, min)
}

func TestClip(t *testing.T) {
	x := Clip(4, 1, 10)
	fmt.Println(x)
}
