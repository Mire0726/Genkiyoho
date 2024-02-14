package model

type User struct {
    ID   int    `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
    Name string `json:"name"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"` 
}
