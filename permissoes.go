package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const arquivoPermissoes = "permissoes.json"

type Permissao struct {
	Path     string `json:"path"`
	Username string `json:"username"`
}

// Carrega permissões do arquivo
func CarregarPermissoes() (map[string]string, error) {
	permissoes := make(map[string]string)

	file, err := os.Open(arquivoPermissoes)
	if err != nil {
		if os.IsNotExist(err) {
			return permissoes, nil
		}
		return nil, fmt.Errorf("erro ao carregar permissões: %v", err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&permissoes)
	if err != nil {
		return nil, fmt.Errorf("erro ao decodificar permissões: %v", err)
	}

	return permissoes, nil
}

// Salva permissões no arquivo
func SalvarPermissoes(permissoes map[string]string) error {
	file, err := os.Create(arquivoPermissoes)
	if err != nil {
		return fmt.Errorf("erro ao salvar permissões: %v", err)
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(permissoes)
}

// Verifica se o usuário tem permissão para acessar o caminho
func VerificarPermissao(caminho, usuario string) (bool, error) {
	permissoes, err := CarregarPermissoes()
	if err != nil {
		return false, err // Retorna false em caso de erro ao carregar as permissões
	}

	// Verifica se o caminho existe nas permissões
	proprietario, existe := permissoes[caminho]
	if !existe {
		return false, fmt.Errorf("o caminho %s não possui um proprietário registrado", caminho)
	}

	// Verifica se o usuário é o proprietário
	if proprietario != usuario {
		return false, fmt.Errorf("Acesso negado, você não tem permissão para apagar esse arquivo.")
	}

	return true, nil // Retorna true se o usuário tem permissão
}

// Registra a permissão para um usuário em um caminho
func RegistrarPermissao(caminho, usuario string) error {
	permissoes, err := CarregarPermissoes()
	if err != nil {
		return err
	}

	permissoes[caminho] = usuario
	return SalvarPermissoes(permissoes)
}
