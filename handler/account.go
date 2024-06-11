package handler

import (
	"TaskGo/model"
	"TaskGo/models"
	"TaskGo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountInterface interface {
	GetAccount(*gin.Context)
	CreateAccount(*gin.Context)
	UpdateAccount(*gin.Context)
	DeleteAccount(*gin.Context)
	GetBalance(*gin.Context)
}

type accountImplement struct{}

func NewAccount() AccountInterface {
	return &accountImplement{}
}

func (a *accountImplement) GetAccount(g *gin.Context) {

	QueryParam := g.Request.URL.Query()

	name := QueryParam.Get("name")

	accounts := []model.Account{}

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	q := orm
	if name != "" {
		q = q.Where("name = ?", name)
	}

	result := q.Find(&accounts)
	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"eror": result.Error,
		})

	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Get account successfully",
		"data":    accounts,
	})
}

// type BodyPayloadAccount struct {
// 	AccountID string
// 	Name      string
// 	Address   string
// }

func (a *accountImplement) CreateAccount(g *gin.Context) {

	BodyPayload := model.Account{}
	err := g.BindJSON(&BodyPayload)

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	result := orm.Create(&BodyPayload)

	if result.Error != nil {
		g.AbortWithError(http.StatusBadRequest, err)
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Get account successfully",
		"data":    BodyPayload,
	})
}

func (a *accountImplement) UpdateAccount(g *gin.Context) {
	bodyPayload := model.Account{}

	err := g.BindJSON(&bodyPayload)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	id := g.Param("id")
	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	user := model.Account{}

	orm.First(&user, "account_id = ?", id)
	user.Name = bodyPayload.Name
	user.Username = bodyPayload.Username
	orm.Save(user)

	if user.AccountID == "" {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "data not found",
		})
		return
	}
	// QueryParam := g.Request.URL.Query()

	// name := QueryParam.Get("name")

	g.JSON(http.StatusOK, gin.H{
		"message": "update account successfully",
		"data":    user,
	})
}

func (a *accountImplement) DeleteAccount(g *gin.Context) {

	id := g.Param("id")
	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()
	defer db.Close()

	result := orm.Where("account_id = ?", id).Delete(&model.Account{})

	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Account removed successfully",
		"data":    id,
	})
}

type BodyPayloadBalance struct {
	Account_ID string
	Month      int
}

// func (a *accountImplement) GetBalance(g *gin.Context) {

// 	bodyPayloadBal := BodyPayloadBalance{}
// 	err := g.BindJSON(&bodyPayloadBal)

// 	if err != nil {
// 		g.AbortWithError(http.StatusBadRequest, err)
// 	}

// 	g.JSON(http.StatusOK, gin.H{
// 		"message": "Hello guys this API rest for later",
// 	})
// }

func (a *accountImplement) GetBalance(g *gin.Context) {

	BodyPayloadBalance := BodyPayloadBalance{}
	err := g.BindJSON(&BodyPayloadBalance)

	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
	}

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()
	defer db.Close()

	sumResult := struct {
		Total int
	}{}

	result := orm.Model(&models.Transaction{}).
		Select("sum(amount) as total").Where("account_id = ? AND date_part( 'Month' , transaction_date) = ?", BodyPayloadBalance.Account_ID, BodyPayloadBalance.Month).
		Group("account_id").
		Scan(&sumResult)

	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Get total successfully",
		"data":    sumResult,
	})

}
