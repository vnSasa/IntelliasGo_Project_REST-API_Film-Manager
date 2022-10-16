package app

type FilmList struct {
	Id int	`json:"id" db:"id"`
	Name string	`json:"name" db:"name" binding:"required"`
	Genre string	`json:"genre" db:"genre" binding:"required"`
	DirectorId int	`json:"director_id" db:"director_id" binding:"required"`
	Rate float32	`json:"rate" db:"rate" binding:"required"`
	Year int	`json:"year" db:"year" binding:"required"`
	Minutes float32	`json:"minutes" db:"minutes" binding:"required"`
}

type DirectorList struct {
	Id int	`json:"id" db:"id"`
	Name string	`json:"name" db:"name" binding:"required"`
	DateOfBirth string	`json:"date_of_birth" db:"date_of_birth" binding:"required"`
}

type UserFavoriteFilms struct {
	Id int
	UserId int	
	FilmId int
}

type UserWishFilms struct {
	Id int
	UserId int
	FilmId int
}