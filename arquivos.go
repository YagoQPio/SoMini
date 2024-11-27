package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Função para criar arquivos
func CriarArquivo(caminho string, usuario string) error {
	dir := filepath.Dir(caminho)
	if dir != "." {
		// Cria o diretório se necessário
		if err := CriarDiretorio(dir, usuario); err != nil {
			return fmt.Errorf("erro ao criar diretórios para o arquivo: %v", err)
		}
	}

	// Cria o arquivo
	file, err := os.Create(caminho)
	if err != nil {
		return err
	}
	defer file.Close()

	// Escreve informações sobre o arquivo
	_, err = file.WriteString(fmt.Sprintf("Criado pelo usuário %s.\n", usuario))
	if err != nil {
		return fmt.Errorf("erro ao escrever no arquivo: %v", err)
	}

	fmt.Printf("Arquivo criado com sucesso: %s\n", caminho)
	return nil
}

// Função para apagar arquivos
func ApagarArquivo(caminho string, usuario string) error {
	if _, err := os.Stat(caminho); os.IsNotExist(err) {
		return fmt.Errorf("o arquivo/diretório não existe: %s", caminho)
	}

	err := os.Remove(caminho)
	if err != nil {
		return fmt.Errorf("erro ao apagar arquivo: %v", err)
	}

	fmt.Printf("Arquivo apagado com sucesso por %s: %s\n", usuario, caminho)
	return nil
}

// Função para listar diretórios
func ListarDiretorio(caminho string) error {
	files, err := os.ReadDir(caminho)
	if err != nil {
		return fmt.Errorf("erro ao listar diretório: %v", err)
	}

	fmt.Println("Conteúdo do diretório:")
	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("[DIR] %s\n", file.Name())
		} else {
			fmt.Printf("[FILE] %s\n", file.Name())
		}
	}
	return nil
}

// Função para criar diretórios
func CriarDiretorio(caminho string, usuario string) error {
	if err := os.MkdirAll(caminho, 0755); err != nil {
		return fmt.Errorf("erro ao criar diretório: %v", err)
	}

	fmt.Printf("Diretório criado com sucesso por %s: %s\n", usuario, caminho)
	return nil
}

// Função para apagar diretórios
func ApagarDiretorio(caminho string, usuario string, force bool) error {
	if _, err := os.Stat(caminho); os.IsNotExist(err) {
		return fmt.Errorf("o diretório não existe: %s", caminho)
	}

	if force {
		err := os.RemoveAll(caminho)
		if err != nil {
			return fmt.Errorf("erro ao apagar diretório: %v", err)
		}
		fmt.Printf("Diretório apagado com sucesso por %s: %s\n", usuario, caminho)
	} else {
		// Verifica se está vazio antes de apagar
		entries, err := os.ReadDir(caminho)
		if err != nil {
			return fmt.Errorf("erro ao verificar conteúdo do diretório: %v", err)
		}
		if len(entries) > 0 {
			return fmt.Errorf("o diretório não está vazio. Use --force para apagar.")
		}

		err = os.Remove(caminho)
		if err != nil {
			return fmt.Errorf("erro ao apagar diretório: %v", err)
		}
		fmt.Printf("Diretório apagado com sucesso por %s: %s\n", usuario, caminho)
	}

	return nil
}
