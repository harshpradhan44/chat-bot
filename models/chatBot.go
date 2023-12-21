package models

type ChatBot struct {
	Intent string `json:"intent,omitempty"`
	Input  string `json:"input,omitempty"`
	Output string `json:"output,omitempty"`
}
