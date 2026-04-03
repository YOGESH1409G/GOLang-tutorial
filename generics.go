package main
import "fmt"


/*we can define type parameter in this also*/
/*func Print[T any, U any](a T, b U) {
    fmt.Println(a, b)
}*/
//generics using structs
// type number[T any ] struct{
// 	value T
// }

// func (n number[T]) getValue() T {
// 	return n.value
// }

//Generics with Slices
// func PrintSlice[T any](s []T) {
//     for _, v := range s {
//         fmt.Print(v)
//     }
}
// func sum[T int | float64](a, b T) T {
// 	return a + b
// }
func main(){

	// s:= sum(2.7,3.5)
	// fmt.Println("Sum:", s)
	// b:= number[int]{value: 42}
	// fmt.Println("Value:", b.getValue())

	// intSlice := []int{1, 2, 3, 4, 5}
	// stringSlice := []string{"Hello", "Generics", "in", "Go"}

	// fmt.Println("Integer Slice:")
	// PrintSlice(intSlice)

	// fmt.Println("\nString Slice:")
	// PrintSlice(stringSlice)

	
}