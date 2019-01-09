package main
import ("fmt"

      )
func wordscore(n int)  {
  fmt.Println("Enter the Strings")
  str := make([]string, n)
  for i := 0; i < n; i++ {
      fmt.Scanln(&str[i])
  }
  count :=0
  sum := 0
  //var userString = "programming is awesome "

  for j := 0; j < n; j++ {

      words := string([]rune(str[j]))
      //fmt.Println(string(words[1]))
      for i :=0; i<len(string(words)); i++{
        if string(words[i])=="a" || string(words[i])=="e" || string(words[i])=="i" || string(words[i])=="o" || string(words[i])=="u" {
            count++
        } //else if (string(words[i])==" " || i+1 == len(string(words))) {

        //}

      }
      if count !=0 {
        if count % 2 ==0{
          sum = sum+2
        }else {
          sum = sum+1
        }
        count = 0
      }
  }
  fmt.Println(sum)

}
func main()  {
  var n int = 0
  fmt.Println("enter number of words + 1 EXample : if requered 3 then enter 4")
  fmt.Scanf("%d",&n)
  wordscore(n)


}
