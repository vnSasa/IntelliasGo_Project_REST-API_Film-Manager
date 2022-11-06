package repository

import (
	"github.com/sirupsen/logrus"
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres  {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateAdmin(admin app.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (id, login, password_hash, age) values (1, $1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, admin.Login, admin.Password, admin.Age)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) CreateUser(user app.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (login, password_hash, age) values ($1, $2, $3) RETURNING id", usersTable)
	if len(user.Login) != 10 {
		logrus.Error("not phone number")
	} 
	row := r.db.QueryRow(query, user.Login, user.Password, user.Age)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(login, password string) (app.User, error) {
	var user app.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE login=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, login, password)

	return user, err
}