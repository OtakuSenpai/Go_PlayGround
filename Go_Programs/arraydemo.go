package main

import (
    "fmt"
    "math/rand"
)

func main() {
    fmt.Println("Fun with arrays :D")
    var (
       cityNo[3] int
       cityName[3] string
       matrix[3][3] int
       i,j,k int
    )
    
    fmt.Println("Input some data!")
    for i=0;i<3;i++ {
       cityNo[i] = i
       cityName[i] = fmt.Sprintf("Enter the name of city %d: ",i)
    }
   
    for j=0;j<3;j++ {
       for k=0;k<3;k++ {
          matrix[j][k] = rand.Intn(100)
       }
    }
    
    fmt.Println("Output now...")
    for i=0;i<3;i++ {
       fmt.Println(cityNo[i])
       fmt.Println(cityName[i])
    }
    for j=0;j<3;j++ {
       for k=0;k<3;k++ {
          fmt.Println(matrix[j][k])
       }
    }
}