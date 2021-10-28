package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {

	//struct para json

	p1 := produto{1, "Notebook", 2341.0, []string{"Promocoes", "Eletronico"}}

	p1Json, _ := json.Marshal(p1) //retorno -> json (slice de bytes), err

	fmt.Println(string(p1Json)) //convertendo os bytes para string

	var p2 produto
	jsonString := `{"id":2, "nome":"Caneta", "preco":8.9, "tags":["Papelaria", "Importado"]}`

	json.Unmarshal([]byte(jsonString), &p2) //passa a referencia da memoria de p2 para o unmarshal jogar la dentro o objeto

	fmt.Println(p2)
	fmt.Println(p2.Nome)
}
