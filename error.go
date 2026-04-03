package main
import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator cannot be zero")
	}
	return a / b, nil
}

var Erroroutoftea = errors.New("no more tea available")
var Erroroutofcoffee = errors.New("can't boil water")

func maketea(arg int) (error){
	if arg ==2{
		return Erroroutoftea
	} else if arg == 4 {
		return fmt.Errorf("making tea: %w ", Erroroutofcoffee)
	}
	return nil
}
func main(){
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	for i := range 5 {
		if err := maketea(i); err != nil {
			if errors.Is(err, Erroroutoftea) {
				fmt.Println("Error: Out of tea")
			} else if errors.Is(err, Erroroutofcoffee) {
				fmt.Println("Error: Out of coffee")
			} else {
				fmt.Println("Unknown error:", err)
			}
			continue
		}
		fmt.Println("tea is ready")
	}

}