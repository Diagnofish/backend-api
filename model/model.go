package model

import (
	"time"
)

type Session struct {
	ID     int       `json:"id" gorm:"primaryKey"`
	Token  string    `json:"token"`
	Email  string    `json:"email"`
	Expiry time.Time `json:"expiry"`
}

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"type:varchar(255);not null"`
	Email     string    `json:"email" gorm:"type:varchar(50);not null"`
	Password  string    `json:"-" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FishDetection struct {
	ID            string    `json:"id" gorm:"PrimaryKey"`
	ImageFilename string    `json:"image_link" gorm:"type:varchar(255)"`
	FishName      string    `json:"fish_name" gorm:"type:varchar(50)"`
	IsSuccess     bool      `json:"is_success"`
	Result        string    `json:"result"`
	Accuracy      float64   `json:"accuracy"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Email         string    `json:"email"`
}

type DetectionClass struct {
	ID          int    `json:"id" gorm:"PrimaryKey"`
	Result      string `json:"result" gorm:"type:varchar(50)"`
	StatusShown string `json:"status_shown" gorm:"type:varchar(50)"`
	Description string `json:"description" gorm:"type:text"`
	Treatment   string `json:"treatment" gorm:"type:text"`
	Prevention  string `json:"prevention" gorm:"type:text"`
}

type UserRegister struct {
	Email          string `json:"email" binding:"required"`
	Password       string `json:"password" binding:"required"`
	RepeatPassword string `json:"repeat_password" binding:"required"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ImageData struct {
	ID            string
	Email         string
	Filename      string
	FileDirectory string
}

type Respp struct {
	IsSuccess         bool    `json:"is_success"`
	Accuracy          float32 `json:"accuracy"`
	PredictionClassId int     `json:"prediction_class_id"`
}

type Credential struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
	Schema       string
}
