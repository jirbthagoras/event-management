package middleware

import (
	"database/sql"
	"github.com/julienschmidt/httprouter"
	"jirbthagoras/event-management/exception"
	"jirbthagoras/event-management/helper"
	"jirbthagoras/event-management/repository"
	"net/http"
)

type AuthMiddleware struct {
	DB              *sql.DB
	AdminRepository repository.AdminRepository
}

func NewAuthMiddleware(DB *sql.DB, adminRepository repository.AdminRepository) *AuthMiddleware {
	return &AuthMiddleware{DB: DB, AdminRepository: adminRepository}
}

func (middleware *AuthMiddleware) Handle(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		authorization := r.Header.Get("X-API-Key")
		if authorization == "" {
			helper.PanicIfError(exception.NewAuthenticationError("Auth Key is missing"))
		}

		tx, err := middleware.DB.Begin()
		helper.PanicIfError(err)
		defer helper.CommitOrRollback(tx)

		_, err = middleware.AdminRepository.FindByToken(r.Context(), tx, authorization)
		if err != nil {
			helper.PanicIfError(exception.NewAuthenticationError("Auth Key Invalid"))
		}

		next(w, r, p)
	}
}
