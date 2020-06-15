package main

type General struct {
	Username     string `json: "username"`
	Dob          string `json: "dob"`
	Caffeine     bool   `json: "caffeine"`
	DrugAddict   bool   `json: "drug_addict"`
	Created      string `json: "created"`
	LastModified string `json: "last_modified"`
}

type Patient struct {
	Username     string  `json: "username"`
	Dob          string  `json: "dob"`
	Caffeine     bool    `json: "caffeine"`
	DrugAddict   bool    `json: "drug_addict"`
	Sex          int     `json: "sex"`
	Pregnant     bool    `json: "pregnant"`
	Weight       float64 `json: "weight"`
	Height       float64 `json "height"`
	BMI          float64 `json: "bmi"`
	Disease      string  `json: disease`
	Created      string  `json: "created"`
	LastModified string  `json: "last_modified"`
}