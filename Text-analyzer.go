package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

type Document struct {
	ID       string `json:"id"`
	Language string `json:"language"`
	Text     string `json:"text"`
}

type Documents struct {
	Documents []Document `json:"documents"`
}

type Sentiment struct {
	ID        string `json:"id"`
	Sentiment string `json:"sentiment"`
}

type SentimentResponse struct {
	Documents []Sentiment `json:"documents"`
}

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")
	endpoint := os.Getenv("ENDPOINT")

	// API Defintion
	sentimentURL := fmt.Sprintf("%s/text/analytics/v3.1/sentiment", endpoint)

	client := resty.New()

	// Define the document to analyze
	document := Document{
		ID:       "1",
		Language: "en",
		Text:     "I am sad",
	}
	documents := Documents{
		Documents: []Document{document},
	}

	// Make the request
	resp, err := client.R().
		SetHeader("Ocp-Apim-Subscription-Key", apiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(documents).
		Post(sentimentURL)

	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}

	var sentimentResponse SentimentResponse
	if err := json.Unmarshal(resp.Body(), &sentimentResponse); err != nil {
		log.Fatalf("Error parsing response: %v", err)
	}

	// sentiment response Format
	for _, sentiment := range sentimentResponse.Documents {
		sentimentText := convertSentimentToText(sentiment.Sentiment)
		fmt.Printf("Code: %s, Mood: %s\n", sentiment.ID, sentimentText)
	}
}

func convertSentimentToText(sentiment string) string {
	switch sentiment {
	case "positive":
		return "It seems to radiate positivity"
	case "neutral":
		return "The sentiment is balanced"
	case "negative":
		return "It seems to carry a touch of negativity"
	default:
		return "Unknown Sentiment"
	}
}
