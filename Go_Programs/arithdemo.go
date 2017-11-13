package main
import "fmt"

func main() {
  var a,b int
  a = 5
  b = 7
  c := a + b
  d := a - b
  e := a * b
  f := b/a
  
  fmt.Printf("%d + %d = %d \n",a,b,c)
  fmt.Printf("%d - %d = %d \n",a,b,d)
  fmt.Printf("%d * %d = %d \n",a,b,e)
  fmt.Printf("%d / %d = %.2f \n",b,a,f)
}
