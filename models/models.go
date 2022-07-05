package models

type ToDoList struct {
	ID     int    `json:"_id,omitempty" bson:"_id,omitempty"`
	Task   string `json:"task,omitempty"`
	Status bool   `json:"status,omitempty"`
}
