package fizzbuzz

import (
	"strconv"
	"fmt"
)


type FizzBuzz struct {
	start, end int
}


func (this *FizzBuzz) PrintOn(printer Printer) {
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
		printer.Println(i, msg)
	}
}


type Printer interface {
	Println(lineNumber int, line string)
}


type ConsolePrinter struct {
	Printer
}

func (p *ConsolePrinter) Println(lineNumber int, line string) {
	fmt.Println(line)
}


