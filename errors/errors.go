package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%s int %d", e.prob, e.arg)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{10, 42, 23} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 81, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}
	m, e := f2(42)
	fmt.Println("\n", m)
	fmt.Println(e)
	fmt.Println(e.Error)
	fmt.Println(e.Error())
	fmt.Println(e.(*argError), "\n")

	_, e = f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
		fmt.Println(ae.Error)
		fmt.Println(ae.Error())
	}
}
