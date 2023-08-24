package auth

type user struct {
	Id       uint64 `db:"user_id"`
	Login    string `db:"user_login"`
	Password string `db:"user_password"`
}
