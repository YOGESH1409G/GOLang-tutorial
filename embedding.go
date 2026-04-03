package main

import "fmt"

type base struct {
	num int
}
//if we do *base then go Automatic address-taking when needed and
//  dereferencing when needed.
//Even though describe() needs *base, Go converts:
// co.describe() → (&co.base).describe()

func (b base) describe()string{
	return fmt.Sprintf("Base struct with num = %d", b.num)
}

type container struct {
	base // embedding base struct
	name string
}

func main(){
	co:= container{
		base: base{num:45},
		name: "yogesh",
	}
	 fmt.Printf("co={num: %v, name: %v}", co.num, co.name)
	 fmt.Println()
	 fmt.Println("co.describe() =", co.describe())
	 fmt.Println("co.base.describe() =", co.base.describe())
	 fmt.Println("co.base.num =", co.base.num)
	 fmt.Println("co.num =", co.num)

}