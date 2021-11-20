package main

import "fmt"

type FizzBuzzCounter interface {
	count(out chan string)
}

type FizzBuzzCount struct {
	cfg *FizzBuzzConfiger
}

func (counter *FizzBuzzCount) count(out chan string) {
	defer close(out)

	cfg := counter.cfg

	l, u := (*cfg).getLower(), (*cfg).getUpper()
	for i := l; i <= u; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			out <- "fizzBuzz"
		case i%3 == 0:
			out <- "fizz"
		case i%5 == 0:
			out <- "buzz"
		default:
			out <- fmt.Sprint(i)
		}
		if i < u {
			out <- ", "
		}
	}

}
