package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// MenuUsuario gerencia as opções disponíveis para o usuário logado.
func MenuUsuario() {
	scanner := bufio.NewScanner(os.Stdin) // Scanner para capturar linhas completas de entrada
	for {
		fmt.Println("\nMenu do Usuário")
		fmt.Println("1. Criar Arquivo/Diretório")
		fmt.Println("2. Apagar Arquivo/Diretório")
		fmt.Println("3. Listar Todos os Arquivos Criados por Usuários")
		fmt.Println("4. Apagar Diretório com --force")
		fmt.Println("5. Logout")
		fmt.Print("Escolha uma opção (ou insira um comando direto como 'apagar diretorio <caminho> --force'): ")

		// Lê a entrada completa do usuário
		scanner.Scan()
		entrada := scanner.Text()
		entrada = strings.TrimSpace(entrada) // Remove espaços extras

		// Verifica se o comando é "apagar diretorio ... --force"
		if strings.HasPrefix(entrada, "apagar diretorio") && strings.HasSuffix(entrada, "--force") {
			ApagarDiretorioForce(entrada)
			continue
		}

		// Processa opções numéricas
		switch entrada {
		case "1":
			fmt.Print("Digite o caminho do arquivo ou diretório a ser criado: ")
			scanner.Scan()
			caminho := scanner.Text()
			CriarArquivo(caminho)
		case "2":
			fmt.Print("Digite o caminho do arquivo ou diretório a ser apagado: ")
			scanner.Scan()
			caminho := scanner.Text()
			ApagarArquivo(caminho)
		case "3":
			ListarTodosArquivos()
		case "4":
			fmt.Print("Digite o caminho do diretório a ser apagado com --force: ")
			scanner.Scan()
			comando := "apagar diretorio " + scanner.Text() + " --force"
			ApagarDiretorioForce(comando)
		case "5":
			fmt.Println("Saindo do menu do usuário...")
			return
		default:
			fmt.Println("Opção inválida!")
		}
	}
}
