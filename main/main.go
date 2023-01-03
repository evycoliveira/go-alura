package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Evelyn"
	reflect.TypeOf(nome)
	versao := 1.1
	fmt.Println("Olá, Sra.", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")

	var comando int
	//fmt.Scan(&comando)
	fmt.Scanf("%d", &comando)
	fmt.Println("O endereço da minha variável comando é", &comando)
	fmt.Println("O comando escolhido foi", comando)
}
