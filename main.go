package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Função para executar um processo e exibir o PID
func executarProcesso(action func() error) error {
	pid := os.Getpid()
	fmt.Printf("Processo iniciado com PID: %d\n", pid)

	err := action()
	if err != nil {
		return fmt.Errorf("erro no processo %d: %v", pid, err)
	}
	return nil
}

func main() {
	fmt.Println("Bem-vindo ao MiniSO!")

	// Verifica se há usuários cadastrados
	if !UsuariosExistem() {
		err := executarProcesso(CriarUsuario)
		if err != nil {
			fmt.Printf("Erro ao criar usuário: %v\n", err)
			return
		}
		fmt.Println("Usuário criado com sucesso!")
	}

	// Loop para login ou gerenciamento de usuários
	for {
		fmt.Println("1. Fazer login")
		fmt.Println("2. Adicionar novo usuário")
		fmt.Println("3. Finalizar programa")
		fmt.Print("Escolha uma opção: ")

		var opcao string
		fmt.Scanln(&opcao)

		switch strings.TrimSpace(opcao) {
		case "1":
			// Login do usuário
			var user *Usuario
			var err error
			for {
				err = executarProcesso(func() error {
					var loginErr error
					user, loginErr = LoginUsuario()
					return loginErr
				})
				if err != nil {
					fmt.Printf("Erro no login: %v. Tente novamente.\n", err)
					continue
				}
				break
			}
			fmt.Printf("Login bem-sucedido! Bem-vindo, %s.\n", user.Username)

			// Inicia o shell principal
			shell(user)
		case "2":
			// Adicionar novo usuário
			err := executarProcesso(CriarUsuario)
			if err != nil {
				fmt.Printf("Erro ao adicionar usuário: %v\n", err)
			} else {
				fmt.Println("Usuário adicionado com sucesso!")
			}
		case "3":
			// Finalizar o programa
			fmt.Println("Finalizando o programa...")
			os.Exit(0)
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}

// Shell principal
func shell(user *Usuario) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("MiniSO> ")
		if !scanner.Scan() {
			break
		}
		comando := strings.TrimSpace(scanner.Text())

		if comando == "sair" {
			fmt.Println("Voltando ao menu de login...")
			return
		}

		if comando == "finalizar" {
			fmt.Println("Finalizando o programa...")
			os.Exit(0)
		}

		err := executarProcesso(func() error {
			return processarComando(comando, user.Username)
		})
		if err != nil {
			fmt.Printf("Erro: %v\n", err)
		}
	}
}
