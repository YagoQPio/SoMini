package main

import (
	"fmt"
	"strings"
)

// Processa comandos do shell
func processarComando(entrada string, usuarioAtual string) error {
	// Divide a entrada do comando em partes
	partes := strings.Fields(entrada)
	if len(partes) < 2 {
		return fmt.Errorf("[ERRO] Comando inválido. Uso: criar diretorio <nomeDiretorio>")
	}

	comando := partes[0]     // Comando principal (ex: criar, apagar)
	subcomando := partes[1]  // Subcomando (ex: diretorio, arquivo)
	argumentos := partes[2:] // Argumentos adicionais (ex: nome do diretório ou arquivo)

	// Comando para criar um diretório
	if comando == "criar" && subcomando == "diretorio" {
		if len(argumentos) < 1 {
			return fmt.Errorf("[ERRO] Nome do diretório não fornecido. Uso: criar diretorio <nomeDiretorio>")
		}
		nomeDiretorio := argumentos[0]
		err := CriarDiretorio(nomeDiretorio, usuarioAtual)
		if err != nil {
			return fmt.Errorf("[ERRO] Falha ao criar diretório: %v", err)
		}
		fmt.Printf("[SUCESSO] Diretório criado: %s\n", nomeDiretorio)
		return nil
	}

	// Comando para criar um arquivo
	if comando == "criar" && subcomando == "arquivo" {
		if len(argumentos) < 1 {
			return fmt.Errorf("[ERRO] Nome do arquivo não fornecido. Uso: criar arquivo <nomeArquivo>")
		}
		nomeArquivo := argumentos[0]
		err := CriarArquivo(nomeArquivo, usuarioAtual)
		if err != nil {
			return fmt.Errorf("[ERRO] Falha ao criar arquivo: %v", err)
		}
		fmt.Printf("[SUCESSO] Arquivo criado: %s\n", nomeArquivo)
		return nil
	}

	// Comando para apagar um diretório
	if comando == "apagar" && subcomando == "diretorio" {
		if len(argumentos) < 1 {
			return fmt.Errorf("[ERRO] Caminho do diretório não fornecido. Uso: apagar diretorio <caminho> [--force]")
		}
		caminho := argumentos[0]
		force := len(argumentos) > 1 && argumentos[1] == "--force"
		err := ApagarDiretorio(caminho, usuarioAtual, force)
		if err != nil {
			return fmt.Errorf("[ERRO] Falha ao apagar diretório: %v", err)
		}
		fmt.Printf("[SUCESSO] Diretório apagado: %s\n", caminho)
		return nil
	}

	// Comando para apagar um arquivo
	if comando == "apagar" && subcomando == "arquivo" {
		if len(argumentos) < 1 {
			return fmt.Errorf("[ERRO] Caminho do arquivo não fornecido. Uso: apagar arquivo <caminho>")
		}
		caminho := argumentos[0]
		err := ApagarArquivo(caminho, usuarioAtual)
		if err != nil {
			return fmt.Errorf("[ERRO] Falha ao apagar arquivo: %v", err)
		}
		fmt.Printf("[SUCESSO] Arquivo apagado: %s\n", caminho)
		return nil
	}

	// Comando não reconhecido
	return fmt.Errorf("[ERRO] Comando não reconhecido.")
}
