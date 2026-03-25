package main

import "fmt"


func factorial(n int) int{
	//base case
	if n ==0{
		return 1
	}
	//recursive case
	return n*factorial(n-1)
}
	
	
func main(){
	fmt.Println(factorial(5))
}
