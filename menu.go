package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func menuPrincipal() {
	for {
		fmt.Println("\n=== Menu Principal ===")
		fmt.Println("1. Criar Arquivo")
		fmt.Println("2. Criar Diretório")
		fmt.Println("3. Apagar Arquivo")
		fmt.Println("4. Apagar Diretório")
		fmt.Println("5. Listar Conteúdo")
		fmt.Println("6. Logout")
		fmt.Print("Escolha uma opção: ")

		var opcao int
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			fmt.Print("Digite o nome do arquivo a ser criado: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			nome := scanner.Text()
			criarArquivo([]string{nome})
		case 2:
			fmt.Print("Digite o nome do diretório a ser criado: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			nome := scanner.Text()
			criarDiretorio([]string{nome})
		case 3:
			fmt.Print("Digite o nome do arquivo a ser apagado: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			nome := scanner.Text()
			apagarArquivo([]string{nome})
		case 4:
			fmt.Print("Digite o nome do diretório a ser apagado: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			nome := scanner.Text()
			apagarDiretorio([]string{nome})
		case 5:
			fmt.Print("Digite o diretório para listar (ou pressione Enter para listar o atual): ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			diretorio := scanner.Text()
			if strings.TrimSpace(diretorio) == "" {
				diretorio = "."
			}
			listar([]string{diretorio})
		case 6:
			fmt.Println("Realizando logout...")
			return
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}
