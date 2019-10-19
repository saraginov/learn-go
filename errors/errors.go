package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
	Errors

		Go programs express error state with error values

		The error type is a built-in interface similar to fmt.Stringer

			type error interface {
				Error() string
			}

			NOTE: As with fmt.Stringer, the fmt package looks for the error interface when printing values.

		IMPORTANT:
		Functions ofter return an error value,
		and calling code should handle errors by testing whether the error equals nil.

			i, err := strconv.Atoi("42")
			if err != nil {
				fmt.Printf("Couldn't convert number: %v\n", err)
				return
			}
			fmt.Println("Converted interger:", i)

		A nil error denotes success;
		a non-nil error denotes failure
*/

// MyError is a "fake" error construct
type MyError struct {
	When time.Time
	What string
}

// StringConversionIntError is also a "fake" error construct
//
type StringConversionIntError struct {
	ConvertedString int
	SomeError       error
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s ", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func convertStringToInt(s string) StringConversionIntError {
	intToReturn, error := strconv.Atoi(s)
	return StringConversionIntError{
		intToReturn,
		error,
	}
}

func main() {
	/*
		NOTE: The expression may be preceded by a simple statement, which executes before the expression is evaluated.

		if x := f(); x < y {...}
	*/
	if err := run(); err != nil {
		fmt.Println(err)
	}

	stringThatCanBeConvertedToInt := "69"
	if convertedValue := convertStringToInt(stringThatCanBeConvertedToInt); convertedValue.SomeError != nil {
		fmt.Println(convertedValue.SomeError)
	} else {
		fmt.Printf("string: %s converted to int is: %d\n", stringThatCanBeConvertedToInt, convertedValue.ConvertedString)
	}

	stringThatCannotBeConvertedToInt := "sixtyNine"
	if convertedValue := convertStringToInt(stringThatCannotBeConvertedToInt); convertedValue.SomeError != nil {
		fmt.Println(convertedValue.SomeError)
	} else {
		fmt.Printf("string: %s converted to int is: %d\n", stringThatCanBeConvertedToInt, convertedValue.ConvertedString)
	}

	canFloatStringBeConvertedToInt := "3.14"
	if convertedValue := convertStringToInt(canFloatStringBeConvertedToInt); convertedValue.SomeError != nil {
		fmt.Println(convertedValue.SomeError)
	} else {
		fmt.Printf("string: %s converted to int is: %d\n", stringThatCanBeConvertedToInt, convertedValue.ConvertedString)
	}

}
