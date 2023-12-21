package errors

import "fmt"

type MissingParamsError struct {
	Code   int
	Params []string
}

func (err MissingParamsError) Error() {
	fmt.Printf("missing param(s) : %v", err.Params)
}
