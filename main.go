package main

import (
	"bufio"
	"fmt"
	"os"
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

	if len(s)%2 != 0 {
		fmt.Println(s)
		return
	}

	var (
		bb     string
		bSlice string
	)
	for n, c := range s {
		if !(48 <= c || c <= 57 ||
			65 <= c || c <= 70 || 97 <= c || c <= 102) {
			fmt.Println(s)
			return
		}
		bb = bb + string(c)
		if (n+1)%2 == 0 {
			bSlice = bSlice + fmt.Sprintf("0x%s, ", bb)
			bb = ""
		}
	}
	fmt.Printf("[]byte{%s}", bSlice)
}
