// Generate the first 13 terms of the Fibonacci sequence:
//      0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, ...
//
// In mathematical terms, the sequence Fn of Fibonacci numbers is defined by the recurrence relation
//      F_n = F_{n-1} + F_{n-2}
// with seed values
//      F_1 = 1, F_2 = 1
// or (modern usage)
//      F_0 = 0, F_1 = 1
//
// See https://en.wikipedia.org/wiki/Fibonacci_number
package main

import (
	"fmt"
)

func transient(n uint) {
	var Fn, Fn1, Fn2 uint = 0, 0, 1

	fmt.Printf("Transient:\t%v %v ", Fn1, Fn2)
	for i := uint(2); i < n; i++ {
		Fn = Fn1 + Fn2
		Fn1 = Fn2
		Fn2 = Fn
		fmt.Printf("%v ", Fn)
	}
	fmt.Print("\n")
}

func slice(n uint) {
	Fn := make([]uint, n)

	for i := range Fn {
		switch {
		case i < 2:
			Fn[i] = uint(i)

		default:
			Fn[i] = Fn[i-1] + Fn[i-2]
		}
	}
	fmt.Printf("Slice:\t\t%v\n", Fn)
}

func closure() func() uint {
	var Fn, Fn1, Fn2 uint = 0, 0, 1

	return func() uint {
		Fn = Fn1 + Fn2
		Fn1 = Fn2
		Fn2 = Fn
		return Fn
	}
}

func main() {
	var n uint = 13

	transient(n)

	slice(n)

	fmt.Printf("Closure:\t")
	Fn := closure()
	for i := uint(0); i < n; i++ {
		fmt.Printf("%d ", Fn())
	}
	fmt.Print("\n")
}
