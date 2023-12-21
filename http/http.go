package http

import (
	"ChatBot/models"
	"ChatBot/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Handler struct {
	imgSvc service.ImageProcessing
	botSvc service.ChatBot
}

func New(imgSvc service.ImageProcessing, botSvc service.ChatBot) Handler {
	return Handler{botSvc: botSvc, imgSvc: imgSvc}
}

func (h Handler) Get(ctx *gin.Context) {
	var body models.ChatBot

	intent := ctx.Request.URL.Query().Get("intent")
	if intent == "" {
		ctx.JSON(http.StatusBadRequest, "parameter intent is missing")
	}

	input := ctx.Request.URL.Query().Get("input")

	body.Intent = intent
	body.Input = input

	err := h.process(ctx, body)
	if err != nil {
		return
	}
}

// process redirects to the respective module on the basis of intent provided
func (h Handler) process(ctx *gin.Context, body models.ChatBot) error {
	switch body.Intent {

	case models.ImageUploadIntent:
		err := ctx.Request.ParseMultipartForm(10 << 20) // Parse the request with a maximum file size of 10MB
		if err != nil {
			log.Fatalf("error parsing the image - %v", err)
		}

		file, handler, err := ctx.Request.FormFile("file") // Access the file part from the request
		if err != nil {
			log.Fatalf("Error Retrieving the File - %v", err)
		}

		defer file.Close()

		err = h.imgSvc.UploadImage(file, handler.Filename)
		if err != nil {
			return err
		}

		ctx.JSON(http.StatusOK, "Image Uploaded Successfully!")

	case models.ImageDownloadIntent:
		img := h.imgSvc.DownloadImage(body.Input)
		ctx.File(img)

	case models.ChatIntent:
		reply := h.botSvc.GenerateReply(body.Input)
		ctx.JSON(http.StatusOK, models.ChatBot{Output: reply})

	default:
		ctx.JSON(http.StatusBadRequest, "cannot process provided intent")
	}

	return nil
}
