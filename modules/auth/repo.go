package auth

import (
	"github.com/jmoiron/sqlx"
)

type authRepo struct {
	connection *sqlx.DB
}

func newAuthRepo(connection *sqlx.DB) *authRepo {
	return &authRepo{connection}
}

func (ar authRepo) GetUser(login, password string) (*user, error) {
	var user user
	err := ar.connection.Get(
		&user,
		`
			SELECT 
				u.id as user_id,
				u.login as user_login,
				u.password as user_password
			FROM users as u
			WHERE u.login = $1 AND u.password = $2
		`, login, password,
	)
	return &user, err
}
