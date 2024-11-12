package main

import (
	"fmt"
	"os"
)

func main() {
	var choice int
	users, err := loadUsers()
	if err != nil {
		fmt.Println("Erro ao carregar usuários:", err)
		return
	}

	for {
		fmt.Println("Selecione uma opção:")
		fmt.Println("1 - Login")
		fmt.Println("2 - Registrar usuário")
		fmt.Println("3 - Sair")
		fmt.Print("Escolha: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// Efetuar login
			user, err := login(users)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Login bem-sucedido!")
				// Após login, exibe o menu de operações de arquivos
				fileOperationsMenu(user)
			}

		case 2:
			// Registrar novo usuário
			_, err := registerUser()
			if err != nil {
				fmt.Println("Erro no registro:", err)
			} else {
				fmt.Println("Usuário registrado com sucesso!")
			}

		case 3:
			// Sair
			fmt.Println("Saindo...")
			os.Exit(0)

		default:
			fmt.Println("Opção inválida, tente novamente.")
		}
	}
}

// Menu para operações de arquivos
func fileOperationsMenu(user *User) {
	var choice int
	for {
		fmt.Println("\nSelecione uma operação de arquivo:")
		fmt.Println("1 - Criar arquivo")
		fmt.Println("2 - Criar diretório")
		fmt.Println("3 - Listar arquivos")
		fmt.Println("4 - Apagar arquivo")
		fmt.Println("5 - Sair para o menu principal")
		fmt.Print("Escolha: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// Criar arquivo
			var filePath string
			fmt.Print("Digite o caminho do arquivo a ser criado (exemplo: dir1/dir2/arquivo1.txt): ")
			fmt.Scanln(&filePath)
			createFile(user, filePath)

		case 2:
			// Criar diretório
			var dirPath string
			fmt.Print("Digite o caminho do diretório a ser criado (exemplo: dir1/dir2): ")
			fmt.Scanln(&dirPath)
			createDir(user, dirPath)

		case 3:
			// Listar arquivos
			listFiles()

		case 4:
			// Apagar arquivo
			var filePath string
			fmt.Print("Digite o caminho do arquivo a ser apagado (exemplo: dir1/arquivo1.txt): ")
			fmt.Scanln(&filePath)
			deleteFile(user, filePath)

		case 5:
			// Voltar ao menu principal
			return

		default:
			fmt.Println("Opção inválida, tente novamente.")
		}
	}
}
