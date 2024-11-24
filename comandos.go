package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func temUsuarios() bool {
	_, err := os.Stat("usuarios.txt")
	return !os.IsNotExist(err)
}

func criarUsuario() {
	var nome, senha string
	fmt.Print("Digite um nome de usuário: ")
	fmt.Scanln(&nome)
	fmt.Print("Digite uma senha: ")
	fmt.Scanln(&senha)
	hash := gerarHash(senha, nome)
	salvarUsuario(nome, hash)
	fmt.Println("Usuário criado com sucesso.")
}

func apagarUsuario() {
	var nome, senha string
	fmt.Print("Digite o nome do usuário a ser apagado: ")
	fmt.Scanln(&nome)
	fmt.Print("Digite a senha do usuário: ")
	fmt.Scanln(&senha)
	hash := gerarHash(senha, nome)

	if verificarUsuario(nome, hash) {
		if removerUsuario(nome) {
			fmt.Println("Usuário apagado com sucesso.")
			removerArquivosDoUsuario(nome)
			if !temUsuarios() {
				fmt.Println("Nenhum usuário restante. Será necessário criar um novo usuário.")
				criarUsuario()
			}
		} else {
			fmt.Println("Erro ao apagar o usuário.")
		}
	} else {
		fmt.Println("Usuário ou senha incorretos.")
	}
}

func removerUsuario(nome string) bool {
	data, err := os.ReadFile("usuarios.txt")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo de usuários:", err)
		return false
	}

	linhas := strings.Split(string(data), "\n")
	var novasLinhas []string
	for _, linha := range linhas {
		if !strings.HasPrefix(linha, nome+":") {
			novasLinhas = append(novasLinhas, linha)
		}
	}

	novoConteudo := strings.Join(novasLinhas, "\n")
	err = os.WriteFile("usuarios.txt", []byte(novoConteudo), 0644)
	if err != nil {
		fmt.Println("Erro ao atualizar o arquivo de usuários:", err)
		return false
	}

	return true
}

func removerArquivosDoUsuario(nome string) {
	diretorio := "./" + nome
	err := os.RemoveAll(diretorio)
	if err != nil {
		fmt.Printf("Erro ao remover arquivos e diretórios do usuário '%s': %v\n", nome, err)
	} else {
		fmt.Printf("Todos os arquivos e diretórios do usuário '%s' foram removidos.\n", nome)
	}
}

func login() bool {
	var nome, senha string
	fmt.Print("Usuário: ")
	fmt.Scanln(&nome)
	fmt.Print("Senha: ")
	fmt.Scanln(&senha)
	hash := gerarHash(senha, nome)
	if verificarUsuario(nome, hash) {
		fmt.Println("Login realizado com sucesso!")
		return true
	}
	fmt.Println("Usuário ou senha incorretos.")
	return false
}

func gerarHash(senha, salt string) string {
	h := sha512.New()
	h.Write([]byte(senha + salt))
	return hex.EncodeToString(h.Sum(nil))
}

func salvarUsuario(nome, hash string) {
	file, _ := os.OpenFile("usuarios.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	file.WriteString(fmt.Sprintf("%s:%s\n", nome, hash))
}

func verificarUsuario(nome, hash string) bool {
	data, err := os.ReadFile("usuarios.txt")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo de usuários:", err)
		return false
	}
	linhas := strings.Split(string(data), "\n")
	for _, linha := range linhas {
		if linha == fmt.Sprintf("%s:%s", nome, hash) {
			return true
		}
	}
	return false
}
