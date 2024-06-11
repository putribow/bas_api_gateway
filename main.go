package main

import (
	"TaskGo/handler"
	"TaskGo/proto"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/go-micro/generator/cmd/protoc-gen-micro/plugin/micro"
	"github.com/gin-contrib/cors"
	"github.com/go-micro/plugins/v4/client/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"*"},
	}))

	addrServiceTransactionOpt := client.WithAddress(":3636")
	clientSrvTransaction := grpc.NewClient()

	srvTransaction := micro.NewService(
		micro.Client(clientSrvTransaction),
	)

	srvTransaction.Init(
		micro.Name("service-transaction"),
		micro.Version("latest"),
	)

	accountRoute := r.Group("/account")
	accountRoute.GET("/get", handler.NewAccount().GetAccount)
	accountRoute.POST("/create", handler.NewAccount().CreateAccount)
	accountRoute.PATCH("/update/:id", handler.NewAccount().UpdateAccount)
	accountRoute.DELETE("/delete/:id", handler.NewAccount().DeleteAccount)
	accountRoute.POST("/getbalance", handler.NewAccount().GetBalance)

	transactionRoute := r.Group("/transaction")
	transactionRoute.POST("/transferbank", handler.NewTransaction().TransferBank)
	transactionRoute.GET("/get", func(g *gin.Context) {
		clientResponse, err := proto.NewServiceTransactionService("service-transaction", srvTransaction.Client()).Login(context.Background(), &proto.LoginRequest{
			Username: "Admin",
			Password: "admin123",
		}, addrServiceTransactionOpt)

		if err != nil {
			g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		g.JSON(http.StatusOK, gin.H{
			"data": clientResponse,
		})

	})

	authRoute := r.Group("/auth")
	authRoute.POST("/login", handler.NewAuth().Login)

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
