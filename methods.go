package main

import "fmt"

type rectangle struct{
	length int
	width int 
}
//This area method has a receiver type of *rect.

func (r *rectangle) area() int {
	return r.length * r.width
}
//Methods can be defined for either pointer or value receiver types. 
// Here’s an example of a value receiver.
func (r rectangle) perimeter() int {
	return 2*(r.length + r.width)
}
/* /Go automatically handles conversion between values and pointers for method calls. 
 You may want to use a pointer receiver type to avoid copying on method calls or to allow 
 the method to mutate the receiving struct.*/
func main(){
	rect := rectangle{length: 10, width: 5}

	fmt.Println("Area of rectangle:", rect.area())
	fmt.Println("Perimeter of rectangle:", rect.perimeter())

	rectptr := &rect

	fmt.Println("Area of rectangle using pointer:", rectptr.area())
	fmt.Println("Perimeter of rectangle using pointer:", rectptr.perimeter())
}