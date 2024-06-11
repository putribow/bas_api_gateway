package handler

import (
	//"TaskGo/models"
	"TaskGo/usecase"
	//"TaskGo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthInterface interface {
	Login(*gin.Context)
}

type authImplement struct{}

func NewAuth() AuthInterface {
	return &authImplement{}
}

type BodyPayloadAuth struct {
	Username string
	Password string
}

func (a *authImplement) Login(g *gin.Context) {

	bodyPayloadAuth := BodyPayloadAuth{}
	err := g.BindJSON(&bodyPayloadAuth)

	// orm := utils.NewDatabase().Orm
	// db, _ := orm.DB()

	// defer db.Close()

	// user := models.Account{}

	// user.Password = bodyPayloadAuth.Password
	// user.Username = bodyPayloadAuth.Username

	// result := orm.Create(&bodyPayloadAuth)

	// if result.Error != nil {
	// 	g.AbortWithError(http.StatusBadRequest, err)
	// }

	//usecase.NewLogin().Authenticate(bodyPayloadAuth.Username, bodyPayloadAuth.Password)

	if usecase.NewLogin().Authenticate(bodyPayloadAuth.Username, bodyPayloadAuth.Password) {
		g.JSON(http.StatusOK, gin.H{
			"message": "Anda berhasil login",
			"data":    bodyPayloadAuth,
		})
	} else {
		g.JSON(http.StatusBadRequest, gin.H{
			"message": "Anda gagal login",
			"data":    err,
		})
	}

}
