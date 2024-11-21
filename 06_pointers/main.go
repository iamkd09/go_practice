package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting Pointer")
	var ptr *string
	fmt.Println("Value in starting is :", ptr)
	mystr := "kanishk"
	fmt.Println("Original string:", mystr)
	myPtr := &mystr
	fmt.Println("Pointer value:", myPtr)
	fmt.Println("Pointer dereferencing:", *myPtr)
	fmt.Println("Changing the string using pointer")
	*myPtr = *myPtr + " dixit"
	print(mystr)
}
