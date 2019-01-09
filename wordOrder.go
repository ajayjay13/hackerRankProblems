package main
import "fmt"

func productReduce(n int,st []string)  {


  distinct := 0
  k := 0
  for i :=0; i<n-1; i++{
    count :=0

    for j :=i+1; j<n; j++{

      if st[i]==st[j] {
        count =count+1
      }
    }

    if count == 0{
      distinct = distinct+1

    }
  }
  distinct = distinct+1

  occ := make([]int, distinct)
  for i:= 0; i < distinct; i++{
    occ[i] = 1
  }

  for i :=0; i<n-1; i++{
    count :=0
    if st[i]==""{
      continue
    }
    for j :=i+1; j<n; j++{

      if st[i]==st[j] {
        st[j]=""
        count =count+1
        occ[k]= occ[k]+1

      }
    }

    if count == 0{
      occ[k]= 1
    }
    k = k+1
  }


  fmt.Print(distinct,"\n")
  for i:=0; i< distinct; i++ {
    fmt.Print(occ[i]," ")
  }
}
func main()  {
  
  n := 5


  fmt.Println("Enter the Strings")
  str := make([]string, n)
  for i := 0; i < n; i++ {
      fmt.Scanln(&str[i])
  }

  productReduce(n, str)
}
