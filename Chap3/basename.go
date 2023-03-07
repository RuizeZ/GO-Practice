// basename removes directory components and a .suffix
// e.g. a => a; a.go => a; a/b => b; a/b.c.go => b.c
package main

import "fmt"

func main() {
	fmt.Println("a => " + basename("a"))
	fmt.Println("a.go => " + basename("a.go"))
	fmt.Println("a/b => " + basename("a/b "))
	fmt.Println("a/b.c.go => " + basename("a/b.c.go"))
}

func basename(s string) string {
	// discard last '/' and everything before
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// preserve everything before last '.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
