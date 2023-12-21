package chatbot

import (
	"ChatBot/service"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"
)

type chat struct {
	url string
}

func New(url string) service.ChatBot {
	return chat{url: url}
}

// in-memory cache to cache bot responses
var (
	cache     = make(map[string]string)
	cacheLock sync.Mutex
)

func (c chat) GenerateReply(input string) string {
	cacheLock.Lock()
	defer cacheLock.Unlock()

	if reply, ok := cache[input]; ok {
		return reply // Return cached response if available
	}

	modelURL := c.url

	requestData := map[string]interface{}{
		"inputs": input,
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		log.Fatalf("Error marshaling request body: %v", err)
	}

	req, err := http.NewRequest("POST", modelURL, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making the request: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %d - %s", response.StatusCode, response.Status)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var responseData map[string]interface{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		log.Fatalf("Error unmarshaling response body: %v", err)
	}

	generatedText, ok := responseData["generated_text"].(string)
	if !ok {
		log.Fatal("Failed to extract generated text")
	}

	cache[input] = generatedText // Cache the response

	log.Println("Reply:", generatedText)

	return generatedText
}
