package main

type Activity struct {
	username     string `json: "username"`
	ActivityName string `json:"activity"`
	Date         string `json:"date"`
}

type Diary struct {
	Username string `json:"username"`
	Title    string `json:"title"`
	Story    string `json:"story"`
	Date     string `json:"date"`
}

type Exercise struct {
	Username  string `json:"username"`
	StepCount int    `json:"step_count"`
	Date      string `json:"date"`
}

type Mood struct {
	Username string `json:"username"`
	Mood     int    `json:"mood"`
	Level    int    `json:"level"`
	Date     string `json:"date"`
}

type Sleep struct {
	Username  string `json:"username"`
	TotalTime string `json:"total_time"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Date      string `json:"date"`
}

func CreateActivity(activity Activity) {

}

func CreateDiary(diary Diary) {

}

func CreateExercise(exercise Exercise){

}

func CreateMood(mood Mood){

}

func CreateSleep(sleep Sleep){

}

func EditActivity(activity Activity) {

}

func EditDiary(diary Diary) {

}

func EditExercise(exercise Exercise){

}

func EditMood(mood Mood){

}

func EditSleep(sleep Sleep){

}

func DeleteActivity(activity Activity) {

}

func DeleteMood(mood Mood){

}

func DeleteExercise(exercise Exercise){

}

func DeleteDiary(diary Diary) {

}

func DeleteSleep(sleep Sleep){

}