package main

import "fmt"

func main() {
    s := "Hello नमस्ते"

    fmt.Println("String:", s)

    // Length in bytes
    fmt.Println("len(s):", len(s))

    // Length in characters (runes)
    fmt.Println("len([]rune(s)):", len([]rune(s)))

    m := "नम"

    fmt.Println("len(m):", len(m))                 // 6 bytes
    fmt.Println("len([]rune(m)):", len([]rune(m))) // 2 characters

	
}