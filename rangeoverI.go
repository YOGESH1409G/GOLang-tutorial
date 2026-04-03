package main
import (
	"fmt"
	"iter"
	"slices"
	"Strings"
)

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val T
}

func ( lst * List[T]) push(v T) {
	if lst.tail == nil {
		lst.head = &element[T ]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

func getFib()iter.Seq[int]{
	return func(yield func(int) bool){
		a, b := 0, 1
		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main(){
	lst := List[int]{}
	lst.push(1)
	lst.push(2)
	lst.push(3)

	for e := range lst.All() {
		fmt.Println(e)
	}

	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	for part:= range strings.SplitSeq("yogesh", ",") {
		fmt.Println(part)
	}

	for f := range getFib() {
		if f > 100 {
			break
		}
		fmt.Println(f)
	}
}