package main

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/handler"
	"bwastartup/helper"
	"bwastartup/transaction"
	"bwastartup/user"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
  // refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
  dsn := "root:root@tcp(127.0.0.1:8889)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil{
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)

	campaignService := campaign.NewService(campaignRepository)
	
	userService := user.NewService(userRepository)

	transactionService := transaction.NewService(transactionRepository, campaignRepository)

	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)

	campaignHandler := handler.NewCampaignHandler(campaignService)

	transactionHandler := handler.NewTransactionHandler(transactionService)

	router := gin.Default()
	api := router.Group("/api/v1")
	router.Static("/images", "./images")

	// user endpoint
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checker", userHandler.CheckEmailAvailability)
	api.POST("/avatars",authMiddleware(authService, userService), userHandler.UploadAvatar)
	// campaign endpoint
	api.GET("/campaigns",campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id",campaignHandler.GetCampaign)
	api.POST("/campaigns",authMiddleware(authService, userService),campaignHandler.CreateCampaign)
	api.PUT("/campaigns/:id",authMiddleware(authService, userService),campaignHandler.UpdateCampaign)
	api.POST("/campaign-images",authMiddleware(authService, userService),campaignHandler.UploadImage)
	// transaction endpoint
	api.GET("/campaigns/:id/transactions",authMiddleware(authService, userService),transactionHandler.GetCampaignTransactions)
	api.GET("/transactions",authMiddleware(authService, userService),transactionHandler.GetUserTransactions)

	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer ") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""

		arrayToken := strings.Split(authHeader, " ")

		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		payload, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(payload["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil{
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}