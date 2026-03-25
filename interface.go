package main
import "fmt"

//“Any type that has these two methods automatically becomes a geometry.”
type geometry interface{
	area() float64
	perimeter() float64
}

type rectangle struct{
	width, height float64
}

type circle struct{
	radius float64
}

func (r rectangle) area() float64{
	return r.width * r.height
}

func (r rectangle) perimeter() float64{
	return 2*(r.width + r.height)
}

func (c circle) area() float64{
	return 3.14 * c.radius * c.radius
}

func (c circle) perimeter() float64{
	return 2*3.14*c.radius
}
//g is not a specific type, it can be any type that implements the geometry interface.
//So at runtime:
//g can be a rectangle
//OR g can be a circle

func printGeometry(g geometry){
	fmt.Printf("Area: %.2f\n", g.area())
	fmt.Printf("Perimeter: %.2f\n", g.perimeter())
}	


/*How Go decides what to call?
👉 Go looks at the actual type inside g
If g contains rectangle → call rectangle.area()
If g contains circle → call circle.area()*/


func main(){
	rect := rectangle{width: 5, height: 3}
	circ := circle{radius: 4}
//Go creates an interface value like
// g = {
//     type: rectangle
//     value: {width: 5, height: 3}
// }
	fmt.Println("Rectangle:")
	printGeometry(rect)

	fmt.Println("\nCircle:")
	printGeometry(circ)	
}