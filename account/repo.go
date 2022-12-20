package account

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"
)

var ReposErr = errors.New("Unable to handle Repo Request")

type repo struct {
	db     *sql.DB
	logger log.Logger
}

// CreateUser implements Repository
func (repo *repo) CreateUser(ctx context.Context, user User) error {
	sql := `INSERT INTO users(id,email,password)
			VALUES($1,$2,$3)`

	if user.Email == "" || user.Password == "" {
		return ReposErr
	}

	_, err := repo.db.ExecContext(ctx, sql, user.ID, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

// GetUser implements Repository
func (repo *repo) GetUser(ctx context.Context, id string) (string, error) {
	var email string
	sql := `SELECT email FROM users WHERE id = $1`
	err := repo.db.QueryRow(sql, id).Scan(&email)
	if err != nil {
		return "", ReposErr
	}

	return email, nil
}

func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}
