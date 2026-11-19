package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//pendencia: será q n tem um jeito de repetir menos esses código, as 2 funções que chamam APIs tem basicamente o mesmo código

func searchInBrasilApi(zipCode string) (Address, error) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", zipCode)

	response, err := http.Get(url)
	if err != nil {
		return Address{}, err
	}

	//agenda fechamento para o fim da função, segundo o stackOverflow não fechar pode causar falta de conexão ou memory leak
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Address{}, fmt.Errorf("Status inesperado ao chamar a API do BrasilApi: %d", response.StatusCode)
	}

	var data brasilApiResponse
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return Address{}, err
	}

	return Address{
		ZipCode:      data.ZipCode,
		Street:       data.Street,
		Neighborhood: data.Neighborhood,
		City:         data.City,
		State:        data.State,
	}, nil
}

func searchInViaCep(zipCode string) (Address, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", zipCode)

	response, err := http.Get(url)
	if err != nil {
		return Address{}, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Address{}, fmt.Errorf("Status inesperado ao chamar a API do searchInViaCep: %d", response.StatusCode)
	}

	var data viaCepResponse
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return Address{}, err
	}

	return Address{
		ZipCode:      data.ZipCode,
		Street:       data.Street,
		Neighborhood: data.Neighborhood,
		City:         data.City,
		State:        data.State,
	}, nil
}

type Address struct {
	ZipCode      string
	Street       string
	Neighborhood string
	City         string
	State        string
}

type brasilApiResponse struct {
	ZipCode      string `json:"cep"`
	Street       string `json:"street"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
}

type viaCepResponse struct {
	ZipCode      string `json:"cep"`
	Street       string `json:"logradouro"`
	Neighborhood string `json:"bairro"`
	City         string `json:"localidade"`
	State        string `json:"uf"`
}