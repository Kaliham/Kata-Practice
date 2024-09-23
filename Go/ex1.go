package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input string
	var err error
	var arr []string
	fmt.Println("Enter String to split")
	reader := bufio.NewReader(os.Stdin)
	input, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error Occured while reading input")
	}
	arr = split(input)

	fmt.Println(arr)
}

func split(str string) (arr []string) {
	a := []rune(str)
	l := len(a)
	s := l/2 + l%2
	arr = make([]string, s)
	var i int
	var c rune
	for i, c = range a {
		arr[i/2] += string(c)
	}
	if l%2 == 1 {
		arr[i/2] += "_"
	}
	return
}
