package main

import "github.com/jbowkett/go-katas/fizzbuzz"

func main(){
  printer := fizzbuzz.ConsolePrinter{}
  fizzer := fizzbuzz.NewFizzBuzz(0, 100)
  fizzer.PrintOn(&printer)
}
