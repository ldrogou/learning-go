package main

import "fmt"

func main() {

	names := []string{"laurent", "Adeline", "Iris", "Maelle"}

	for i, name := range names {
		fmt.Printf("le nom %s est à l'indice %d \n", name, i)
	}

	i := 42
	str := &i
	ptr := *str

	fmt.Printf("valeur de i : %v \n", i)
	fmt.Printf("adresse de i %v dans str : %v \n", str, *str)
	fmt.Printf("valeur contenue à l'adresse str %v dans ptr : %v \n", &ptr, ptr)
}
