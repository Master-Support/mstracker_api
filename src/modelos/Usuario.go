package modelos

import "time"


//Usuário representa um usuário que utliza o sistema
type Usuario struct {
	ID uint64 `json:"id,omitempty"`
	Nome string `json:"nome,omitempty"`
	Email string `json:"email,omitempty"`
	Senha string `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriandoEm,omitempty"`
}
