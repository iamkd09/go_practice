package main

import "fmt"

// public variable using constant
const publicKey = "fdhg673geghsd36"

func main() {

	// using the public key
	fmt.Println("Public Key:", publicKey)
	fmt.Printf("Variable is of type: %T \n", publicKey)

	// string
	var firstname string = "Kanishk"
	fmt.Println("First Name: ", firstname)
	fmt.Printf("Variable is of type: %T \n", firstname)

	// wallrus operator :=
	lastname := "Dixit"
	fmt.Println("Last Name: ", lastname)
	fmt.Printf("Variable is of type: %T \n", lastname)

	// unsigned integer small
	var smallVal uint8 = 255
	fmt.Println("Small Value: ", smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	// unsigned float small
	var floatVal float64 = 255.45544511254451885
	fmt.Println("Float Value: ", floatVal)
	fmt.Printf("Variable is of type: %T \n", floatVal)

	// boolean
	var isLoggedIn bool = false
	fmt.Println("Is Logged In: ", isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	// implicit type
	var website = "www.google.com"
	fmt.Println("Website: ", website)
	fmt.Printf("Variable is of type: %T \n", website)

	// default value
	var defaultVal int
	fmt.Println("Default Value: ", defaultVal)
	fmt.Printf("Variable is of type: %T \n", defaultVal)

}
