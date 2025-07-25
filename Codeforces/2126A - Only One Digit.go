package main
import "fmt"

func main() {
  x:=0
  fmt.Scanln(&x)
  for i:= 0; i<x; i++ {
    var y int
    fmt.Scanln(&y)
    res := y%10
    y /= 10
    for y > 0 {
       if res >( y%10){
           res = y%10
       }
      y /= 10
    }
    fmt.Println(res)
  }
}