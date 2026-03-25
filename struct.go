package main
import "fmt"

//Go’s structs are typed collections of fields. 
// They’re useful for grouping data together to form records.
type person struct{
	name string
	age int
}


func main(){
	p := person{name: "Alice", age: 30}
	fmt.Println(p)

	fmt.Println(person{name: "bob", age: 25})

	// You can also name the fields when initializing a struct.
	fmt.Println(person{age: 40, name: "Charlie"})

	// If the fields are omitted when initializing a struct, they will be set to the zero value of their type.
	fmt.Println(person{name: "Dave"})

	// Accessing struct fields using dot notation
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)

	// You can also modify struct fields using dot notation
	p.age = 31
	fmt.Println("Updated Age:", p.age)

	// Structs can also be used with pointers
	pPointer := &p
	fmt.Println("Name through pointer:", pPointer.name)
	fmt.Println("Age through pointer:", pPointer.age)

	// Modifying struct fields through a pointer
	pPointer.age = 32
	fmt.Println("Updated Age through pointer:", pPointer.age)

}
