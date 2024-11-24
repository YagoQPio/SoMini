package main

import (
	"fmt"
)

func main() {
	for {
		if !temUsuarios() {
			fmt.Println("Nenhum usuário encontrado. Criando o primeiro usuário...")
			criarUsuario()
		} else {
			fmt.Println("\nBem-vindo ao MiniSO!")
			fmt.Println("Escolha uma opção:")
			fmt.Println("1. Criar Usuário")
			fmt.Println("2. Login")
			fmt.Println("3. Apagar Usuário")
			fmt.Println("4. Sair")

			var opcao int
			fmt.Print("Opção: ")
			fmt.Scanln(&opcao)

			switch opcao {
			case 1:
				criarUsuario()
			case 2:
				if login() {
					menuPrincipal()
				}
			case 3:
				apagarUsuario()
			case 4:
				fmt.Println("Saindo do MiniSO.")
				return
			default:
				fmt.Println("Opção inválida. Tente novamente.")
			}
		}
	}
}
