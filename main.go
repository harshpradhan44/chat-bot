package main

import (
	"ChatBot/http"
	"ChatBot/service/chatbot"
	"ChatBot/service/image"
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// Load environment variables from the .env file
	err := godotenv.Load("./configs/.env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	viper.AutomaticEnv()

	imgSvc := image.New()
	botSvc := chatbot.New(viper.GetString("MODEL_URL"))
	handler := http.New(imgSvc, botSvc)

	app := gin.Default()

	app.Handle("GET", "/bot", handler.Get)

	err = app.Run()
	if err != nil {
		log.Fatal("server failure")
		return
	}
}
