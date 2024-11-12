package main

import (
	"fmt"
)

type FileMeta struct {
	Owner string
}

var permissions = make(map[string]FileMeta)

func setFileOwner(filePath string, owner string) {
	permissions[filePath] = FileMeta{Owner: owner}
}

func checkPermissions(user *User, filePath string) bool {
	fileMeta, exists := permissions[filePath]
	if !exists {
		fmt.Println("Arquivo ou diretório não encontrado.")
		return false
	}
	return fileMeta.Owner == user.Username
}
