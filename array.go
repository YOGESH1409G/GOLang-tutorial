package main

import "fmt"

func main(){
/*
	Here we create an array a that will hold exactly 5 ints. 
	The type of elements and length are both part of the array’s type.
	 By default an array is zero-valued, which for ints means 0s.
	*/
	var a [5]int
	fmt.Println(" array is = ", a)
/*We can set a value at an index using the array[index] = value syntax, 
and get a value with array[index].*/
	a[4] = 200
	fmt.Println("array at 4th index is =",a[4])
	fmt.Println(" array is = ", a)

fmt.Println("len :", len(a))


/*Use this syntax to declare and initialize an array in one line.*/

b := [4]int{3,2,4,5}
fmt.Println("array b is = " , b)

c := [...]int{3,2,4,5}
fmt.Println("array c is = " , c)
 /* yeha per jab hum ( :something ) like rahe hai iska matlab yeh hai ki iss index per yeh
  value hogi aur bakki previous index per 0 rahega*/

 d:= [...]int{100,400, 500, 5: 300}
 fmt.Println("idx:", d)
/*Array types are one-dimensional, 
but you can compose types to build multi-dimensional data structures.*/
 var d2 [3][4]int
for i:= range 3 {
	for j:= range 4{
	d2[i][j] = i+j
}
}
fmt.Println("2d array is = ", d2)
/* we can create and also change the value of the existing array */
d2 = [3][4]int{
	{4,3,5,2},
	{4,4,4,4},
	{6,7,8,9},
}
fmt.Println("new d2 is = ", d2)

/* we can also change the value of the existing array by using the index */
d2[0][0] = 100
fmt.Println("new d2 is = ", d2)


}