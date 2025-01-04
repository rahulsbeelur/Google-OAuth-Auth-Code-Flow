package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOAuthConfig = &oauth2.Config{}

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading env file")
	}

	googleOAuthConfig = &oauth2.Config{
		ClientID: os.Getenv("GOOGLE_CLIENT_ID"), 
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL: os.Getenv("GOOGLE_REDIRECT_URL"), 
		Scopes: []string{"profile", "email"},
		Endpoint: google.Endpoint,
	}
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/auth/google/callback", callbackHandler)
	_ = r.Run()
}

func callbackHandler(c *gin.Context) {
	code := c.Query("code")
	token, err := googleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Println("Error exchanging code", code)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userInfo, err := GetUserInfo(token.AccessToken)
	if err != nil {
		fmt.Println("Error getting user info", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	signedToken, err := SignJWT(userInfo)
	if err != nil {
		fmt.Println("Error signing token", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": signedToken})
}