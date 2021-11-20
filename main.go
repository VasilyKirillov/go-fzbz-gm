package main

import "fmt"

func main() {
	out := make(chan string)

	var cfg FizzBuzzConfiger = &FizzBuzzConfig{}
	cfg.promptRange()

	var counter FizzBuzzCounter = &FizzBuzzCount{cfg: &cfg}
	go counter.count(out)

	for n := range out {
		fmt.Print(n)
	}
}
