package model

type Login struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Jwt  string `json:"Jwt"`
}

type FormLogin struct {
	User     string `json:"user"`
	Password string `json:"password"`
}
