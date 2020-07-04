package main

import (
	"errors"
	"fmt"
)

// CustomError fsa
type CustomError struct {
	error
	msg string
}

func (c CustomError) Error() string {
	return fmt.Sprintf("Error: %s, msg: %s", c.error, c.msg)
}

func main() {
	err := CustomError{
		error: errors.New("I am a error"),
		msg:   "My custom error",
	}

	foo(err)
}

func foo(err error) {
	fmt.Println(err)
}
