package main

import "fmt"

func findbday (x, y int) int { return x - y }

func main() {
	fmt.Printf("Hello world!!\n")
        myPetType := "dog"
          fmt.Println("Breezy is my",myPetType)
        const bday string = "in October"
          fmt.Println("My birthday is", bday)
        fmt.Println("This year I will be", findbday(2018, 1968)) 
}
