package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Processa comandos do shell
func processarComando(comando, usuario string) error {
	parts := strings.Fields(comando)
	if len(parts) < 1 {
		return fmt.Errorf("comando inválido. Use: criar/apagar/listar/adicionar_usuario <caminho>")
	}

	switch parts[0] {
	case "adicionar_usuario":
		return CriarUsuario()
	case "criar":
		if len(parts) < 2 {
			return fmt.Errorf("uso: criar <caminho_do_arquivo>")
		}
		return CriarArquivo(parts[1], usuario)
	case "apagar":
		if len(parts) < 2 {
			return fmt.Errorf("uso: apagar <caminho>")
		}
		if strings.HasPrefix(parts[1], "diretorio") {
			force := len(parts) > 2 && parts[2] == "--force"
			return ApagarDiretorio(parts[1], usuario, force)
		}
		return ApagarArquivo(parts[1], usuario)
	case "listar":
		caminho := "."
		if len(parts) >= 2 {
			caminho = parts[1]
		}
		return ListarDiretorio(caminho)
	case "criar_dir":
		if len(parts) < 2 {
			return fmt.Errorf("uso: criar_dir <caminho_do_diretorio>")
		}
		return CriarDiretorio(parts[1], usuario)
	default:
		cmd := exec.Command(parts[0], parts[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}
}
