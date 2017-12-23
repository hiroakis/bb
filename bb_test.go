package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestRun(t *testing.T) {
	input := "16bf0f1e88de"
	expect := "[]byte{0x16, 0xbf, 0xf, 0x1e, 0x88, 0xde}"

	buf := &bytes.Buffer{}
	cl := &cli{outStream: buf, errStream: ioutil.Discard}
	err := cl.run([]string{input})
	if err != nil {
		t.Errorf("err should be nil but: %s", err)
	}
	out := buf.String()
	if out != expect {
		t.Errorf("something went wrong.\n   out: %s\nexpect: %s", out, expect)
	}
}
