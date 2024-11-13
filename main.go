package main

import (
	"fmt"
	"strings"
	"os"
	"path/filepath"
)

func main() {
	for {
		fmt.Println("\n==================== MENU ====================")
		fmt.Println("1. Criar usuário")
		fmt.Println("2. Logar")
		fmt.Println("3. Sair")
		fmt.Print("Escolha uma opção (1/2/3): ")

		var opcao int
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			fmt.Println("\n== Criar Novo Usuário ==")
			fmt.Print("Digite o nome de usuário: ")
			var username string
			fmt.Scanln(&username)

			fmt.Print("Digite a senha: ")
			var senha string
			fmt.Scanln(&senha)

			criarUsuario(username, senha)
			fmt.Println("Usuário criado com sucesso!")

		case 2:
			fmt.Println("\n== Logar ==")
			fmt.Print("Digite o nome de usuário: ")
			var username string
			fmt.Scanln(&username)

			fmt.Print("Digite a senha: ")
			var senha string
			fmt.Scanln(&senha)

			logarUsuario(username, senha)

		case 3:
			fmt.Println("Saindo...")
			return

		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}

		// Exibir o menu de comandos após logar
		if usuarioLogado.Username != "" {
			fmt.Println("\nBem-vindo,", usuarioLogado.Username)
			for {
				fmt.Print("\nComando (listar/criar/apagar/sair): ")
				var comando string
				fmt.Scanln(&comando)

				switch comando {
				case "listar":
					listarComandos()
				case "criar":
					fmt.Print("Digite 'arquivo' ou 'diretorio' para criar: ")
					var tipo string
					fmt.Scanln(&tipo)

					fmt.Print("Digite o nome do caminho: ")
					var caminho string
					fmt.Scanln(&caminho)

					// Garantir que o caminho seja absoluto e fora do diretório do projeto
					absPath, err := filepath.Abs(caminho)
					if err != nil {
						fmt.Println("Erro ao gerar caminho absoluto:", err)
						break
					}

					if tipo == "arquivo" {
						criarArquivo(absPath)
					} else if tipo == "diretorio" {
						criarDiretorio(absPath)
					} else {
						fmt.Println("Tipo inválido. Use 'arquivo' ou 'diretorio'.")
					}

				case "apagar":
					fmt.Print("Digite 'arquivo' ou 'diretorio' para apagar: ")
					var tipo string
					fmt.Scanln(&tipo)

					fmt.Print("Digite o nome do caminho: ")
					var caminho string
					fmt.Scanln(&caminho)

					// Garantir que o caminho seja absoluto e fora do diretório do projeto
					absPath, err := filepath.Abs(caminho)
					if err != nil {
						fmt.Println("Erro ao gerar caminho absoluto:", err)
						break
					}

					if tipo == "arquivo" {
						apagarArquivo(absPath)
					} else if tipo == "diretorio" {
						fmt.Print("Deseja forçar a remoção? (sim/nao): ")
						var forca string
						fmt.Scanln(&forca)
						forca = strings.ToLower(forca)
						apagarDiretorio(absPath, forca == "sim")
					} else {
						fmt.Println("Tipo inválido. Use 'arquivo' ou 'diretorio'.")
					}

				case "sair":
					usuarioLogado = Usuario{}
					fmt.Println("Você saiu da sua conta.")
					break

				default:
					fmt.Println("Comando desconhecido.")
				}
			}
		}
	}
}

func listarComandos() {
	fmt.Println("Comandos disponíveis:")
	fmt.Println("- listar <diretório> : Lista os arquivos de um diretório")
	fmt.Println("- criar arquivo <caminho> : Cria um arquivo no caminho especificado")
	fmt.Println("- criar diretorio <caminho> : Cria um diretório no caminho especificado")
	fmt.Println("- apagar arquivo <caminho> : Apaga o arquivo no caminho especificado")
	fmt.Println("- apagar diretorio <caminho> : Apaga o diretório no caminho especificado")
	fmt.Println("- sair : Sair da conta")
}
