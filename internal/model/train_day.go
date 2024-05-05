package model

type TrainDay struct {
	Id          int64      `json:"id"`
	DayName     string     `json:"day_name"`
	Description string     `json:"description"`
	Exercises   []Exercise `json:"exercises"`
}
