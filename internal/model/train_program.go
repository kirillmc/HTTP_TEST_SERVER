package model

type TrainProgram struct {
	Id          int64      `json:"id"`
	ProgramName string     `json:"program_name"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	TrainDays   []TrainDay `json:"train_days"`
}
