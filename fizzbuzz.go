package main

import(
  "fmt"
)

func fizzBuzz(num int) string {
  switch {
  case num % 3 == 0 && num % 5 == 0:
    return "FizzBuzz"
  case num % 3 == 0:
    return "Fizz"
  case num % 5 == 0:
    return "Buzz"
  }
  return ""
}

func main() {
  n := 30
  for i:=1; i<=n; i++ {
    fmt.Println(i,fizzBuzz(i))
  }
}
