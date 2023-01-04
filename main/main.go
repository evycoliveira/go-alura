package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
)

func main() {
	exibeNomes()
	exibeIntroducao()

	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa")
			// Indica que o usuário deseja sair com sucesso
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
			// Indica que ocorreu uma ação inesperada
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Evelyn"
	versao := 1.1
	fmt.Println("Olá, Sra.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://www.alura.com.br/"
	sites[2] = "https://www.caelum.com.br/"

	fmt.Println(sites)

	site := "https://random-status-code.herokuapp.com/"
	// Realiza a consulta na URL atribuida a variável
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}

func exibeNomes() {
	nomes := []string{"Evelyn", "Amanda", "João"}
	fmt.Println("O meu slice tem", len(nomes))
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")
	// Adicionando um nome ao slice
	nomes = append(nomes, "Aparecida")
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")
}
