package main

import "fmt"

func main() {
     for x := 4; x <=6; x++{
       fmt.Println("this is", x)
       switch x {
       case 4:
          fmt.Println("four")
       case 5:
          fmt.Println("five")
       case 6:
          fmt.Println("six")
       }
     }
}
