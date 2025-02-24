package models

type User struct {
	ID        uint64 `json:id,omitempty`
	Nome      string `json:nome,omitempty`
	Nick      string `json:nick,omitempty`
	Email     string `json:email,omitempty`
	Password  string `json:password,omitempty`
	CreatedAt string `json:createdAt,omitempty`
}
