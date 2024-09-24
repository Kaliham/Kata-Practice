package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	var input string
	var err error
	var arr []string

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("How Many ")
	fmt.Println("Enter String to split")
	input, err = reader.ReadString(10)
	if err != nil {
		fmt.Println("Error Occured while reading input")
	}
	//removing the return carrage and line feed
	input = strings.TrimRightFunc(input, func(r rune) bool { return r <= 13 })
	arr = split(input)

	fmt.Print(arr)
}

func split(str string) (arr []string) {
	a := make([]string, len(str))
	for i, r := range str {
		a[i] = string(r)
	}

	l := len(a)
	s := l/2 + l%2
	t := []string{""}
	arr = slices.Repeat(t, s)
	var i int
	var c string
	fmt.Println(a)
	for i, c = range a {
		arr[i/2] = arr[i/2] + string(c)
	}
	if l%2 == 1 {
		arr[i/2] += "_"
	}

	for _, e := range arr {
		fmt.Printf("%v,", e)
	}
	return
}
