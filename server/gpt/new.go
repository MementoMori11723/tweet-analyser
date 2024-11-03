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
	log.Println("GPT New called")
	api, url := config.GetGptData()
	request := Request{
		Model: "gpt-4",
		Messages: []message{
			{
				Role: "system",
				Content: `Analyze the text to determine if it has a positive, neutral, or negative tone. Then, provide a brief summary that captures the main message and mood in one simple, cohesive paragraph. Display the tone as upbeat, neutral, or downbeat based on the sentiment in a separate line.

Here is the format for the output:
<div style="width: 100%; max-width: 512px; margin: 0 auto;">
    <p>Provide the summary here in one cohesive sentence, capturing the overall sentiment and main idea of the text.</p>
    <br>
    Tone: <span style="color: green;">positive</span> <!-- Change "green" to red or gray based on the sentiment result, and also change the positive to either positive or negaitive or neutral based on the result -->
</div>

Output the <div> structure in HTML with the <p> tag for the summary, and the <span> indicating the tone. Do not use markdown or any other formatting, only HTML tags that can be used inside of a div tag.`,
			},
			{
				Role:    "user",
				Content: data,
			},
		},
	}

	log.Println("Request created")
	log.Println("creating json")

	jsonReq, err := json.Marshal(request)
	if err != nil {
		log.Println("json error: ", err)
		log.Fatal(err)
	}

	log.Println("created json")
	log.Println("creating client")

	client := &http.Client{
		Timeout: time.Second * 50,
	}

	log.Println("created client")
	log.Println("creating request")

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Println("request error: ", err)
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api)

	log.Println("created request")
	log.Println("sending request")
	res, err := client.Do(req)
	if err != nil {
		log.Println("response error: ", err)
		log.Fatal(err)
	}

	log.Println("sent request")
	log.Println("reading response")

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("body error: ", err)
		log.Fatal(err)
	}

	log.Println("read response")
	log.Println("unmarshalling response")

	var responce Responce
	err = json.Unmarshal(body, &responce)
	if err != nil {
		log.Println("unmarshal error: ", err)
		log.Fatal(err)
	}

	log.Println("unmarshalled response")
	log.Println("returning response")
	return responce.Choices[0].Message.Content
}
