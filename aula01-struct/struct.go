package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

//metodo: funcao com receiver (receptor)

func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {

	var produto1 produto

	produto1 = produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}

	fmt.Println(produto1, produto1.precoComDesconto())

	produto2 := produto{"notebook", 5432.3, 0.9}

	fmt.Println(produto2)
}
