package app

type User struct {
	Id			int		`json:"-" db:"id"`
	Login		string	`json:"login" binding:"required"`
	Password	string	`json:"password" binding:"required"`
	Age			string	`json:"age" binding:"required"`
}