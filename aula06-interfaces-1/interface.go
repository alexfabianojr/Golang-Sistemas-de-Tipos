package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

//interfaces sao implementadas implicitamente

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(i imprimivel) {
	fmt.Println(i.toString())
}

func main() {

	var coisa imprimivel = pessoa{"Roberto", "Silva"}
	fmt.Println(coisa.toString())

	imprimir(coisa)

	coisa = produto{"notebook", 2000} // polimorfismo - a mesma variavel consegue ter multiplas formas

	imprimir(coisa)

	p2 := produto{"Calça", 45}

	imprimir(p2)
}
