package repository

import (
	"fmt"
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager"
	"github.com/jmoiron/sqlx"
)

type DirectorPostgres struct {
	db *sqlx.DB
}

func NewDirectorPostgres(db *sqlx.DB) *DirectorPostgres {
	return &DirectorPostgres{db: db}
}

func (r *DirectorPostgres) Create(userLogin string, director app.DirectorList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	
	var id int
	createDirectorQuery := fmt.Sprintf("INSERT INTO %s (name, date_of_birth) VALUES ($1, $2) RETURNING id", directorTable)
	row := tx.QueryRow(createDirectorQuery, director.Name, director.DateOfBirth)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0,err
	}

	return id, tx.Commit()
}