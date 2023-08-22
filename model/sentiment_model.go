package model

type SentimentPayload struct {
	Negative string `json:"negative"`
	Postive  string `json:"positive"`
	Neutral  string `json:"neutral"`
}
