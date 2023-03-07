package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "print in same line")
var sep = flag.String("s", " ", "separator")

func main() {
	// var s, sep string
	// for i := 1; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }
	// fmt.Println("name: " + os.Args[0])
	// fmt.Println(s)

	// for index, arg := range os.Args[1:] {
	// 	fmt.Println(index, " "+arg)
	// }

	// fmt.Println(strings.Join(os.Args[1:], " "))

	// echo with command line flag
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	fmt.Println("this is next line")
}
