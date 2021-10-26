package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       // campos anonimos - tudo que eh de carro fica disponivel dentro de ferrari
	turboLigado bool
}

func main() {

	f := ferrari{}

	f.nome = "F40"
	f.velocidadeAtual = 0
	f.turboLigado = true

	fmt.Println(f)

}
