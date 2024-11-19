package main

import (
    "fmt"
    "os"
)

// Variável global que armazena o usuário logado atualmente
var usuarioAtual string

func main() {
    for {
        fmt.Println("Mini SO - Menu Principal")
        fmt.Println("1. Criar Usuário")
        fmt.Println("2. Login")
        fmt.Println("3. Fechar Programa")
        fmt.Print("Escolha uma opção: ")

        var opcao int
        fmt.Scan(&opcao)

        switch opcao {
        case 1:
            CriarUsuario()
        case 2:
            if Login() {
                MenuUsuario()
            }
        case 3:
            fmt.Println("Encerrando o programa...")
            os.Exit(0)
        default:
            fmt.Println("Opção inválida!")
        }
    }
}
