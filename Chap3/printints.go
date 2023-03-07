// fmt.Sprintf(values) but adds commas.
package main

import (
	"bytes"
	"fmt"
)

func intsTiString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(intsTiString([]int{1, 2, 3}))
}
