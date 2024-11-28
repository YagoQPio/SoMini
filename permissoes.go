package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
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
	// Carrega as permissões do sistema
	permissoes, err := CarregarPermissoes()
	if err != nil {
		return false, fmt.Errorf("[ERRO] Falha ao carregar permissões: %v", err)
	}

	// Verifica se o caminho possui permissões registradas
	proprietario, existe := permissoes[caminho]
	if !existe {
		return false, fmt.Errorf("[ERRO] O caminho %s não possui um proprietário registrado.", caminho)
	}

	// Verifica se o usuário atual é o proprietário
	if proprietario != usuario {
		return false, fmt.Errorf("Acesso negado: você não tem permissão para apagar o caminho %s.", caminho)
	}

	// Retorna sucesso se o usuário for o proprietário
	return true, nil
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

func RemoverPermissao(caminho string) error {
	permissoes, err := CarregarPermissoes()
	if err != nil {
		return fmt.Errorf("falha ao carregar permissões: %v", err)
	}

	if _, existe := permissoes[caminho]; existe {
		delete(permissoes, caminho)
		err = SalvarPermissoes(permissoes)
		if err != nil {
			return fmt.Errorf("falha ao atualizar permissões: %v", err)
		}
	}
	return nil
}

func RemoverPermissoesRecursivas(caminho string) error {
	permissoes, err := CarregarPermissoes()
	if err != nil {
		return fmt.Errorf("falha ao carregar permissões: %v", err)
	}

	// Remove todas as permissões de subcaminhos
	for path := range permissoes {
		if strings.HasPrefix(path, caminho) {
			delete(permissoes, path)
		}
	}

	// Salva as permissões atualizadas
	err = SalvarPermissoes(permissoes)
	if err != nil {
		return fmt.Errorf("falha ao salvar permissões: %v", err)
	}

	return nil
}
