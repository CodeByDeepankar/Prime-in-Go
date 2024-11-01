package main
import "fmt"

func main() {
  var(
      a int
      b int = 0
  )
  fmt.Print("Enter a number: ")
  fmt.Scan(&a)
  for i:= 1; i <= a; i++ {
      if a % i == 0{
          b++
      }
  }
  if b == 2{
      fmt.Print(a," is Prime")
  }else{
      fmt.Print(a," is Not Prime")
  }
  }
