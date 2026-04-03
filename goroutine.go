package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i:= range 3{
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("start")

	go f("goroutine")

	 go func (msg string){
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")


}