package service

import "mime/multipart"

type ChatBot interface {
	GenerateReply(input string) string
}

type ImageProcessing interface {
	UploadImage(file multipart.File, fileName string) error
	DownloadImage(fileName string) string
}
