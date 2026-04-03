package main
import "fmt"
//Type safety ✅
//Prevents mixing with other ints ❌
// type status int

 const (
 	Active status = iota
 	Inactive
 	Pending
 )
 func main(){
    fmt.Println(Pending)
 }
// if we want to print the status in string format then we can use switch case with using 
// String() method to convert the status to string. 
 func (s status) String() string{
 	switch s {
 		case Active:
 			return "Active"
 		case Inactive:
 			return "Inactive"
		case Pending:
			return "Pending"
		default:
			return "Unknown Status"
 	}
 }

