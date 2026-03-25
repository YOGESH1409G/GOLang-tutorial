package main

import (
"fmt"
"maps"
)

func main(){

	// A map is a collection of key-value pairs. It is also known as a dictionary in 
	// other programming languages. In Go, we can create a map using the make function or 
	// using a map literal.
	 m:= make(map[string]int)

	 m["yogesh"] = 137
	 m["bhumica"] = 032

	 fmt.Println("map:", m)

	 // we can also create a map using a map literal
	 m2 := map[string]int{
		 "yogesh": 137,
		 "bhumica": 032,
	 }
	 fmt.Println("map literal:", m2)

    //Get a value for a key with name[key].

	 v1:= m["yogesh"]
	 fmt.Println("value of key 'yogesh':", v1)
	 
	//If the key doesn’t exist, the zero value of the value type is returned.

	v3 := m["ankit"]
	fmt.Println("value of key 'ankit':", v3)

	delete(m,"bhumica") // this will delete the key "bhumica" from the map m (delete(mapName, key))
	fmt.Println("map after deleting key 'bhumica':", m)

	clear(m)// this will clear the map m (clear(mapName))
	fmt.Println("map after clearing:", m)

	//The maps package contains a number of useful utility functions for maps

	m3 := map[string]int{
		 "yogesh": 137,
		 "bhumica": 032,
	}

	if maps.Equal(m2,m3){
	fmt.Println("m2 and m3 are equal")
}


	 
}