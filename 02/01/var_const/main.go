package main

import "fmt"

func main() {
	var name string
	name = "Salvador"
	// you can also declare with := to set the type
	surname := "Muñoz"
	// several declarations at the same time
	country, city, age := "Spain", "Granada", 28

	fmt.Println(name, surname, age, country, city)
}
