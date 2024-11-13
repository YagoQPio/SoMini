package main

import (
	"fmt"
)

type Usuario struct {
	Username string
	Senha    string
}

var usuarios = make(map[string]Usuario)
var usuarioLogado Usuario

func criarUsuario(username, senha string) {
	usuarios[username] = Usuario{Username: username, Senha: senha}
}

func logarUsuario(username, senha string) {
	usuario, existe := usuarios[username]
	if !existe || usuario.Senha != senha {
		fmt.Println("Nome de usuário ou senha incorretos.")
		return
	}
	usuarioLogado = usuario
	fmt.Println("Usuário logado com sucesso!")
}
