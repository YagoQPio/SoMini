package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Arquivo struct {
	Path    string `json:"path"`
	Criador string `json:"criador"`
}

const arquivoPermissoes = "permissoes.json"

// Carrega as permissões do arquivo
func CarregarPermissoes() ([]Arquivo, error) {
	var arquivos []Arquivo
	file, err := os.Open(arquivoPermissoes)
	if err != nil {
		if os.IsNotExist(err) {
			return []Arquivo{}, nil
		}
		return nil, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&arquivos)
	return arquivos, err
}

// Salva as permissões no arquivo
func SalvarPermissoes(arquivos []Arquivo) error {
	file, err := os.Create(arquivoPermissoes)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(arquivos)
}

// Adiciona um arquivo/diretório às permissões
func AdicionarPermissao(caminho string, usuario string) error {
	arquivos, err := CarregarPermissoes()
	if err != nil {
		return err
	}

	for _, a := range arquivos {
		if a.Path == caminho {
			return fmt.Errorf("o arquivo/diretório já existe nas permissões")
		}
	}

	arquivos = append(arquivos, Arquivo{Path: caminho, Criador: usuario})
	return SalvarPermissoes(arquivos)
}

// Verifica permissões de acesso
func VerificarPermissao(caminho string, usuario string) error {
	arquivos, err := CarregarPermissoes()
	if err != nil {
		return err
	}

	for _, a := range arquivos {
		if a.Path == caminho {
			if a.Criador != usuario {
				return fmt.Errorf("você não tem permissão para acessar: %s", caminho)
			}
			return nil
		}
	}
	return fmt.Errorf("o arquivo/diretório não existe")
}

// Função para criar arquivos
func CriarArquivo(caminho string, usuario string) error {
	dir := filepath.Dir(caminho)

	// Cria os diretórios necessários
	if dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("erro ao criar diretórios: %v", err)
		}
	}

	// Verifica permissões no arquivo (se já existe)
	if _, err := os.Stat(caminho); err == nil {
		// O arquivo já existe, verifica permissões
		if err := VerificarPermissao(caminho, usuario); err != nil {
			return err
		}
	}

	// Cria o arquivo
	file, err := os.Create(caminho)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo: %v", err)
	}
	defer file.Close()

	// Escreve conteúdo inicial no arquivo
	_, err = file.WriteString(fmt.Sprintf("Criado pelo usuário %s.\n", usuario))
	if err != nil {
		return fmt.Errorf("erro ao escrever no arquivo: %v", err)
	}

	// Adiciona permissões para o novo arquivo
	err = AdicionarPermissao(caminho, usuario)
	if err != nil {
		return fmt.Errorf("erro ao adicionar permissões: %v", err)
	}

	fmt.Printf("Arquivo criado com sucesso: %s\n", caminho)
	return nil
}

// Função para apagar arquivos
func ApagarArquivo(caminho string, usuario string) error {
	err := VerificarPermissao(caminho, usuario)
	if err != nil {
		return err
	}

	err = os.Remove(caminho)
	if err != nil {
		return fmt.Errorf("erro ao apagar arquivo: %v", err)
	}

	fmt.Printf("Arquivo apagado com sucesso por %s: %s\n", usuario, caminho)
	return nil
}

// Função para listar diretórios
func ListarDiretorio(caminho string) error {
	files, err := os.ReadDir(caminho)
	if err != nil {
		return fmt.Errorf("erro ao listar diretório: %v", err)
	}

	fmt.Println("Conteúdo do diretório:")
	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("[DIR] %s\n", file.Name())
		} else {
			fmt.Printf("[FILE] %s\n", file.Name())
		}
	}
	return nil
}

// Função para criar diretórios
func CriarDiretorio(caminho string, usuario string) error {
	err := os.MkdirAll(caminho, 0755)
	if err != nil {
		return fmt.Errorf("erro ao criar diretório: %v", err)
	}

	err = AdicionarPermissao(caminho, usuario)
	if err != nil {
		return fmt.Errorf("erro ao adicionar permissões ao diretório: %v", err)
	}

	fmt.Printf("Diretório criado com sucesso: %s\n", caminho)
	return nil
}

// Função para apagar diretórios
func ApagarDiretorio(caminho string, usuario string, force bool) error {
	err := VerificarPermissao(caminho, usuario)
	if err != nil {
		return err
	}

	if force {
		err = os.RemoveAll(caminho)
	} else {
		err = os.Remove(caminho)
	}

	if err != nil {
		return fmt.Errorf("erro ao apagar diretório: %v", err)
	}

	arquivos, err := CarregarPermissoes()
	if err != nil {
		return err
	}

	var arquivosAtualizados []Arquivo
	for _, a := range arquivos {
		if !strings.HasPrefix(a.Path, caminho) {
			arquivosAtualizados = append(arquivosAtualizados, a)
		}
	}

	err = SalvarPermissoes(arquivosAtualizados)
	if err != nil {
		return fmt.Errorf("erro ao atualizar permissões: %v", err)
	}

	fmt.Printf("Diretório apagado com sucesso por %s: %s\n", usuario, caminho)
	return nil
}
