package main

import "fmt"

func addAndMultiply(a, b int) (int, int) {
	sum := a + b
	product := a * b
	return sum, product
}

//value + error pattern.This is VERY IMPORTANT in Go
//Almost every real-world function follows this pattern.
func divide(a,b int) (int,error){
		if b == 0{
			return 0, fmt.Errorf("denominator cannot be zero")
		}
		return a/b, nil
	}


func main(){
	sum , product := addAndMultiply(2,3)
	fmt.Println("Sum is =", sum)
	fmt.Println("Product is =", product)

   if result, err := divide(10, 7); err == nil {
	fmt.Println("Result is =", result)
   }
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("Result is =", result)
	// }
}