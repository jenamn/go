package main

import "fmt"

func findbday (x, y int) int { return x - y }

func main() {
	fmt.Printf("Hello Sexy Coach D!!\n")
        myPetType := "dog"
          fmt.Println("Breezy is my",myPetType)
        const bday string = "October 15,1968"
          fmt.Println("My birthday is", bday)
        fmt.Println("This year I will be", findbday(2018, 1968)) 
}
