# chat-bot
A basic chat bot written in Golang.

Endpoints:

GET /bot 

example : localhost:8080/bot?intent=chat&input=how are you?

**intent, input** are query param.

**intent** = [image_upload, image_download or chat]

NOTE: When intent is image_uploade, provide image under "file" key within form-data.
