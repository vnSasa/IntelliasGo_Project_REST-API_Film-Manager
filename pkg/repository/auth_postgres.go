package repository

import (
	"strings"
	"errors"
	"strconv"
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
	query := fmt.Sprintf("INSERT INTO %s (login, password_hash, age) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, admin.Login, admin.Password, admin.Age)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) CreateUser(user app.User) (int, error) {
	var id int
	err := r.validLogin(user.Login)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("your enter data for login is not valid: %s", err))
	}
	query := fmt.Sprintf("INSERT INTO %s (login, password_hash, age) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Login, user.Password, user.Age)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) validLogin(login string) error {
	if len(login) != 10 {
		return errors.New("must have 10 numbers")
	}
	_, err := strconv.Atoi(login)
	if err != nil {
		return errors.New(fmt.Sprintf("your enter data have not only numbers: %s", login))
	}
	loginSlice := strings.Split(login, "")
	for i := 0; i < len(loginSlice); i++ {
		x, _ := strconv.Atoi(loginSlice[0])
		if x != 0 {
			return errors.New(fmt.Sprintf("phone number must have strart from 0, but you have: %d", x))
		}
	}
	return nil
}

func (r *AuthPostgres) DeleteUser(user app.User) (int, error) {
	var id int
	query := fmt.Sprintf("DELETE FROM %s WHERE login=$1 RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Login)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(login, password string) (app.User, error) {
	var user app.User
	query := fmt.Sprintf("SELECT id, login FROM %s WHERE login=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, login, password)

	return user, err
}

func (r *AuthPostgres) GetUserById(id int) (app.User, error) {
	var user app.User
	query := fmt.Sprintf("SELECT login, age FROM %s WHERE id=$1", usersTable)
	err := r.db.Get(&user, query, id)

	return user, err
}