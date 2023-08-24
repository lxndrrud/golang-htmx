package auth

import (
	"database/sql"
	"fmt"
)

type authDataService struct {
	authRepo *authRepo
}

func newAuthDataService(authRepo *authRepo) *authDataService {
	return &authDataService{authRepo}
}

func (ads authDataService) LoginAction(login, password string) (*user, error) {
	user, err := ads.authRepo.GetUser(login, password)
	if err == sql.ErrNoRows {
		err = fmt.Errorf("Пользователь не найден!")
		return user, err
	}
	if err != nil {
		err = fmt.Errorf("authDataService - LoginAction: %v", err.Error())
		return user, err
	}
	return user, nil
}
