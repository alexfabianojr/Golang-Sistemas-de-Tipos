package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	intens []item
}

func (p pedido) valorTotal() float64 {

	total := 0.0

	for _, item := range p.intens {
		total += item.preco * float64(item.qtde)
	}

	return total
}

func main() {

	itens := []item{
		item{
			produtoID: 1,
			qtde:      20,
			preco:     876.32,
		},
		item{
			produtoID: 2,
			qtde:      988,
			preco:     12.1,
		},
	}

	pedido := pedido{
		userID: 1,
		intens: itens,
	}

	fmt.Println(pedido.valorTotal())
}
