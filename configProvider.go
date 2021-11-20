package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	LOW_PROMPT  = "Please input where to start counting the fizzBuzz numbers (upper boundry):"
	UP_PROMPT   = "Please input where to stop counting the fizzBuzz numbers (lower boundry):"
	LOW_ERR_MSG = "Wrong lower boundry"
	UP_ERR_MSG  = "Wrong upper boundry"
)

type FizzBuzzConfiger interface {
	promptRange()
	getUpper() int
	getLower() int
}

type FizzBuzzConfig struct {
	l, u int
}

func (c *FizzBuzzConfig) getUpper() int {
	return c.u
}
func (c *FizzBuzzConfig) getLower() int {
	return c.l
}

func (c *FizzBuzzConfig) promptRange() {
	reader := bufio.NewReader(os.Stdin)

	for {
		l_text := readNumberStr(reader, LOW_PROMPT, LOW_ERR_MSG)
		c.l = parseInt(l_text, LOW_ERR_MSG)

		u_text := readNumberStr(reader, UP_PROMPT, UP_ERR_MSG)
		c.u = parseInt(u_text, UP_ERR_MSG)

		if r := recover(); r == nil {
			break
		}
	}
}

func readNumberStr(reader *bufio.Reader, prompt string, errMsg string) (text string) {
	fmt.Println(prompt)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(errMsg + text)
	}
	return
}

func parseInt(text string, errMsg string) int {
	var n int
	_, err := fmt.Sscan(text, &n)
	if err != nil {
		panic(errMsg + text)
	}
	return n
}
