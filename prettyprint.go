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

type Printer struct {
	counter int
}

func NewPrinter() *Printer {
	return &Printer{counter: 0}
}

func (p *Printer) Print(i interface{}) {
	defer p.count()

	_, _, line, _ := runtime.Caller(1)

	switch v := i.(type) {
	case int, int8, int16, int64, uint, uint16, uint32, uintptr, string, bool, byte, rune, float32, float64:
		fmt.Printf(printed, p.counter, v, v, line)
	default:
		fmt.Printf(printedDynamic, p.counter, reflect.TypeOf(i).String(), reflect.ValueOf(i), line)
	}

	return
}

func (p *Printer) count() {
	p.counter++
}
