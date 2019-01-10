package main
import "fmt"
func main()  {
  count := 1
  var input int
  fmt.Println("Enter the number")
  fmt.Scanf("%d", &input)
  var n, r int = 0, 0

  var r1, r2 int
  for {
    r = input % 10
    input = input / 10
    n = n * 10 + r
    if input == 0 {
      break
    }
  }

  r1 = n % 10
  n = n/10
  for {
    r2 = n % 10
    n = n / 10
    if r1 == r2 {
      count++
    }else {
      fmt.Print("(",count,",",r1,")")
      r1 = r2
      count = 1
    }
    if n == 0 {
      break
    }
  }
  fmt.Print("(",count,",",r1,")")
}
