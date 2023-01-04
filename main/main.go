package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeIntroducao()
	exibeMenu()

	nome, idade := devolveNomeEIdade()
	fmt.Println("Meu nome é:", nome, "e tenho", idade, "anos.")

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
	site := "https://www.alura.com.br"
	// Realiza a consulta na URL atribuida a variável
	resp, _ := http.Get(site)
	fmt.Println(resp)
}

func devolveNomeEIdade() (string, int) {
	nome := "Evelyn"
	idade := 25
	return nome, idade
}
