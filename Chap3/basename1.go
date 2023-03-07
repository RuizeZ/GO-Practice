// basename removes directory components and a .suffix
// e.g. a => a; a.go => a; a/b => b; a/b.c.go => b.c
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("a => " + basename("a"))
	fmt.Println("a.go => " + basename("a.go"))
	fmt.Println("a/b => " + basename("a/b "))
	fmt.Println("a/b.c.go => " + basename("a/b.c.go"))
}

// use strings package
func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
