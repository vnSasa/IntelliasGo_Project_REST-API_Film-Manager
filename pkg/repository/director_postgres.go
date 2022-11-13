package repository

import (
	"strings"
	"fmt"
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type DirectorPostgres struct {
	db *sqlx.DB
}

func NewDirectorPostgres(db *sqlx.DB) *DirectorPostgres {
	return &DirectorPostgres{db: db}
}

func (r *DirectorPostgres) Create(director app.DirectorList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	
	var id int
	createDirectorQuery := fmt.Sprintf("INSERT INTO %s (name, date_of_birth)" + 
								"VALUES ($1, $2) RETURNING id", directorTable)
	row := tx.QueryRow(createDirectorQuery, director.Name, director.DateOfBirth)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0,err
	}

	return id, tx.Commit()
}

func (r *DirectorPostgres) GetAll() ([]app.DirectorList, error) {
	var directors []app.DirectorList
	query := fmt.Sprintf("SELECT id, name, date_of_birth FROM %s", directorTable)
	err := r.db.Select(&directors, query)

	return directors, err
}

func (r *DirectorPostgres) GetById(directorId int) (app.DirectorList, error) {
	var directors app.DirectorList
	query := fmt.Sprintf("SELECT id, name, date_of_birth FROM %s WHERE id = $1", directorTable)
	err := r.db.Get(&directors, query, directorId)

	return directors, err
}

func (r *DirectorPostgres) Update(directorId int, input app.UpdateDirectorInput) error {
	setValue := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValue = append(setValue, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.DateOfBirth != nil {
		setValue = append(setValue, fmt.Sprintf("date_of_birth=$%d", argId))
		args = append(args, *input.DateOfBirth)
		argId++
	}

	setQuery := strings.Join(setValue, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", directorTable, setQuery, argId)

	args = append(args, directorId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *DirectorPostgres) Delete(directorId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", directorTable)
	_, err := r.db.Exec(query, directorId)

	return err
}