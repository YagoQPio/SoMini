package main

import (
	"fmt"
	"os"
	"strings"
)

func salvarMetaDados(caminho, dono string) {
	file, _ := os.OpenFile(caminho+".meta", os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	file.WriteString(dono)
}

func verificarPermissao(caminho, usuario string) bool {
	metaFile := caminho + ".meta"
	data, err := os.ReadFile(metaFile)
	if err != nil {
		fmt.Printf("Permissão negada para '%s': %v\n", caminho, err)
		return false
	}
	dono := strings.TrimSpace(string(data))
	return dono == usuario
}
