package model

import "time"

type Login struct {
	Password string `json:"password" binding:"required,decryptiontext"`
	Username string `json:"username" binding:"required,alphanum"`
}

type TokenResponse struct {
	AccessToken    string     `json:"access_token"`
	RefreshToken   string     `json:"refresh_token"`
	AccessExpired  *time.Time `json:"access_expired"`
	RefreshExpired *time.Time `json:"refresh_expired"`
}
