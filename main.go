package main

import "fmt"

func main() {
	saudation()
	zipCode := showMenu()
	address, err := search(zipCode)

	if err != nil {
		fmt.Println("Algo inesperado aconteceu")
		return
	}
	fmt.Printf("Rua: %s\nBairro: %s\nCidade: %s - %s\n", address.Street, address.Neighborhood, address.City, address.State)

}
