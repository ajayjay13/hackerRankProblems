package main
import "fmt"

func gcd(a,b int) int {
  if a == 0 {
    return b
  }

  return gcd(b % a, a)
}

func productReduce(n int,num []int,den []int)  {
  new_num := 1
  new_den := 1


  for i :=0; i<n; i++{
    new_num = new_num * num[i]
    new_den = new_den * den[i]
  }



  GCD := gcd(new_num, new_den)


  new_num = new_num / GCD
  new_den = new_den / GCD

  fmt.Print(new_num, "/", new_den)
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
  /*var num [] int
  var den [] int*/
  //n := 3






 /*var num = []int{10,20,30}
  var den = []int{10,20,30}*/
  /*num = {1, 2, 5}
  den = {2, 1, 6}*/

/*  var num [3]int
  var den [3]int
   num[0]=1
   num[1]=2
   num[2]=5
   den[0]=2
   den[1]=1
   den[2]=6*/
  productReduce(n, num, den)
}
