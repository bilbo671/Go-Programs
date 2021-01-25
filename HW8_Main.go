#
#<HW8_Main.go>
#
#<Main function for Swift>
#
#NCU.edu
#School of Business and Technology Management
#<TIM-6110>
#
#Author: <William Jenkins>
#Date: <1-24-21>
#
#/
package main
import "fmt"
import "strings"

func main() {

#Hello world!
fmt.Println("Hello, World!")

#Odd Num Func
var numbs = []int{1,2,3,4,5,6,7,8,9}

for _, num := range numbs {
  if num%2 !=0 {
  fmt.Printf("%v,", num)
  }
}

#Explode
var explode = strings.Split("apple", "")
fmt.Printf("%q\n", explode)

#Implode
implode := []string{"a", "p", "p", "l", "e"}
fmt.Println(strings.Join(implode, ""))

