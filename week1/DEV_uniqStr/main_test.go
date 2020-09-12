package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var testOK = `1
2
3
3
3
3
4
5`
var testOKResult = `1
2
3
4
5
`

func TestOK(t *testing.T)  {
	in := bufio.NewReader(strings.NewReader(testOK))
	out := new(bytes.Buffer)
	err := uniq(in, out)
	if err != nil {
		t.Errorf("test TestOK Failed")
	}
	if out.String() != testOKResult {
		t.Errorf("Out != testOKResult \n %v %v", out.String(), testOKResult)
	}
}

var testFail = `1
2
3
1`

func TestForError(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testFail))
	out := new(bytes.Buffer)
	err := uniq(in, out)
	if err == nil {
		t.Errorf("test TestForError Failed")
	}
}
