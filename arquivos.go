package main

import (
	"fmt"
	"os"
)

func criarArquivo(caminho string) {
	// Criação do arquivo
	file, err := os.Create(caminho)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer file.Close()

	// Escreve conteúdo aleatório no arquivo
	_, err = file.WriteString("Conteúdo aleatório gerado para o arquivo.")
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}

	fmt.Println("Arquivo criado com sucesso em:", caminho)
}

func criarDiretorio(caminho string) {
	// Criação do diretório
	err := os.MkdirAll(caminho, os.ModePerm)
	if err != nil {
		fmt.Println("Erro ao criar o diretório:", err)
		return
	}
	fmt.Println("Diretório criado com sucesso em:", caminho)
}

func apagarArquivo(caminho string) {
	// Remove o arquivo
	err := os.Remove(caminho)
	if err != nil {
		fmt.Println("Erro ao apagar o arquivo:", err)
		return
	}
	fmt.Println("Arquivo apagado com sucesso:", caminho)
}

func apagarDiretorio(caminho string, force bool) {
	// Se for para apagar forçadamente (mesmo com arquivos dentro), usa RemoveAll
	if force {
		err := os.RemoveAll(caminho)
		if err != nil {
			fmt.Println("Erro ao apagar diretório forçadamente:", err)
			return
		}
		fmt.Println("Diretório apagado com sucesso (forçadamente):", caminho)
	} else {
		// Caso contrário, só apaga se estiver vazio
		err := os.Remove(caminho)
		if err != nil {
			fmt.Println("Erro ao tentar apagar diretório:", err)
			return
		}
		fmt.Println("Diretório apagado com sucesso:", caminho)
	}
}
