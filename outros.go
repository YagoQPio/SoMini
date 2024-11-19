package main

import (
	"fmt"
)

// MenuUsuario gerencia as opções disponíveis para o usuário logado.
func MenuUsuario() {
	for {
		fmt.Println("\nMenu do Usuário")
		fmt.Println("1. Criar Arquivo/Diretório")
		fmt.Println("2. Apagar Arquivo/Diretório")
		fmt.Println("3. Listar Arquivos/Diretórios")
		fmt.Println("4. Logout")
		fmt.Print("Escolha uma opção: ")

		var opcao int
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			fmt.Print("Digite o caminho do arquivo ou diretório a ser criado: ")
			var caminho string
			fmt.Scan(&caminho)
			CriarArquivo(caminho) // Certifique-se que esta função é exportada
		case 2:
			fmt.Print("Digite o caminho do arquivo ou diretório a ser apagado: ")
			var caminho string
			fmt.Scan(&caminho)
			ApagarArquivo(caminho) // Certifique-se que esta função é exportada
		case 3:
			fmt.Print("Digite o diretório para listar os arquivos: ")
			var diretorio string
			fmt.Scan(&diretorio)
			ListarArquivos(diretorio) // Certifique-se que esta função é exportada
		case 4:
			fmt.Println("Saindo do menu do usuário...")
			return
		default:
			fmt.Println("Opção inválida!")
		}
	}
}
