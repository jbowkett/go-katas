package fizzbuzz

import (
	"testing"
	"fmt"
	"strconv"
)

type fizzBuzz struct {
	start, end int
}

func (p *fizzBuzz) printOn(printer Printer) {
	for i := p.start; i < p.end; i++ {
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


type TestPrinter struct {
	Printer
	lineMap map[int]string
}

func NewTestPrinter() *TestPrinter {
	return &TestPrinter{ lineMap: make(map[int]string) }
}


func (p *TestPrinter) println(lineNumber int, line string) {
	p.lineMap[lineNumber] = line
}


type ConsolePrinter struct {
	Printer
}

func (p *ConsolePrinter) print(line string) {
	fmt.Println(line)
}



func doTest(toTest *fizzBuzz, testPrinter *TestPrinter, testing *testing.T, expected string, lineNumber int) {
	toTest.printOn(testPrinter)
	actual := testPrinter.lineMap[lineNumber]
	if (actual != expected) {
		testing.Error("Expected:", expected, "Got:", actual)
	}
}


func TestFizzIsPrintedOnMultiplesOfThree(testing *testing.T) {
	testPrinter := NewTestPrinter()
	toTest := fizzBuzz{1, 100}
	doTest(&toTest, testPrinter, testing, "FIZZ", 3)
}

func TestBuzzIsPrintedOnMultiplesOfFive(testing *testing.T) {
	testPrinter := NewTestPrinter()
	toTest := fizzBuzz{1, 100}
	doTest(&toTest, testPrinter, testing, "BUZZ", 5)
}


func TestFizzBuzzIsPrintedOnMultiplesOfFiveAndThree(testing *testing.T) {
	testPrinter := NewTestPrinter()
	toTest := fizzBuzz{1, 100}
	doTest(&toTest, testPrinter, testing, "FIZZBUZZ", 15)
}

func TestJustTheNumberIsPrintedOnNonMultiplesOfFiveNorThree(testing *testing.T) {
	testPrinter := NewTestPrinter()
	toTest := fizzBuzz{1, 100}
	expected := "16"
	doTest(&toTest, testPrinter, testing, expected, 16)
}


func main(){
	printer := ConsolePrinter{}
	f := fizzBuzz{1, 100}
	f.printOn(&printer)
}


