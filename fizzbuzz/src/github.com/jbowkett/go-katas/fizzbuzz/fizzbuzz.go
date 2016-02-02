package fizzbuzz

import (
	"strconv"
	"fmt"
)


type fizzBuzz struct {
	start, end int
}

func (this *fizzBuzz) printOn(printer Printer) {
	for i := this.start; i <= this.end; i++ {
		msg := ""
		divisibleByThree := i % 3 == 0
		divisibleByFive := i % 5 == 0
		if divisibleByThree {
			msg += "FIZZ"
		}
		if divisibleByFive {
			msg += "BUZZ"
		}
		if(!divisibleByThree && !divisibleByFive) {
 			msg = strconv.Itoa(i)
		}
		printer.println(i, msg)
	}
}


type Printer interface {
	println(lineNumber int, line string)
}


type ConsolePrinter struct {
	Printer
}

func (p *ConsolePrinter) println(lineNumber int, line string) {
	fmt.Println(line)
}


