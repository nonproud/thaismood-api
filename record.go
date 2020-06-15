package main

type Activity struct {
	username     string `json: "username"`
	ActivityName string `json:"activity"`
	Date         string `json:"date"`
}

type Mood struct {
	Username string `json:"username"`
	Mood     int    `json:"mood"`
	Level    int    `json:"level"`
	Date     string `json:"date"`
}

type Exercise struct {
	Username  string `json:"username"`
	StepCount int    `json:"step_count"`
	Date      string `json:"date"`
}

type Diary struct {
	Username string `json:"username"`
	Title    string `json:"title"`
	Story    string `json:"story"`
	Date     string `json:"date"`
}

type Sleep struct {
	Username  string `json:"username"`
	TotalTime string `json:"total_time"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Date      string `json:"date"`
}
