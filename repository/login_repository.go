package repository

import (
	"database/sql"
	"fmt"
	"gin-api/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type LoginRepository struct {
	connection *sql.DB
}

func NewLoginRepository(connection *sql.DB) LoginRepository {
	return LoginRepository{
		connection: connection,
	}
}

func (lr *LoginRepository) LoginUser(formLogin model.FormLogin) (*model.Login, error) {

	query, err := lr.connection.Prepare("SELECT id,usuario, '' as Jwt FROM usuario WHERE usuario  = $1 and senha = $2")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var login model.Login

	err = query.QueryRow(formLogin.User, formLogin.Password).Scan(
		&login.ID,
		&login.Name,
		&login.Jwt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	// Gerar Jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": login.ID,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("87878567"))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	login.Jwt = tokenString

	query.Close()

	return &login, nil

}
