package app

import (
	"errors"
)

type DirectorsList struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name" binding:"required"`
	DateOfBirth string `json:"date_of_birth" db:"date_of_birth" binding:"required"`
}

type FilmsList struct {
	ID         int     `json:"id" db:"id"`
	Name       string  `json:"name" db:"name" binding:"required"`
	Genre      string  `json:"genre" db:"genre" binding:"required"`
	DirectorID int     `json:"director_id" db:"director_id" binding:"required"`
	Rate       float32 `json:"rate" db:"rate" binding:"required"`
	Year       int     `json:"year" db:"year" binding:"required"`
	Minutes    float32 `json:"minutes" db:"minutes" binding:"required"`
}

type UserFavoriteFilms struct {
	ID     int
	UserID int
	FilmID int
}

type UserWishFilms struct {
	ID     int
	UserID int
	FilmID int
}

type UpdateDirectorInput struct {
	Name        *string `json:"name"`
	DateOfBirth *string `json:"date_of_birth"`
}

func (i UpdateDirectorInput) Validate() error {
	if i.Name == nil && i.DateOfBirth == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type UpdateFilmInput struct {
	Name       *string  `json:"name"`
	Genre      *string  `json:"genre"`
	DirectorID *int     `json:"direcor_id"`
	Rate       *float32 `json:"rate"`
	Year       *int     `json:"year"`
	Minutes    *float32 `json:"minutes"`
}

func (i UpdateFilmInput) Validate() error {
	if i.Name == nil &&
		i.Genre == nil &&
		i.DirectorID == nil &&
		i.Rate == nil &&
		i.Year == nil &&
		i.Minutes == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type FiltersFilmsInput struct {
	Count    *int     `json:"count"`
	Genre    *string  `json:"genre"`
	MinRate  *float32 `json:"minimum-rate"`
	SortRate *string  `json:"sort-rate"`
	SortYear *string  `json:"sort-year"`
	SortTime *string  `json:"sort-time"`
}
