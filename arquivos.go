package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// CriarArquivo ou diretório.
func CriarArquivo(caminho string) {
	err := os.MkdirAll(filepath.Dir(caminho), os.ModePerm)
	if err != nil {
		fmt.Println("Erro ao criar diretório:", err)
		return
	}

	file, err := os.Create(caminho)
	if err != nil {
		fmt.Println("Erro ao criar arquivo:", err)
		return
	}
	defer file.Close()

	fmt.Println("Arquivo criado com sucesso:", caminho)
}

// ApagarArquivo apaga um arquivo ou diretório.
func ApagarArquivo(caminho string) {
	err := os.RemoveAll(caminho)
	if err != nil {
		fmt.Println("Erro ao apagar arquivo/diretório:", err)
		return
	}

	fmt.Println("Arquivo/diretório apagado com sucesso:", caminho)
}

// ListarArquivos lista todos os arquivos no diretório do usuário.
func ListarArquivos(diretorio string) {
	err := filepath.Walk(diretorio, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path)
		return nil
	})
	if err != nil {
		fmt.Println("Erro ao listar arquivos:", err)
	}
}
