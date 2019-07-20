package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//fmt.Println("Hello Hao")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(">")

	//var total int

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		var inputText string
		inputText = scanner.Text()
		s := strings.Split(inputText, " ")
		var total int
		for i, num := range s {
			if (i % 2) == 0 {

			}
		}
	}
	//if err := scanner.Err(); err != nil {
	//	fmt.Fprintln(os.Stderr, "reading standard input:", err)
	//}
	fmt.Println(">")
}

func isOperater(oper string, num1, num2 float32) int {
	switch oper {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
		return true
	}
	return false
}
