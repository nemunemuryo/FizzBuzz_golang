package main

import(
  "testing"
)

func TestFizzBuzz(t *testing.T) {
  actual := fizzBuzz(15)
  if actual != "FizzBuzz" {
    t.Fatal("failed test")
  }
}

func TestFizz(t *testing.T) {
  actual := fizzBuzz(3)
  if actual != "Fizz" {
    t.Fatal("failed test")
  }
}

func TestBuzz(t *testing.T) {
  actual := fizzBuzz(5)
  if actual != "Buzz" {
    t.Fatal("failed test")
  }
}
