package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("name: " + os.Args[0])
	fmt.Println(s)

	for index, arg := range os.Args[1:] {
		fmt.Println(index, " "+arg)
	}

	fmt.Println(strings.Join(os.Args[1:], " "))

}
