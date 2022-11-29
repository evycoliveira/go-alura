package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome = "Evelyn"
	var idade = 25
	var versao = 1.1

	fmt.Println("Olá, Sra.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(nome))
	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(idade))
	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(versao))
}
