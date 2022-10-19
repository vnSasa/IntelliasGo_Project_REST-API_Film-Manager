package repository

import (
	"github.com/jmoiron/sqlx"
)

type Authorization interface {

}

type FilmsList interface {

}

type DirectorsFilms interface {

}

type FavouriteFilms interface {

}

type WishFilms interface {

}

type Repository struct {
	Authorization
	FilmsList
	DirectorsFilms
	FavouriteFilms
	WishFilms
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}