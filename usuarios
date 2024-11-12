package main

import (
    "crypto/sha512"
    "encoding/hex"
    "fmt"
    "os"
)

type User struct {
    Username string
    Password string
}

func createUser(username, password string) (*User, error) {
    hashedPassword := hashPassword(password)
    user := &User{
       Username: username,
       Password: hashedPassword,
    }
    return user, nil
}

func hashPassword(password string) string {
    hash := sha512.New()
    hash.Write([]byte(password))
    return hex.EncodeToString(hash.Sum(nil))
}

func saveUser(user User) error {
    // Simulando a gravação do usuário no arquivo (somente o nome de usuário)
    file, err := os.OpenFile("users.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
       return err
    }
    defer file.Close()

    _, err = file.WriteString(fmt.Sprintf("%s %s\n", user.Username, user.Password)) // Salvando o nome de usuário e senha
    return err
}

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
