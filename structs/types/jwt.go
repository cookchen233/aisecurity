package types

import (
	"time"
)

type JWTAdmin struct {
	ID          int      `json:"id"`
	DisplayName string   `json:"displayName"`
	PhotoURL    string   `json:"photoURL"`
	PhoneNumber string   `json:"phoneNumber"`
	AccessIDs   []string `json:"access_ids"`
}

type JWTData struct {
	Admin   JWTAdmin `json:"admin"`
	Persist time.Duration
}

type JWTResponse struct {
	Admin       JWTAdmin `json:"admin"`
	AccessToken string   `json:"accessToken"`
}
