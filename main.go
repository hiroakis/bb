package main

import (
	"bufio"
	"encoding/hex"
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

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
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
	log.SetOutput(cl.errStream)
	log.SetFlags(0)
	fs := flag.NewFlagSet("bb", flag.ContinueOnError)
	fs.SetOutput(cl.errStream)
	fs.Usage = func() {
		fmt.Fprintf(cl.errStream, `bb - command line tool to replace hex string with Go byte slice

Usage:
    %% bb 16bf0f1e88de
    []byte{0x16, 0xbf, 0x0f, 0x1e, 0x88, 0xde}

Options:
`)
		fs.PrintDefaults()
	}
	err := fs.Parse(args)
	if err != nil {
		return err
	}
	args = fs.Args()

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

	decoded, err := hex.DecodeString(s)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(cl.outStream, "%#v", decoded)
	return err
}
