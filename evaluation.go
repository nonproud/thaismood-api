package main

type Evaluation struct {
	Username string `json:"username"`
	Type     string `json:"type"`
	Score    int    `json:"score"`
	Date     string `json:"date"`
}
