package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	openAIToken := os.Getenv("OPENAI_TOKEN")
	cli := http.DefaultClient
	messageBodyStruct := map[string]any{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]any{
			{
				"role":    "system",
				"content": "You are a helpfull assistant.",
			},
			{
				"role":    "user",
				"content": "Hello!",
			},
		},
	}
	messageBody, _ := json.Marshal(messageBodyStruct)
	fmt.Println(string(messageBody))
	req, _ := http.NewRequest(
		http.MethodPost,
		"https://api.openai.com/v1/chat/completions",
		bytes.NewBuffer(messageBody),
	)
	req.Header.Add("Authorization", "Bearer "+openAIToken)
	req.Header.Add("Content-Type", "application/json")
	res, err := cli.Do(req)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	defer res.Body.Close()
	b, _ := ioutil.ReadAll(res.Body)
	fmt.Println(res.StatusCode)
	fmt.Println(string(b))
}
