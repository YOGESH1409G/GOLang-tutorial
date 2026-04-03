package main
import "fmt"

type MyError struct {
	Message string
	Code int
}
func (e MyError) Error() string{
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}
func main(){
	err:= MyError{Message: "Something went wrong", Code: 500}
	fmt.Println(err.Error())
}