package models

import "github.com/google/uuid"

type User struct {
	BaseModel
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
	Tasks       []Task `gorm:"foreignKey:UserID"`
}

type GetUserTasks struct {
	UserId uuid.UUID `json:"userId"`
}
