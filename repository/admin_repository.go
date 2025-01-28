package repository

import (
	"context"
	"database/sql"
	"errors"
	"jirbthagoras/event-management/domain/model"
	"jirbthagoras/event-management/exception"
)

type AdminRepository interface {
	FindByToken(ctx context.Context, tx *sql.Tx, token string) (model.Admin, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (model.Admin, error)
	Update(ctx context.Context, tx *sql.Tx, admin model.Admin) error
}

type AdminRepositoryImpl struct {
}

func NewAdminRepositoryImpl() *AdminRepositoryImpl {
	return &AdminRepositoryImpl{}
}

func (repository AdminRepositoryImpl) FindByToken(ctx context.Context, tx *sql.Tx, token string) (model.Admin, error) {
	query := "SELECT id FROM admin WHERE token = ?"
	rows, err := tx.QueryContext(ctx, query, token)
	defer rows.Close()
	admin := model.Admin{}
	if err != nil {
		return admin, err
	}

	if rows.Next() {
		err := rows.Scan(&admin.Id)
		return admin, err
	} else {
		return admin, errors.New("not found")
	}
}

func (repository AdminRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (model.Admin, error) {
	query := "SELECT id, email, password, token FROM admin WHERE email = ?"
	rows, err := tx.QueryContext(ctx, query, email)
	admin := model.Admin{}

	defer rows.Close()
	if err != nil {
		return admin, err
	}

	if rows.Next() {
		err := rows.Scan(&admin.Id, &admin.Email, &admin.Password, &admin.Token)
		return admin, err
	} else {
		return admin, exception.NewNotFoundError("Account Not Found")
	}
}

func (repository AdminRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, admin model.Admin) error {
	query := "UPDATE admin SET token = ?, email = ?, password = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, admin.Token, admin.Email, admin.Password, admin.Id)
	return err
}
