package main

import (
    "fmt"
    "os"
)

// Função para apagar um diretório
func DeleteDirectory(path string) error {
    // Verificando se o diretório existe
    info, err := os.Stat(path)
    if os.IsNotExist(err) {
       return fmt.Errorf("diretório não encontrado")
    }
    // Verificando permissões de leitura e escrita
    if info.IsDir() {
       // Tentando excluir o diretório
       err := os.RemoveAll(path)
       if err != nil {
          return fmt.Errorf("erro ao apagar o diretório: %v", err)
       }
       return nil
    }
    return fmt.Errorf("o caminho não é um diretório")
}
