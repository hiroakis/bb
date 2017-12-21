package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	if len(os.Args) < 2 {
		in := bufio.NewScanner(os.Stdin)
		if in.Scan() {
			s = in.Text()
		}
	} else {
		s = os.Args[1]
	}
	s = strings.TrimSpace(s)

	if len(s)%2 != 0 {
		os.Exit(1)
	}

	var (
		str    string
		bSlice string
	)
	for n, c := range s {
		if !((48 <= c && c <= 57) ||
			(65 <= c && c <= 70) || (97 <= c && c <= 102)) {
			os.Exit(1)
		}
		str = str + string(c)
		if (n+1)%2 == 0 {
			if n == len(s)-1 {
				bSlice = bSlice + fmt.Sprintf("0x%s", str)
			} else {
				bSlice = bSlice + fmt.Sprintf("0x%s, ", str)
				str = ""
			}
		}
	}
	fmt.Printf("[]byte{%s}", bSlice)
}
