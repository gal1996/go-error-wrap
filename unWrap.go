package main

import(
	"errors"
	"fmt"
)

func main() {
	wrappedErr := MakeWrappedError()
	fmt.Printf("type:%T", wrappedErr)

	unWrappedErr := errors.Unwrap(wrappedErr)
	fmt.Printf("type:%T", unWrappedErr)

	nilErr := errors.Unwrap(unWrappedErr)
	fmt.Printf("type:%T", nilErr)
}
