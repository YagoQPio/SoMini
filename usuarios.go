package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
)

// Estrutura do Usuário
type User struct {
	Username string
	Password string
}

// Função para criar um novo usuário (com senha criptografada)
func createUser(username, password string) (*User, error) {
	hashedPassword := hashPassword(password)
	user := &User{
		Username: username,
		Password: hashedPassword,
	}
	return user, nil
}

// Função para gerar o hash da senha
func hashPassword(password string) string {
	hash := sha512.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

// Função para salvar um usuário no arquivo
func saveUser(user User) error {
	file, err := os.OpenFile("users.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%s %s\n", user.Username, user.Password))
	return err
}

// Função para carregar os usuários a partir do arquivo
func loadUsers() ([]User, error) {
	var users []User
	file, err := os.Open("users.json")
	if err != nil {
		return users, err
	}
	defer file.Close()

	var username, password string
	for {
		_, err := fmt.Fscanf(file, "%s %s\n", &username, &password)
		if err != nil {
			break
		}
		users = append(users, User{Username: username, Password: password})
	}
	return users, nil
}

// Função para registrar um novo usuário
func registerUser() (*User, error) {
	var username, password string
	fmt.Print("Digite o nome de usuário: ")
	fmt.Scanln(&username)
	fmt.Print("Digite a senha: ")
	password, err := readPassword()
	if err != nil {
		return nil, fmt.Errorf("erro ao ler a senha")
	}

	// Criação e salvamento do usuário
	user, err := createUser(username, password)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar o usuário")
	}

	err = saveUser(*user)
	if err != nil {
		return nil, fmt.Errorf("erro ao salvar o usuário")
	}

	return user, nil
}

// Função para fazer login
func login(users []User) (*User, error) {
	var username, password string
	fmt.Print("Digite o nome de usuário: ")
	fmt.Scanln(&username)
	fmt.Print("Digite a senha: ")
	password, err := readPassword()
	if err != nil {
		return nil, fmt.Errorf("erro ao ler a senha")
	}

	hashedPassword := hashPassword(password)

	for _, user := range users {
		if user.Username == username && user.Password == hashedPassword {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("usuário ou senha inválidos")
}

// Função para ler a senha (sem exibir no terminal)
func readPassword() (string, error) {
	var password string
	fmt.Scanln(&password)
	return password, nil
}
