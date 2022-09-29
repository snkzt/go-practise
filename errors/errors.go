package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		// Errors are the last return value and have type error
		// errors.New constructs a basic error value with the given error message.
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

// We can create custom types as errors by implementing
// the Error() method on them.
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		// Set r and e with returned value with f2 and check the if condition.
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	// e.(*argError) is type assertion to get error as an instance of the custom error type
	// Type assertion provide access to the exact type of variable of an interface.
	// If already the data type is present in the interface, then it will retrieve the actual data type value
	// held by the interface, arg and prob in this case.
	// t := value.(typeName) // value is a variable whose type must be an interface,
	// typeName is the concrete type we want to check and underlying typeName value is assigned to variable t.
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
