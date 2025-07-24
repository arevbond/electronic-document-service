package domain

type CreateUserParams struct {
	Login          string `db:"login"`
	HashedPassword string `db:"hashed_password"`
}
