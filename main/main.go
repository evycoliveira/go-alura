package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	monitoramentos = 3
	delay          = 5
)

func main() {
	exibeIntroducao()
	leSitesDoArquivo()

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
	fmt.Println(" ")

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := leSitesDoArquivo()

	// For para repetição do monitoramento
	for i := 0; i < monitoramentos; i++ {
		// For para percorrer a lista de sites
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		// Intervalo de 5 segundos para cada vez que todos os sites são momitorados
		time.Sleep(delay * time.Second)
		fmt.Println(" ")
	}
	fmt.Println(" ")
}

func testaSite(site string) {
	// Realiza a consulta na URL atribuida a variável
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um errro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}

func leSitesDoArquivo() []string {
	var sites []string

	//arquivo, err := os.Open("sites.txt")
	// Essa função do pacote ioutil traz o conteúdo do arquivo como array de bytes
	arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	// Conversão do arquivo de um array de bytes para string
	fmt.Println(string(arquivo))
	return sites
}
