package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	// read from standard input
	// input := bufio.NewScanner(os.Stdin)
	// for input.Scan() {
	// 	counts[input.Text()]++
	// }
	// for line, count := range counts {
	// 	if count > 1 {
	// 		fmt.Printf("%d\t%s\n", count, line)
	// 	}
	// }

	// read from file
	// files := os.Args[1:]
	// if len(files) == 0 {
	// 	countLines(os.Stdin, counts)
	// } else {
	// 	for _, arg := range files {
	// 		f, err := os.Open(arg)
	// 		if err != nil {
	// 			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	// 			continue
	// 		}
	// 		countLines(f, counts)
	// 		f.Close()
	// 	}
	// 	for line, count := range counts {
	// 		fmt.Printf("%d\t%s\n", count, line)
	// 	}
	// }

	// read file using os.ioutil.ReadFile
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\r\n") {
			if line != "psu" {
			}
			counts[line]++
		}
	}
	for line, count := range counts {
		fmt.Printf("%d\t%s\n", count, line)
	}

}

func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
	}
}
