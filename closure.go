package main

import "fmt"

func seq() func() int{
	i:=0
	return func() int{
		i++
		return i
	}
}
func main(){
 numsseq := seq()

 fmt.Println(numsseq())
 fmt.Println(numsseq())
 fmt.Println(numsseq())	

}