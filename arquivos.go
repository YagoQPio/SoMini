package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Função para criar arquivos
func CriarArquivo(caminho string, usuario string) error {
	// Extrai o diretório pai do arquivo
	dir := filepath.Dir(caminho)

	// Verifica se o diretório pai não é o diretório atual
	if dir != "." {
		// Checa se o diretório pai existe
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			// Cria o diretório pai se ele não existir
			err := CriarDiretorio(dir, usuario)
			if err != nil {
				return fmt.Errorf("erro ao criar diretório pai para o arquivo: %v", err)
			}
		}
	}

	// Cria o arquivo
	file, err := os.Create(caminho)
	if err != nil {
		return fmt.Errorf("erro ao criar o arquivo: %v", err)
	}
	defer file.Close()

	// Escreve informações iniciais no arquivo
	_, err = file.WriteString(fmt.Sprintf("Criado pelo usuário %s.\n", usuario))
	if err != nil {
		return fmt.Errorf("erro ao escrever no arquivo: %v", err)
	}

	// Registra a permissão do arquivo
	err = RegistrarPermissao(caminho, usuario)
	if err != nil {
		return fmt.Errorf("erro ao registrar permissão para o arquivo: %v", err)
	}

	fmt.Printf("[SUCESSO] Arquivo criado com sucesso: %s\n", caminho)
	return nil
}

// Função para apagar arquivos
func ApagarArquivo(caminho, usuario string) error {
	// Verifica permissão
	permitido, err := VerificarPermissao(caminho, usuario)
	if err != nil {
		return fmt.Errorf("[ERRO] Verificação de permissão falhou: %v", err)
	}
	if !permitido {
		return fmt.Errorf("[ERRO] Você não tem permissão para apagar este arquivo.")
	}

	// Remove o arquivo
	err = os.Remove(caminho)
	if err != nil {
		return fmt.Errorf("erro ao apagar o arquivo: %v", err)
	}

	// Remove permissões associadas
	err = RemoverPermissao(caminho)
	if err != nil {
		return fmt.Errorf("[ERRO] Falha ao remover permissão: %v", err)
	}

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
func CriarDiretorio(caminho, usuario string) error {
	// Lógica para criar o diretório
	err := os.Mkdir(caminho, 0755)
	if err != nil {
		return fmt.Errorf("falha ao criar o diretório: %v", err)
	}

	// Registra o proprietário do diretório
	err = RegistrarPermissao(caminho, usuario)
	if err != nil {
		return fmt.Errorf("falha ao registrar permissão: %v", err)
	}

	return nil
}

// Função para apagar diretórios
func ApagarDiretorio(caminho string, usuario string, force bool) error {
	// Verifica se o diretório existe
	if _, err := os.Stat(caminho); os.IsNotExist(err) {
		return fmt.Errorf("o diretório não existe: %s", caminho)
	}

	// Caso force não seja usado, verifica a permissão do usuário
	if !force {
		permitido, err := VerificarPermissao(caminho, usuario)
		if err != nil {
			return fmt.Errorf("[ERRO] Falha ao verificar permissão: %v", err)
		}
		if !permitido {
			return fmt.Errorf("[ERRO] Você não tem permissão para apagar este diretório. Use --force para forçar a remoção.")
		}

		// Verifica se o diretório está vazio
		entries, err := os.ReadDir(caminho)
		if err != nil {
			return fmt.Errorf("erro ao verificar conteúdo do diretório: %v", err)
		}
		if len(entries) > 0 {
			return fmt.Errorf("[ERRO] Diretório não está vazio. Para remover, use: apagar diretorio <caminho> --force")
		}
	}

	// Caso force seja usado, apaga tudo, mesmo que o diretório tenha subitens
	if force {
		err := filepath.Walk(caminho, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				permitido, err := VerificarPermissao(path, usuario)
				if err == nil && permitido {
					// Remove permissões do arquivo
					if remErr := RemoverPermissao(path); remErr != nil {
						fmt.Printf("[ERRO] Falha ao remover permissão do arquivo: %v\n", remErr)
					}
				}
				// Remove o arquivo
				return os.Remove(path)
			}
			return nil
		})
		if err != nil {
			return fmt.Errorf("erro ao apagar itens do diretório: %v", err)
		}
	}

	// Finalmente, apaga o diretório raiz
	err := os.RemoveAll(caminho)
	if err != nil {
		return fmt.Errorf("erro ao apagar diretório: %v", err)
	}

	// Remove permissões associadas ao diretório e subcaminhos
	err = RemoverPermissoesRecursivas(caminho)
	if err != nil {
		return fmt.Errorf("[ERRO] Falha ao remover permissões associadas: %v", err)
	}

	fmt.Printf("[SUCESSO] Diretório apagado por %s: %s\n", usuario, caminho)
	return nil
}
