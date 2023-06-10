/*** This Program Gives Basic Introduction Of Go Langauge ***/

package main 

import "fmt"

func main () {
  
  fmt.Println("Hello Gophers !!!")
  
  var i int = 0
  var j float64
  var s string 
  
  k := 10
  l := 111.34
  name := "Sandeep"
  num := 0
  j = 3.17
  s = "Go Langauge"
  
  fmt.Println("Enter The Number : ")
  fmt.Scanln(&num)
  fmt.Printf("Number = %d\n",num)
}
