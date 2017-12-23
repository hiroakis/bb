package bb

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type cli struct {
	outStream, errStream io.Writer
}

const (
	exitCodeOK = iota
	exitCodeErr
)

func Run(args []string) int {
	cl := &cli{outStream: os.Stdout, errStream: os.Stderr}
	err := cl.run(args)
	if err != nil {
		if err != flag.ErrHelp {
			log.Println(err)
		}
		return exitCodeErr
	}
	return exitCodeOK
}

func (cl *cli) run(args []string) error {
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
		return fmt.Errorf("malformed hex string: %q", s)
	}

	var (
		str    string
		bSlice string
	)
	for n, c := range s {
		if !((48 <= c && c <= 57) ||
			(65 <= c && c <= 70) || (97 <= c && c <= 102)) {
			return fmt.Errorf("malformed hex string contained invalid character: %c in %q", c, s)
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
	_, err := fmt.Fprintf(cl.outStream, "[]byte{%s}", bSlice)
	return err
}
