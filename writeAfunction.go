package main

import "fmt"

func main() {

  var input int
  fmt.Println("Enter a year:")
  fmt.Scanf("%d", &input)
  if (input % 4 == 0)
   {
    if (input % 100 == 0)
     {
      if (input % 400 == 0)
      {
        fmt.Println(input , " True")
      }
      else {
       fmt.Println(input , " False")
      }
     }
    else {
       fmt.Println(input , " True")
     }
  }
  else {
    fmt.Println(input , " False")
  }
}
