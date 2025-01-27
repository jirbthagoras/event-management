package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"jirbthagoras/event-management/domain/web"
	"jirbthagoras/event-management/helper"
	"jirbthagoras/event-management/repository"
)

type AdminService interface {
	Login(ctx context.Context, request *web.AdminLoginRequest) *web.AdminResponse
}

type AdminServiceImpl struct {
	AdminRepository repository.AdminRepository
	DB              *sql.DB
	Validator       *validator.Validate
}

func NewAdminServiceImpl(adminRepository repository.AdminRepository, DB *sql.DB, validator *validator.Validate) *AdminServiceImpl {
	return &AdminServiceImpl{AdminRepository: adminRepository, DB: DB, Validator: validator}
}

func (service AdminServiceImpl) Login(ctx context.Context, request *web.AdminLoginRequest) *web.AdminResponse {
	err := service.Validator.StructCtx(ctx, &request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	admin, err := service.AdminRepository.FindByEmail(ctx, tx, request.Email)
	helper.PanicIfError(err)

	err = helper.ComparePassword(admin.Password, request.Password)
	helper.PanicIfError(err)

	admin.Token = uuid.NewString()

	err = service.AdminRepository.Update(ctx, tx, admin)
	helper.PanicIfError(err)

	return helper.ToAdminResponse(admin)
}
