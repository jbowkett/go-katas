package fizzbuzz

import (
	"testing"
)


type TestPrinter struct {
	Printer
	lineMap map[int]string
}


func NewTestPrinter() *TestPrinter {
	return &TestPrinter { lineMap: make(map[int]string) }
}


func (p *TestPrinter) println(lineNumber int, line string) {
	p.lineMap[lineNumber] = line
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

