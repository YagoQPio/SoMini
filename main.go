package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func readPassword() (string, error) {
    // Função para ler a senha do usuário
    fmt.Print("Digite a senha: ")
    password, err := bufio.NewReader(os.Stdin).ReadString('\n')
    if err != nil {
       return "", err
    }
    return strings.TrimSpace(password), nil
}

func main() {
    for {
       fmt.Println("\nBem-vindo ao MiniSO")
       fmt.Println("1. Cadastre-se")
       fmt.Println("2. Login")
       fmt.Println("3. Sair do SO")
       fmt.Print("Escolha uma opção: ")

       var choice int
       fmt.Scanln(&choice)

       switch choice {
       case 1:
          registerUser()
       case 2:
          user, err := loginUser()
          if err != nil {
             fmt.Println("Erro no login:", err)
          } else {
             fmt.Printf("Bem-vindo, %s!\n", user.Username)
             runShell(user)
          }
       case 3:
          fmt.Println("Saindo do SO...")
          return
       default:
          fmt.Println("Opção inválida.")
       }
    }
}

func registerUser() {
    var username, password string
    fmt.Print("Criação de usuário - Digite o nome: ")
    fmt.Scanln(&username)
    password, err := readPassword()
    if err != nil {
       fmt.Println("Erro ao ler a senha:", err)
       return
    }

    user, err := createUser(username, password)
    if err != nil {
       fmt.Println("Erro ao criar o usuário:", err)
       return
    }

    err = saveUser(*user)
    if err != nil {
       fmt.Println("Erro ao salvar o usuário:", err)
    } else {
       fmt.Println("Usuário cadastrado com sucesso!")
    }
}

func loginUser() (*User, error) {
    users, err := loadUsers()
    if err != nil {
       fmt.Println("Erro ao carregar usuários:", err)
       os.Exit(1)
    }

    if len(users) == 0 {
       fmt.Println("Nenhum usuário cadastrado. Por favor, cadastre-se primeiro.")
       registerUser()
       return nil, fmt.Errorf("cadastro realizado, por favor faça o login")
    }

    user, err := login(users)
    if err != nil {
       return nil, fmt.Errorf("usuário ou senha inválidos")
    }

    return user, nil
}

func runShell(currentUser *User) {
    reader := bufio.NewReader(os.Stdin)

    for {
       fmt.Print("> ")
       command, err := reader.ReadString('\n')
       if err != nil {
          fmt.Println("Erro ao ler comando:", err)
          continue
       }
       command = strings.TrimSpace(command)

       if command == "sair" {
          fmt.Println("Saindo da conta...")
          return
       }

       executeCommand(command, currentUser)
    }
}

func executeCommand(command string, currentUser *User) {
    parts := strings.Fields(command)

    if len(parts) == 0 {
       fmt.Println("Comando vazio.")
       return
    }

    switch parts[0] {
    case "listar":
       listFiles(currentUser)
    case "criar":
       if len(parts) < 3 {
          fmt.Println("Uso: criar <arquivo/diretorio> <nome>")
          return
       }
       if parts[1] == "arquivo" {
          createFile(parts[2], currentUser)
       } else if parts[1] == "diretorio" {
          createDirectory(parts[2], currentUser)
       } else {
          fmt.Println("Comando inválido. Use 'arquivo' ou 'diretorio'.")
       }
    case "apagar":
       if len(parts) < 3 {
          fmt.Println("Uso: apagar <arquivo/diretorio> <nome>")
          return
       }
       if parts[1] == "arquivo" {
          deleteFile(parts[2], currentUser)
       } else if parts[1] == "diretorio" {
          deleteDirectory(parts[2], currentUser)
       } else {
          fmt.Println("Comando inválido. Use 'arquivo' ou 'diretorio'.")
       }
    default:
       fmt.Println("Comando desconhecido:", parts[0])
    }
}
