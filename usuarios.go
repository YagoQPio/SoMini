package main

import (
	"bufio"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

const userFile = "users.txt"

// CriarUsuario cria um novo usuário e salva no arquivo.
func CriarUsuario() {
	fmt.Println("\n== Criar Novo Usuário ==")

	// Captura o nome de usuário
	fmt.Print("Digite o nome de usuário: ")
	var username string
	fmt.Scanln(&username)

	// Captura a senha
	fmt.Print("Digite a senha: ")
	var senha string
	fmt.Scanln(&senha)

	// Aplica o hash na senha
	senhaHashed := gerarHashSHA512(senha)

	// Salva o usuário no arquivo
	if salvarUsuario(username, senhaHashed) {
		fmt.Println("Usuário criado com sucesso!")
	} else {
		fmt.Println("Erro ao criar usuário.")
	}
}

// Login autentica o usuário com base no arquivo.
func Login() bool {
	fmt.Println("\n== Logar ==")

	// Captura o nome de usuário
	fmt.Print("Digite o nome de usuário: ")
	var username string
	fmt.Scanln(&username)

	// Captura a senha
	fmt.Print("Digite a senha: ")
	var senha string
	fmt.Scanln(&senha)

	// Aplica o hash na senha para autenticar
	senhaHashed := gerarHashSHA512(senha)

	// Autentica o usuário
	if autenticarUsuario(username, senhaHashed) {
		fmt.Println("Login realizado com sucesso!")
		return true
	}

	fmt.Println("Usuário ou senha incorretos.")
	return false
}

// gerarHashSHA512 gera o hash da senha usando SHA-512.
func gerarHashSHA512(senha string) string {
	hash := sha512.New()
	hash.Write([]byte(senha))
	return hex.EncodeToString(hash.Sum(nil))
}

// salvarUsuario salva os dados do usuário no arquivo.
func salvarUsuario(username, senhaHashed string) bool {
	file, err := os.OpenFile(userFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return false
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%s:%s\n", username, senhaHashed))
	return err == nil
}

// autenticarUsuario verifica se as credenciais estão no arquivo.
func autenticarUsuario(username, senhaHashed string) bool {
	file, err := os.Open(userFile)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		if len(parts) == 2 && parts[0] == username && parts[1] == senhaHashed {
			return true
		}
	}

	return false
}
