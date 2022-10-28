package main

import (
	"fmt"
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

	for {
		exibeMenu()
		opcao := leOpcao()

		switch opcao {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("2")
		case 0:
			fmt.Println("0")
			os.Exit(0)
		default:
			fmt.Println("Erro, opção inválida")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Cadu"
	fmt.Println("Olá mundo", nome)
}

func exibeMenu() {
	fmt.Println("1 - Check url")
	fmt.Println("2 - ")
	fmt.Println("0 - Exit")
}

func leOpcao() int {
	var opcaoLida int

	fmt.Scan(&opcaoLida)

	return opcaoLida
}

func iniciarMonitoramento() {
	sites := leSitesNoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			testaSite(site)
		}

		time.Sleep(delay * time.Second)
	}

}

func testaSite(site string) {
	// faz uma request na url informada
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("OK")
	} else {
		fmt.Println("Ocorreu erro:", resp.StatusCode)
	}
}

func leSitesNoArquivo() []string {
	var sites []string
	file, err := os.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Erro:", err)
	}

	fmt.Println(string(file))
	return sites
}
