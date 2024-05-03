package generics_test

import (
	"fmt"
)

type GenIntf[T any] interface {
	comparable
	GetType() T
}

type impl struct{}

func (impl) GetType() int {
	return 42
}

func GetT[U GenIntf[T], T any](u U) string {
	return fmt.Sprintf("%T", u.GetType())
}

func main() {
	var x impl

	fmt.Println(GetT(x))
}
