package model

type User struct {
	BaseModel
	Email        string
	Username     string
	Bio          string
	Image        string
	PasswordHash string
}

type Profile struct {
	BaseModel
	Username  string
	Bio       string
	Image     string
	following bool
}
