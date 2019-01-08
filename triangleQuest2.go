package main
import "fmt"
func main() {

  var input int
  fmt.Println("Enter a number")
  fmt.Scanf("%d", &input)
  fmt.Print("\n")
    for n := 1; n <= input; n++ {
        for i:=1; i<=n; i++ {
          fmt.Print(i)
        }
        for i:=n-1; i>=1; i-- {
          fmt.Print(i)
        }
        fmt.Print("\n")
    }
}
