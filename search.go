package main



func search(zipCode string) (Address, error) {
	type result struct {
		address Address
		err     error
	}
	
	//faço um canal que terá 2 resultados
	ch := make(chan result, 2)

	//faz as duas chamadas à API ao mesmo tempo a primeira a responder é a que "ganha" faz a chamada e guarda o responde em ch 
	go func() {
		address, err := searchInBrasilApi(zipCode)
		ch <- result{address: address, err: err}
	}()

	go func() {
		address, err := searchInViaCep(zipCode)
		ch <- result{address: address, err: err}
	}()

	//é aqui que eu espero o primeiro resultado
	//mas assim se der erro já não falha no primeiro que responder com erro ?
	response:= <-ch
	if response.err != nil{
		return Address{}, response.err
	}

	return response.address, nil
}

	