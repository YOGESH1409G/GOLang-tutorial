package main

import ("fmt"
// "time"
)

func main(){

	// i := 3
	// i := []int{10, 20, 30}
	// fmt.Print("write"," ", i[0] ," ", "as", " ")
	// switch i[0] {
	// case 10:
	// 	fmt.Println("ten")

	// case 20:
	// 	fmt.Println("twenty")

	// case 30:
	// 	fmt.Println("thirty")
	// }
/*You can use commas to separate 
multiple expressions in the same case statement. 
We use the optional default case in this example as well.*/


	// switch time.Now().Weekday(){
	// case time.Sunday, time.Saturday:
	// 	fmt.Println(" it is a weekend")

	// default:
	// 	fmt.Println("it is a weekday")
	
	// }
/*switch without an expression is an alternate way to express 
if/else logic. Here we also show how the case expressions can be non-constants.*/
	// t := time.Now()
	// switch {
	// case t.Hour() < 14:
	// 	fmt.Println("good morning")

	// case t.Hour() < 18:
	// 	fmt.Println("good afternoon")
	// }

	/*A type switch compares types instead of values. 
	You can use this to discover the type of an interface value. 
	In this example, the variable t will have the type corresponding to its clause.*/

	//    whatAmI := func(i interface{}) {
    //     switch t := i.(type) {
    //     case bool:
    //         fmt.Println("I'm a bool")
    //     case int:
    //         fmt.Println("I'm an int")
    //     default:
    //         fmt.Printf("Don't know type %T\n", t)
    //     }
    // }
    // whatAmI(true)
    // whatAmI(1)
    // whatAmI("hey")
}
