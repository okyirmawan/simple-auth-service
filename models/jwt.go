package models

import "os"

var JwtKey []byte

func InitJwtKey() {
	JwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

	if len(JwtKey) == 0 {
		JwtKey = []byte("default_secret_key")
	}
}
