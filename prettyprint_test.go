package prettyprint

import (
	"io/ioutil"
	"os"
	"testing"
)

type TestStruct1 struct {
	integer int
	str     string
}

func TestPrettyPrint(t *testing.T) {
	p := NewPrinter(PrinterOptions{countFunction: false})

	p = testPrettyPrintInt(t, p)
	p = testPrettyPrintPointerStruct(t, p)
	p = testPrettyPrintStruct(t, p)
}

func testPrettyPrintInt(t *testing.T, p *Printer) *Printer {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	p.Print(12)

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = stdOut

	expected := "0: Type: int, Value: 12"
	if expected != string(out) {
		t.Errorf("Expected %v, actual %v", expected, string(out))
	}

	return p
}

func testPrettyPrintPointerStruct(t *testing.T, p *Printer) *Printer {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	p.Print(&TestStruct1{integer: 1, str: "hey"})

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = stdOut

	expected := "1: Type: *prettyprint.TestStruct1, Value: &{1 hey}"
	if expected != string(out) {
		t.Errorf("Expected %v, actual %v", expected, string(out))
	}

	return p
}

func testPrettyPrintStruct(t *testing.T, p *Printer) *Printer {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	p.Print(TestStruct1{integer: 1, str: "hey"})

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = stdOut

	expected := "2: Type: prettyprint.TestStruct1, Value: {1 hey}"
	if expected != string(out) {
		t.Errorf("Expected %v, actual %v", expected, string(out))
	}

	return p
}

func TestFuncCounter(t *testing.T) {
	p := NewPrinter(PrinterOptions{countFunction: true})
	testFunction := func() {
		p.Print("var")
	}

	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	testFunction()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = stdOut

	expected := "prettyprint.TestFuncCounter.func1 2: Type: prettyprint.TestStruct1, Value: {1 hey}"
	if expected != string(out) {
		t.Errorf("Expected %v, actual %v", expected, string(out))
	}
}
