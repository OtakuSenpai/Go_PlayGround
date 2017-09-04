package main

import ( 
    "fmt"
    "math/rand"
)

func main() {
  fmt.Println("This is an array slice demo program.")
  var ( 
     array1[5] int
     array2[][] int
     row int
     col int
  )
  
  fmt.Println("Print the data for first array: \n")
  for i:=0;i<5;i++ {
    fmt.Scanf("%d", &array1[i])  
  }
  
  fmt.Println("Print the row and column: \n")
  fmt.Scanf("%d", &row) 
  fmt.Scanf("%d", &col)
  
  array2 = make([][]int,row)   

  fmt.Println("Entering the data for second array: \n") 
  for j:=0;j<row;j++ {
    if array2[j] == nil {
      array2[j] = make([]int,col)
    } 
    for k:=0;k<col;k++ {
      array2[j][k] = rand.Intn(100)
    }
  }

  fmt.Println("Printing the data of first array: ")
  for i:=0;i<5;i++ {
    fmt.Printf("%d ",array1[i])
  } 
  fmt.Print("\n")
 
  fmt.Println("And for second array: ")
  for j:=0;j<row;j++ { 
    for k:=0;k<col;k++ {
      fmt.Printf("%d ",array2[j][k])
    }
    fmt.Print("\n")
  }
}        
