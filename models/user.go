package models

type User struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Password     string `json:"-"`
	RefreshToken string `json:"refresh_token"`
}
