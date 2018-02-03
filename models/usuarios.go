package models

//Representa a tabela de usuarios no banco de dados
type Usuarios struct {
	ID    int64  `db:"id" json:"id"`
	Nome  string `db:"nome" json:"nome"`
	Email string `db:"email" json:"email"`
}
