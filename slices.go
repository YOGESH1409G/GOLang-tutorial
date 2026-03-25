package main

import (
	"slices"
	"fmt"
)

func main(){

	// How to declare a slice

	// Method 1: Using literal
	s := []int{10,20,30}
	fmt.Println(s)

	// 	Method 2: Using make function
	s1 := make([]int , 3)
	fmt.Println(s1)
	fmt.Println(len(s1)) // this will give the length of the slice
	fmt.Println(cap(s1)) // this will give the capacity of the slice

	// Method 3: From an array
	a := [5]int{1,2,3,4,5}
	b:= a[1:4] // this will create a slice from index 1 to 3 (4 is exclusive)
	// we can also do a[1:] which will create a slice from index 1 to the end of the array
	/*we can also do a[:4] which will create a slice from the beginning of the array to 
	 index 3 (4 is exclusive)*/
	fmt.Println(b)

	f := []int{10,20,30}
	f = append(f,40,50) // this will add 40 and 50 to the slice f
	fmt.Println(f)
    
	// we can also copy a slice to another slice using the copy function
	c:= make([]int , len(f))
	c = append(c,0,3,2)
	copy(c,f) // this will copy the elements of slice f to slice c
	fmt.Println(c)

	t:= []string{"yogesh","gupta"}

	t2:= []string{"yogesh","gupta"}
	if slices.Equal(t,t2) {
		fmt.Println("t and t2 are equal")
	}


}