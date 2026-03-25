package main

import "fmt"


func add(a int , b int)int{
	return a+b
}
/*When you have multiple consecutive parameters of the same type, you may omit the type 
 name for the like-typed parameters up to the final parameter that declares the type.*/
func add3(a,b,c int)int{
		return a+b+c
}

func main(){
  sum := add(10,3)
  fmt.Println("Sum is =", sum)

  sum3 := add3(3,4,5)
  fmt.Println("Sum3 is =", sum3)
}