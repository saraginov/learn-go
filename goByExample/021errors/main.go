package main

import (
	"errors"
	"fmt"
)

// as shown in f1, errors are the last return value and have a type error
func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New() constructs basic error value
		return -1, errors.New("can't work with 42")
	}
	// nil value in the error position indicates that there was no error.
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

// possible to use custom types as errors by implementing Error() method
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// use &argError syntax to build a new struct, supplying values for the
		// two fields arg and prob
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	// the 2 loops below test out each of the error returning functions.
	// inline error check on the if line
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// to programmatically use the data in a custom error,
	// get the error as an instance of the custom error via type assertion.
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
