package gpt

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
	"tweet-analyser/config"
)

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Request struct {
	Model    string    `json:"model"`
	Messages []message `json:"messages"`
}

type Responce struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func New(data string) string {
	api, url := config.GetGptData()
	request := Request{
		Model: "gpt-4",
		Messages: []message{
			{
				Role:    "system",
				Content: "Apply Sentiment Analysis on the given text, and provide a summary of the text by first getting the sentimental value of the original text and then provide a summary of the text with the explanation of weather the data is positive or not based on the value, also use simple language. and these are definitly twitter (also x.com) tweets and make sure to make it one simple paragraph don't menshion sentimental analysis.",
			},
			{
				Role:    "user",
				Content: data,
			},
		},
	}

	jsonReq, err := json.Marshal(request)
	if err != nil {
		log.Println("json error: ", err)
		log.Fatal(err)
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Println("request error: ", err)
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api)

	res, err := client.Do(req)
	if err != nil {
		log.Println("response error: ", err)
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("body error: ", err)
		log.Fatal(err)
	}

	var responce Responce
	err = json.Unmarshal(body, &responce)
	if err != nil {
		log.Println("unmarshal error: ", err)
		log.Fatal(err)
	}

	return responce.Choices[0].Message.Content
}
