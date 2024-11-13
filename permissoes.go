package main

import (
	"fmt"
)

type FileMeta struct {
	Owner string
}

// Armazena permissões para cada arquivo/diretório
var permissions = make(map[string]FileMeta)

// Define o proprietário de um arquivo
func setFileOwner(filePath string, owner string) {
	permissions[filePath] = FileMeta{Owner: owner}
}

// Verifica se o usuário tem permissão para acessar um arquivo
func checkPermissions(usuario *Usuario, filePath string) bool {
	fileMeta, exists := permissions[filePath]
	if !exists {
		fmt.Println("Arquivo ou diretório não encontrado.")
		return false
	}
	return fileMeta.Owner == usuario.Username // Alterado de usuario.Nome para usuario.Username
}
