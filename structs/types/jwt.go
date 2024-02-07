package types

type JWTAdmin struct {
	ID          int      `json:"id"`
	DisplayName string   `json:"displayName"`
	Email       string   `json:"email"`
	Password    string   `json:"password"`
	PhotoURL    string   `json:"photoURL"`
	PhoneNumber string   `json:"phoneNumber"`
	Country     string   `json:"country"`
	Address     string   `json:"address"`
	State       string   `json:"state"`
	City        string   `json:"city"`
	ZipCode     string   `json:"zipCode"`
	About       string   `json:"about"`
	Role        string   `json:"role"`
	IsPublic    int      `json:"isPublic"`
	AccessIDs   []string `json:"access_ids"`
}

type JWTResponse struct {
	AccessToken string   `json:"accessToken"`
	Admin       JWTAdmin `json:"admin"`
}
