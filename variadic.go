package main

import "fmt"


func add(nums ...int) int{
	result:=0
	for _,num := range nums{
		result += num
	}
	return result
}

func main(){
	sum := add(10,20,30,40,50)
	fmt.Println("Sum is =", sum)

	sum2:= []int{1,2,3,4,5}
	fmt.Println("Sum2 is =", add(sum2...)) // this will pass the slice as variadic arguments to the add function
}