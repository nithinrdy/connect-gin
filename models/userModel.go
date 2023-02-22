package models

type UserModel struct {
	Id           int
	Created_at   string
	Username     string
	Email        string
	Nickname     string
	PasswordHash string
	RefreshToken string
}
