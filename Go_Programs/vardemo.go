package main
import "fmt"

func main() {
  //declare variables
  var n int
  var s string
  var f float32
  //assign values
  n = 10
  s = "Hello,this is a demo program to work on variables..."
  f = 1.2
  f = f + 4
  //print them
  fmt.Println(s)
  fmt.Println("f= ",f)
  fmt.Println("n= ",n)
  //another way of assign and declaring at once
  var city string = "Kolkata"
  var name string = "Neel"
  var food string = "tea"
  var result string = name + ",who lives in " + city + " loves " + food
  fmt.Println(result)
  
  //use type inference with the := operator
  choice := true
  str1 := "have some cake!!!"
  
  fmt.Println("choice= ",choice)
  fmt.Println("str1= ",str1)

  //declare multipe variables at once
  var (
    fun int
    baz float32
    foo string
  )
  fun = 1
  baz = 4.3
  foo = "bye..."
  fmt.Println("fun= ",fun)
  fmt.Println("baz= ",baz)
  fmt.Println(foo) 
}
