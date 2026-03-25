package main

import "fmt"

func main(){
	nums := [] int {10, 20, 30, 40, 50}
	sum:= 0
	/*range on arrays and slices provides both the index and value for each entry. 
	Above we didn’t need the index, so we ignored it with the blank identifier _. 
	Sometimes we actually want the indexes though.*/
	for _, v := range nums {
		sum += v
	}
	fmt.Println("sum is =",sum)

	// for i:= range nums{
	// 	fmt.Println(nums[i])
	// }

	for i , num := range nums{
		if num == 30{
			fmt.Println("index of 30 is =",i)
		}
	}

	//maps are randomly ordered in Go, so we can use range to iterate over 
	// the map and get the key and value of each entry in the map
	maps := map[string]int{
		"yogesh": 10,
		"gupta": 20,
		"iiit": 30,
	}
	for h,v:= range maps{
		fmt.Printf("%s -> %d\n", h, v)
	}
	
	for k,_ := range maps{
		fmt.Println("key is =",k)
	}

	for _, v := range maps{
		fmt.Println("value is =",v)
	}

}