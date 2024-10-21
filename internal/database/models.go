package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `gorm:"unique;not null"`
	Email        string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
}

type Room struct {
	gorm.Model
	Name  string `gorm:"unique;not null"`
	Url   string `gorm:"unique;not null"`
	Scale int
}

type RoomUserRef struct {
	UserID int
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	RoomID int
	Room   Room `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type Task struct {
	gorm.Model
	RoomID       int
	Room         Room `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Name         string
	Description  string
	MeanEstimate *float32
}

type RoomTasksRef struct {
	RoomID int
	Room   Room `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	TaskID int
	Task   Task `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type Estimate struct {
	gorm.Model
	UserID int
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Value  int
}

type TaskEstimatesRef struct {
	TaskID     int
	Task       Task
	EstimateID int
	Estimate   Estimate `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
