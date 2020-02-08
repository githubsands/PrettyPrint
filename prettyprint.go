package prettyprint

import (
	"fmt"
	"reflect"
	"runtime"
)

const (
	printed        = "%v: Type: %T, Value: %v, Line: %v\n"
	printedDynamic = "%v: Type: %v, Value: %v, Line: %v\n"
)

type PrinterOptions struct {
	CountFunction bool
}

type Printer struct {
	printCounter    int
	funcNameCounter int
}

func NewPrinter(o PrinterOptions) *Printer {
	p := new(Printer)

	if o.CountFunction == false {
		p.countFunction()
	}

	return p
}

func (p *Printer) PrintCheck() {
	defer p.countPrints()

	var (
		pc, _, line, _ = runtime.Caller(1)
	)

	// only display the function name once
	if p.funcNameCounter == 0 {
		fs := runtime.FuncForPC(pc)
		fmt.Printf("----------%v:----------\n", fs.Name())
		p.countFunction()
	}

	fmt.Printf("%v: Made it to %v", p.printCounter, line)

	return
}

func (p *Printer) PrintVar(i interface{}) {
	defer p.countPrints()

	var (
		pc, _, line, _ = runtime.Caller(1)
	)

	// only display the function name once
	if p.funcNameCounter == 0 {
		fs := runtime.FuncForPC(pc)
		fmt.Printf("----------%v:----------\n", fs.Name())
		p.countFunction()
	}

	switch v := i.(type) {
	case int, int8, int16, int64, uint, uint16, uint32, uintptr, string, bool, byte, rune, float32, float64:
		fmt.Printf(printed, p.printCounter, v, v, line)
	default:
		fmt.Printf(printedDynamic, p.printCounter, reflect.TypeOf(i).String(), reflect.ValueOf(i), line)
	}

	return
}

// countPrints prints how many times print has been called
func (p *Printer) countPrints() {
	p.printCounter++
}

// countFuncs prints the function where this printer was called.
func (p *Printer) countFunction() {
	// funcNameCounter only needs to count once, this program does not need to know how many times the func is called.  Just if its called.
	if p.funcNameCounter > 0 {
		return
	}

	p.funcNameCounter++
}
