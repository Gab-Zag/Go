package main

import (
	"fmt"
  "math/rand"
)

func factorial(x int) int{
  resultado := 3
  resultado = 1
  for i := 1; i <= x ; i++{
    resultado *= i
  }
  return resultado
}

func main() {
  x := 3
  x = rand.Intn(10)
  resultado := 33
  resultado = factorial(x)
  fmt.Println(resultado)
}
