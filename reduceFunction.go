package main
import "fmt"

func gcd(a,b int) int {
  if a == 0 {
    return b
  }
  return gcd(b % a, a)
}

func productReduce(n int,num []int,den []int)  {
  nnum := 1
  nden := 1


  for i :=0; i<n; i++{
    nnum = nnum * num[i]
    nden = nden * den[i]
  }

  CGCD := gcd(nnum, nden)

  nnum = nnum / CGCD
  nden = nden / CGCD

  fmt.Print(nnum, "/", nden)
}
func main()  {
  /*var n int
  fmt.Println("Enter a number")
  fmt.Scanf("%d",&n)*/
  n := 3

  fmt.Println("Enter the numeraters")
  num := make([]int, n)
  for i := 0; i < n; i++ {
      fmt.Scanln(&num[i])
  }

  fmt.Println("Enter the denominaters")
  den := make([]int, n)
  for i := 0; i < n; i++ {
      fmt.Scanln(&den[i])
  }

  productReduce(n, num, den)
}
