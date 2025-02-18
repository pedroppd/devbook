package models

type Usuario struct {
	ID        uint   `json:id,omitempty`
	Nome      string `json:nome,omitempty`
	Nick      string `json:nick,omitempty`
	Email     string `json:email,omitempty`
	Password  string `json:password,omitempty`
	CreatedAt string `json:createdAt,omitempty`
}
