package main

import (
	"fmt"
	"strings"
	"unicode"
)

func saudation() {
	var name string
	fmt.Println("Digite seu nome: ")
	fmt.Scan(&name)
	fmt.Println("Saudações sr.", name)
}

func showMenu() string {
	var zipCode string
	fmt.Println("Digite o cep que deseja visualizar informações men: ")
	fmt.Scan(&zipCode)
	return zipCodeReadAndValidator(zipCode)
}

func zipCodeReadAndValidator(zipCode string) string {
	
	for{
		zipCodeWithouHyphen := strings.ReplaceAll(zipCode, "-", "")

		if isValidZipCode(zipCodeWithouHyphen){
			return zipCodeWithouHyphen
		}

		fmt.Println("O CEP que você digitou é inválido!")
		fmt.Println("Digite o CEP novamente (ele deve conter 8 dígitos):")
		fmt.Scan(&zipCode)
	}
}

func isValidZipCode(zipCode string) bool {
	if len (zipCode) != 8{
		return false
	}

	for _, digit := range zipCode {
		if !unicode.IsDigit(digit){//p ter certeza que é um númeo
			return false
		}
	}
	return true
}
