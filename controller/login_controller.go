package controller

import (
	"gin-api/model"
	"gin-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginController struct {
	LoginUsecase usecase.LoginUsecase
}

func NewLoginController(usecase usecase.LoginUsecase) loginController {
	return loginController{
		LoginUsecase: usecase,
	}
}

func (l *loginController) LoginUser(ctx *gin.Context) {

	var formLogin model.FormLogin
	err := ctx.BindJSON(&formLogin)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	login, err := l.LoginUsecase.LoginUser(formLogin)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if login == nil {
		respose := model.Response{
			Message: "usuario ou senha não encontrado",
		}
		ctx.JSON(http.StatusNotFound, respose)
		return
	}

	ctx.JSON(http.StatusOK, login)

}
