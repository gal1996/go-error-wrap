package main

import(
	"errors"
	"fmt"
)

type User struct {
	name string
	email string
}
type invalidStructError struct {
	original interface{} //元の構造体
	err error
	code int
}


func (e *invalidStructError)Error() string {
	return fmt.Sprintf("original: %v, %s", e.original, e.err.Error())
}

func marshal(v interface{}) error {
	err := &invalidStructError{original:v, err: errors.New("invalid Struct"), code: 0}
	if err != nil{
		return fmt.Errorf("marshal err: %w", err)
	}
	return nil
}


func main() {
	val := User{name: "hoge", email: "1"}
	err := marshal(val)
	var target *invalidStructError
	if errors.As(err, &target) {
		fmt.Printf("code: %d\n", target.code)
		fmt.Printf("%s\n", target.Error())
	} else {
		fmt.Printf("defaultError")
	}
}

