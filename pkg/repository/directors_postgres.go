package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager"
	"strings"
)

type DirectorPostgres struct {
	db *sqlx.DB
}

func NewDirectorPostgres(db *sqlx.DB) *DirectorPostgres {
	return &DirectorPostgres{db: db}
}

func (r *DirectorPostgres) Create(director app.DirectorsList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createDirectorQuery := fmt.Sprintf("INSERT INTO %s (name, date_of_birth)"+
		"VALUES ($1, $2) RETURNING id", directorTable)
	row := tx.QueryRow(createDirectorQuery, director.Name, director.DateOfBirth)
	if err := row.Scan(&id); err != nil {
		if rb := tx.Rollback(); rb != nil {
			return 0, err
		}

		return 0, err
	}

	return id, tx.Commit()
}

func (r *DirectorPostgres) GetAll() ([]app.DirectorsList, error) {
	var directors []app.DirectorsList
	query := fmt.Sprintf("SELECT id, name, date_of_birth FROM %s", directorTable)
	err := r.db.Select(&directors, query)

	return directors, err
}

func (r *DirectorPostgres) GetByID(directorID int) (app.DirectorsList, error) {
	var directors app.DirectorsList
	query := fmt.Sprintf("SELECT id, name, date_of_birth FROM %s WHERE id = $1", directorTable)
	err := r.db.Get(&directors, query, directorID)

	return directors, err
}

func (r *DirectorPostgres) Update(directorID int, input app.UpdateDirectorInput) error {
	setValue := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.Name != nil {
		setValue = append(setValue, fmt.Sprintf("name=$%d", argID))
		args = append(args, *input.Name)
		argID++
	}

	if input.DateOfBirth != nil {
		setValue = append(setValue, fmt.Sprintf("date_of_birth=$%d", argID))
		args = append(args, *input.DateOfBirth)
		argID++
	}

	setQuery := strings.Join(setValue, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", directorTable, setQuery, argID)

	args = append(args, directorID)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)

	return err
}

func (r *DirectorPostgres) Delete(directorID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", directorTable)
	_, err := r.db.Exec(query, directorID)

	return err
}
