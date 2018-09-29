package main

import "fmt"
import "time"

func main() {

time := time.Now()
       switch {
       case time.Hour() < 12: 
          fmt.Println("It's before noon") 
       default:
          fmt.Println("It's afternoon")
      }
}
