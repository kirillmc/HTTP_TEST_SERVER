package model

type Exercise struct {
	Id           int64    `json:"id"`
	ExerciseName string   `json:"exercise_name"`
	Pictures     []string `json:"pictures"`
	Description  string   `json:"description"`
	Sets         []Set    `json:"sets"`
}
