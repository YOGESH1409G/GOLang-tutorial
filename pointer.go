package main

import "fmt"

// pass by reference using pointers
func changeValue(x *int){
	*x = 20. //*x means accessing the value at the address x is pointing to 
}
func main(){
	x:= 10
	p:= &x

	fmt.Println("Value of x is", x)
	fmt.Println("Address of x is", &x)
	fmt.Println("Value of p is", p)
	fmt.Println("Value at address p is", *p)

	changeValue(&x)
	fmt.Println("Value of x after change is", x)

	var q *int
	fmt.Println(q) // nil

	z:= new(int)
	fmt.Println(z)// new(int) will allocate memory for an int and return a pointer to it
	fmt.Println(*z) // 0, default value of int is 0
}