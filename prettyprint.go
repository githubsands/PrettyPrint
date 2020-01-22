package prettyprint

import (
	"fmt"
	"reflect"
)

const (
	printed        = "%v: Type: %T, Value: %v\n"
	printedDynamic = "%v: Type: %v, Value: %v\n"
)

type Printer struct {
	counter int
}

func NewPrinter() *Printer {
	return &Printer{counter: 0}
}

func (p *Printer) Print(i interface{}) {
	defer p.count()

	switch v := i.(type) {
	case int, int8, int16, int64, uint, uint16, uint32, uintptr, string, bool, byte, rune, float32, float64:
		fmt.Printf(printed, p.counter, v, v)
	default:
		fmt.Printf(printedDynamic, p.counter, reflect.TypeOf(i).String(), reflect.ValueOf(i))
	}

	return
}

func (p *Printer) count() {
	p.counter++
}
