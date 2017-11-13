package main

import (
    "fmt"
)

func main() {
    fmt.Println("Inches to centimeter program")
    fmt.Print("Enter a value: ")
    var inchVal float32 
    fmt.Scanf("%f", &inchVal)
    
    centiVal := inchVal * 2.54
    fmt.Printf("Your height in centimeters is: %.2f \n",centiVal)
}