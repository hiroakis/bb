package bb

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Run(args []string) int {
	var s string
	if len(args) < 1 {
		in := bufio.NewScanner(os.Stdin)
		if in.Scan() {
			s = in.Text()
		}
	} else {
		s = args[0]
	}
	s = strings.TrimSpace(s)

	if len(s)%2 != 0 {
		return 1
	}

	var (
		str    string
		bSlice string
	)
	for n, c := range s {
		if !((48 <= c && c <= 57) ||
			(65 <= c && c <= 70) || (97 <= c && c <= 102)) {
			return 1
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
	return 0
}
