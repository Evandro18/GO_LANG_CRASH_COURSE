package main

import "fmt"


func main() {
  var age int32 = 37
  const isCool = true
  var size float32 = 1.3
  name, email := "Brad", "brad@gmail.com"

  fmt.Println(name, email, age, isCool, size)
}
