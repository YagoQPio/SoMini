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
	if len(parts) < 3 || parts[1] != "arquivo" {
		return fmt.Errorf("uso: criar/apagar arquivo <caminho>")
	}

	switch parts[0] {
	case "criar":
		caminho := parts[2]
		if strings.HasSuffix(caminho, "/") {
			// Criar diretório com permissões
			err := CriarDiretorio(caminho, usuario)
			if err != nil {
				return err
			}
			// Registrar permissão para o usuário
			return RegistrarPermissao(caminho, usuario)
		}
		// Criar arquivo com permissões
		err := CriarArquivo(caminho, usuario)
		if err != nil {
			return err
		}
		// Registrar permissão para o usuário
		return RegistrarPermissao(caminho, usuario)

	case "apagar":
		caminho := parts[2]
		force := len(parts) > 3 && parts[3] == "--force"

		// Verificar permissões antes de apagar
		ehProprietario, err := VerificarPermissao(caminho, usuario)
		if err != nil {
			return fmt.Errorf("erro ao verificar permissões: %v", err)
		}
		if !ehProprietario {
			return fmt.Errorf("usuário %s não tem permissão para apagar %s", usuario, caminho)
		}

		// Apagar diretório ou arquivo
		if strings.HasSuffix(caminho, "/") {
			return ApagarDiretorio(caminho, usuario, force)
		}
		return ApagarArquivo(caminho, usuario)

	case "listar":
		caminho := parts[2]

		// Verificar permissões antes de listar
		_, err := VerificarPermissao(caminho, usuario)
		if err != nil {
			return fmt.Errorf("erro ao verificar permissões: %v", err)
		}

		// Listar conteúdo de diretório
		return ListarDiretorio(caminho)

	default:
		// Caso não seja um comando conhecido, tenta executar como comando externo
		cmd := exec.Command(parts[0], parts[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		if err := cmd.Start(); err != nil {
			return fmt.Errorf("erro ao iniciar o processo: %v", err)
		}
		fmt.Printf("Processo iniciado com PID: %d\n", cmd.Process.Pid)

		return cmd.Wait()
	}
}
