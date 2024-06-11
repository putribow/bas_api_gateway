package handler

import (
	//"TaskGo/model"
	"TaskGo/models"
	"TaskGo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionInterface interface {
	TransferBank(*gin.Context)
}

type transactionImplement struct{}

func NewTransaction() TransactionInterface {
	return &transactionImplement{}
}

//type BodyPayloadTransaction struct{}

func (b *transactionImplement) TransferBank(g *gin.Context) {

	bodyPayloadTxn := models.Transaction{}
	err := g.BindJSON(&bodyPayloadTxn)

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()
	result := orm.Create(&bodyPayloadTxn)

	if result.Error != nil {
		g.AbortWithError(http.StatusBadRequest, err)
	}

	// if err != nil {
	// 	g.AbortWithError(http.StatusBadRequest, err)
	// }

	g.JSON(http.StatusOK, gin.H{
		"message": "Get account successfully",
		"data":    bodyPayloadTxn,
	})
}
