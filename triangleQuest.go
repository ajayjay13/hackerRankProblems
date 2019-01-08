package main
import "fmt"
func main() {

  var input int
  fmt.Println("Enter a number")
  fmt.Scanf("%d", &input)
  fmt.Print("\n")
    for n := 1; n < input; n++ {
        for i:=n; i>=1; i-- {
          fmt.Print(n)
        }
        
        fmt.Print("\n")
    }
}
