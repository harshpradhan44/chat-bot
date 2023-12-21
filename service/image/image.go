package image

import (
	"ChatBot/service"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
)

type image struct{}

func New() service.ImageProcessing {
	return image{}
}

func (i image) UploadImage(file multipart.File, fileName string) error {
	// Create a new file on the server
	dst, err := os.Create("./uploaded_Images/" + fileName)
	if err != nil {
		fmt.Println("Error Creating the File")
		fmt.Println(err)
		return err
	}

	defer dst.Close()

	// Copy the file content from the uploaded file to the destination file
	if _, err := io.Copy(dst, file); err != nil {
		fmt.Println(err)
		return err
	}

	log.Print("Successfully Uploaded File")

	return nil
}

func (i image) DownloadImage(fileName string) string {
	imagePath := "./uploaded_Images/" + fileName

	return imagePath
}
