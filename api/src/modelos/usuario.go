package modelos

import (
	"errors"
	"strings"
	"time"
)

// Usuario representa um usuário utilizando a rede social
type Usuario struct {
	ID       uint64    `json:"id, omitempty`
	Nome     string    `json:"nome, omitempty`
	Nick     string    `json:"nick, omitempty`
	Email    string    `json:"email, omitempty`
	Senha    string    `json:"senha, omitempty`
	CriadoEm time.Time `json:"CriadoEm, omitempty`
}

// Preparar chama métodos de validar e formatar usuário recebido
func (usuario *Usuario) Preparar() error {
	if erro := usuario.validar(); erro != nil {
		return erro
	}

	usuario.formatar()
	return nil
}

func (usuario *Usuario) validar() error {
	if usuario.Nome == "" {
		return errors.New("O campo nome é obrigatório.")
	}

	if usuario.Nick == "" {
		return errors.New("O campo nick é obrigatório.")
	}

	if usuario.Email == "" {
		return errors.New("O campo e-mail é obrigatório.")
	}

	if usuario.Senha == "" {
		return errors.New("O campo senha é obrigatório.")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
