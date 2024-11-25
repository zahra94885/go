package main

import "fmt"

func main() {
	firstName := "man"
	var lastName string = "moradi"

	fmt.Println(firstName)
	fmt.Println(lastName)

	currentYear := 2023
	var birthYear = 1999
	age := currentYear - birthYear
	fmt.Println(age)
	currentYear = currentYear + 1
	age = currentYear - birthYear
	fmt.Println(age)

	var newRune rune = '$'

	fmt.Println(newRune)
	fmt.Println(string(newRune))

	var newByte byte = 'a'
	fmt.Println(string(newByte))
	firstName1 := "man"
	lastName1 := "moradi"
	fmt.Println(firstName1 + lastName1)

}
