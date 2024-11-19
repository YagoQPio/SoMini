package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// CriarArquivo cria um arquivo ou diretório associado ao usuário logado.
func CriarArquivo(caminho string) {
	if usuarioAtual == "" {
		fmt.Println("Erro: Nenhum usuário logado.")
		return
	}

	// Cria o diretório do caminho, se necessário
	err := os.MkdirAll(filepath.Dir(caminho), os.ModePerm)
	if err != nil {
		fmt.Println("Erro ao criar diretório:", err)
		return
	}

	// Cria o arquivo
	file, err := os.Create(caminho)
	if err != nil {
		fmt.Println("Erro ao criar arquivo:", err)
		return
	}
	defer file.Close()

	// Adiciona o dono ao arquivo
	file.WriteString(fmt.Sprintf("Dono: %s\n", usuarioAtual))

	fmt.Println("Arquivo criado com sucesso:", caminho)
}

// ApagarArquivo apaga um arquivo ou diretório se o usuário for o dono.
func ApagarArquivo(caminho string) {
	if usuarioAtual == "" {
		fmt.Println("Erro: Nenhum usuário logado.")
		return
	}

	// Verificar se o usuário tem permissão
	dono := verificarDono(caminho)
	if dono != usuarioAtual {
		fmt.Println("Erro: Você não tem permissão para apagar este arquivo.")
		return
	}

	err := os.RemoveAll(caminho)
	if err != nil {
		fmt.Println("Erro ao apagar arquivo/diretório:", err)
		return
	}

	fmt.Println("Arquivo/diretório apagado com sucesso:", caminho)
}

// ApagarDiretorioForce apaga um diretório com a flag --force.
func ApagarDiretorioForce(comando string) {
	if usuarioAtual == "" {
		fmt.Println("Erro: Nenhum usuário logado.")
		return
	}

	// Divide o comando em partes para obter o caminho
	partes := strings.Split(comando, " ")
	if len(partes) < 3 || partes[0] != "apagar" || partes[1] != "diretorio" || partes[len(partes)-1] != "--force" {
		fmt.Println("Comando inválido. Use o formato: apagar diretorio <caminho> --force")
		return
	}

	caminho := strings.Join(partes[2:len(partes)-1], " ") // Extrai o caminho
	err := os.RemoveAll(caminho)                          // Apaga o diretório

	if err != nil {
		fmt.Println("Erro ao apagar o diretório:", err)
		return
	}

	fmt.Printf("Diretório '%s' apagado com sucesso.\n", caminho)
}

// ListarTodosArquivos lista todos os arquivos criados pelo programa com seus respectivos donos.
func ListarTodosArquivos() {
	// Obtém o diretório atual como raiz
	raiz, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter o diretório atual:", err)
		return
	}

	fmt.Println("Listando todos os arquivos e diretórios criados pelos usuários:")

	err = filepath.Walk(raiz, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Verifica se é um arquivo e possui um dono registrado
		if !info.IsDir() {
			dono := verificarDono(path)
			if dono != "Desconhecido" {
				fmt.Printf("Arquivo: %s, Dono: %s\n", path, dono)
			}
		}

		return nil
	})
	if err != nil {
		fmt.Println("Erro ao listar arquivos:", err)
	}
}

// verificarDono verifica quem é o dono de um arquivo.
func verificarDono(caminho string) string {
	file, err := os.Open(caminho)
	if err != nil {
		return "Desconhecido"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Dono: ") {
			return strings.TrimPrefix(line, "Dono: ")
		}
	}

	return "Desconhecido"
}
