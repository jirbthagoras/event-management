package repository

import (
	"context"
	"database/sql"
	"errors"
	"jirbthagoras/event-management/domain/model"
	"jirbthagoras/event-management/helper"
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
	helper.PanicIfError(err)

	admin := model.Admin{}

	if rows.Next() {
		err := rows.Scan(&admin.Id)
		helper.PanicIfError(err)
	} else {
		return admin, errors.New("not found")
	}

	return admin, nil
}

func (repository AdminRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (model.Admin, error) {
	query := "SELECT id, email, password, token FROM admin WHERE email = ?"
	rows, err := tx.QueryContext(ctx, query, email)
	defer rows.Close()
	helper.PanicIfError(err)

	admin := model.Admin{}

	if rows.Next() {
		err := rows.Scan(&admin.Id, &admin.Email, &admin.Password, &admin.Token)
		helper.PanicIfError(err)
	} else {
		return admin, errors.New("not found")
	}

	return admin, nil
}

func (repository AdminRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, admin model.Admin) error {
	query := "UPDATE admin SET token = ?, email = ?, password = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, admin.Token, admin.Email, admin.Password, admin.Id)
	helper.PanicIfError(err)

	return nil
}
