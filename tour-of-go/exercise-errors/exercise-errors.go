package main

import "fmt"

// ErrNegativeSqrt is a place holder for -ive value errors
type ErrNegativeSqrt float64

/*
	We do not need an interface type because, Error() and String()
	are methods which belong to 'type ErrNegativeSqrt float64'

	Had we used functions, which do not have a receiver argument,
	an interface would've been required to let 'GO' know
	how to define a set of method signatures for ErrNegativeSqrt

		NOTE:
			The functions in the interface would've had to use
			ErrNegativeSqrt as a parameter in the function declaration
*/

// Error method belongs to type ErrNegativeSqrt float64
func (err ErrNegativeSqrt) Error() string {
	errorString := err.String()
	return errorString
}

// String is used to format the -ive val root to "cannot Sqrt negative number: -2"
func (err ErrNegativeSqrt) String() string {
	/*
		Note: A call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop.
		You can avoid this by converting e first: fmt.Sprint(float64(e)). Why?

			The reason why an infinite loop is created
			when fmt.Sprint(e) is called inside Error() is because
			i) both methods belong to the same instance of 'ErrNegativeSqrt'
			and
			ii) 'The error type is a built-in interface similar to fmt.Stringer'

			Thus each time Error() is called,
			and it, in-turn calls fmt.Sprint(e),

			fmt.Sprint(being an built-in interface,
			re-initializes the value of 'ErrNegativeSqrt in the memory address

			which in turn calls re-initializes Error()
			as it, too is a built-in interface

			which then calls fmt.Sprint(e) again,
			and this process continues, creating an infinite loop

			On the other hand when we cast err as a new float64,
			the 'ErrNegativeSqrt' does not re-initialize
			since a new memory address is allocated to the local variable
	*/

	return fmt.Sprint("cannot Sqrt negative number: ", float64(err))
}

// OurNewSqrt is identical to OurSqrt in loops-and-functions exercise with error handling added
func OurNewSqrt(x float64) (float64, error) {
	if x < 0 {
		errorToReturn := ErrNegativeSqrt(x)
		return x, errorToReturn
	}

	z := float64(1)
	zPrev := float64(0)

	for z-zPrev > 0.001 {
		zPrev = z
		z -= (z*z - x) / (2 * z)
	}

	return z, nil
}

func main() {
	fmt.Println(OurNewSqrt(2))
	fmt.Println(OurNewSqrt(-2))
}
