package main

import "fmt"

func main() {
	var s = []string{"a", "a", "b", "c", "c"}
	fmt.Printf("%q\n", s)
	fmt.Println(removeDuplicat(s))
	s = []string{"a", "a", "a", "a", "a"}
	fmt.Println(removeDuplicat(s))
	s = []string{"a", "c", "b", "c", "c"}
	fmt.Println(removeDuplicat(s))
	s = []string{"a", "b", "c", "d", "e"}
	fmt.Println(removeDuplicat(s))
	s = []string{"a", "b", "c", "b", "a"}
	fmt.Println(removeDuplicat(s))
}

func removeDuplicat(s []string) []string {
	i := 0
	for _, str := range s {
		if i == 0 {
			i++
			continue
		}
		if s[i-1] != str {
			s[i] = str
			i++
		}
	}
	return s[:i]
}
