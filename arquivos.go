package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Função para criar um arquivo, criando diretórios no caminho se necessário
func createFile(user *User, filePath string) {
	// Cria diretórios no caminho, se não existirem
	dir := filepath.Dir(filePath) // Extrai o diretório do caminho do arquivo
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Println("Erro ao criar diretórios:", err)
		return
	}

	// Cria o arquivo
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Erro ao criar arquivo:", err)
		return
	}
	defer file.Close()

	fmt.Printf("Arquivo %s criado com sucesso!\n", filePath)
}

// Função para listar arquivos no diretório
func listFiles() {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Erro ao listar arquivos:", err)
		return
	}

	if len(files) == 0 {
		fmt.Println("Não há arquivos no diretório.")
	} else {
		fmt.Println("Arquivos no diretório atual:")
		for _, file := range files {
			fmt.Println(file.Name())
		}
	}
}

// Função para apagar um arquivo
func deleteFile(user *User, filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("Erro ao apagar arquivo:", err)
		return
	}

	fmt.Printf("Arquivo %s apagado com sucesso!\n", filePath)
}

// Função para criar diretórios (incluindo diretórios aninhados)
func createDir(user *User, dirPath string) {
	// Cria diretórios e subdiretórios se não existirem
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		fmt.Println("Erro ao criar diretórios:", err)
		return
	}

	fmt.Printf("Diretório %s criado com sucesso!\n", dirPath)
}
