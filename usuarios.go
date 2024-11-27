package main

import (
	"bufio"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

type Usuario struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}

const arquivoUsuarios = "usuarios.json"

// Verifica se há usuários cadastrados
func UsuariosExistem() bool {
	_, err := os.Stat(arquivoUsuarios)
	return !os.IsNotExist(err)
}

// Carrega os usuários do arquivo
func CarregarUsuarios() ([]Usuario, error) {
	var usuarios []Usuario
	file, err := os.Open(arquivoUsuarios)
	if err != nil {
		if os.IsNotExist(err) {
			return []Usuario{}, nil
		}
		return nil, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&usuarios)
	return usuarios, err
}

// Salva os usuários no arquivo
func SalvarUsuarios(usuarios []Usuario) error {
	file, err := os.Create(arquivoUsuarios)
	if err != nil {
		return fmt.Errorf("erro ao salvar usuários: %v", err)
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(usuarios)
}

// Gera um salt aleatório
func GerarSalt() (string, error) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// Gera o hash da senha com o salt
func GerarHashSenha(senha, salt string) string {
	hash := sha512.New()
	hash.Write([]byte(senha + salt))
	return hex.EncodeToString(hash.Sum(nil))
}

// Captura senha de forma segura
func CapturarSenha(mensagem string) (string, error) {
	fmt.Print(mensagem)
	if term.IsTerminal(0) {
		senhaBytes, err := term.ReadPassword(0)
		fmt.Println()
		if err != nil {
			return "", fmt.Errorf("erro ao capturar senha: %v", err)
		}
		return strings.TrimSpace(string(senhaBytes)), nil
	} else {
		reader := bufio.NewReader(os.Stdin)
		senha, err := reader.ReadString('\n')
		if err != nil {
			return "", fmt.Errorf("erro ao capturar senha: %v", err)
		}
		return strings.TrimSpace(senha), nil
	}
}

// Cria um novo usuário, mantendo os usuários existentes
func CriarUsuario() error {
	fmt.Println("Adicionando um novo usuário.")

	fmt.Print("Digite o nome de usuário: ")
	var username string
	fmt.Scanln(&username)
	username = strings.TrimSpace(username)

	if len(username) == 0 {
		return fmt.Errorf("o nome de usuário não pode ser vazio")
	}

	senha, err := CapturarSenha("Digite a senha: ")
	if err != nil {
		return err
	}
	if len(senha) == 0 {
		return fmt.Errorf("a senha não pode ser vazia")
	}

	// Gera salt e hash para a senha
	salt, err := GerarSalt()
	if err != nil {
		return fmt.Errorf("erro ao gerar salt: %v", err)
	}
	hashSenha := GerarHashSenha(senha, salt)

	// Carrega usuários existentes
	usuarios, err := CarregarUsuarios()
	if err != nil {
		return fmt.Errorf("erro ao carregar usuários existentes: %v", err)
	}

	// Verifica se o nome de usuário já existe
	for _, usuario := range usuarios {
		if usuario.Username == username {
			return fmt.Errorf("o nome de usuário '%s' já está em uso", username)
		}
	}

	// Adiciona o novo usuário à lista
	novoUsuario := Usuario{Username: username, Password: hashSenha, Salt: salt}
	usuarios = append(usuarios, novoUsuario)

	// Salva todos os usuários novamente no arquivo
	err = SalvarUsuarios(usuarios)
	if err != nil {
		return fmt.Errorf("erro ao salvar usuários no arquivo: %v", err)
	}

	return nil
}

// Login de usuário
func LoginUsuario() (*Usuario, error) {
	usuarios, err := CarregarUsuarios()
	if err != nil {
		return nil, err
	}

	fmt.Print("Digite o nome de usuário: ")
	var username string
	fmt.Scanln(&username)

	senha, err := CapturarSenha("Digite a senha: ")
	if err != nil {
		return nil, err
	}

	for _, usuario := range usuarios {
		if usuario.Username == username {
			hashSenha := GerarHashSenha(senha, usuario.Salt)
			if hashSenha == usuario.Password {
				return &usuario, nil
			}
			return nil, fmt.Errorf("senha incorreta")
		}
	}

	return nil, fmt.Errorf("usuário não encontrado")
}
