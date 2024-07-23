package usecase

import (
	"gin-api/model"
	"gin-api/repository"
)

type LoginUsecase struct {
	repository repository.LoginRepository
}

func NewLoginUsecase(repo repository.LoginRepository) LoginUsecase {
	return LoginUsecase{
		repository: repo,
	}
}

func (lu *LoginUsecase) LoginUser(formLogin model.FormLogin) (*model.Login, error) {

	login, err := lu.repository.LoginUser(formLogin)

	if err != nil {
		return nil, err

	}

	return login, nil
}
