package postgres

import (
	"github.com/jmoiron/sqlx"
	pb "monorepo/src/idl/auth_service"
)

type authRepo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *authRepo {
	return &authRepo{db: db}
}

func (r *authRepo) CustomerLogin(req pb.CustomerLoginRequest) (pb.AuthResponse, error) {
	var res pb.AuthResponse

	return res, nil
}

func (r *authRepo) CustomerSignUp(req pb.CustomerSignUpRequest) (pb.AuthResponse, error) {
	var res pb.AuthResponse

	return res, nil
}

func (r *authRepo) StaffLogin(req pb.StaffLoginRequest) (pb.AuthResponse, error) {
	var res pb.AuthResponse

	return res, nil
}

func (r *authRepo) StaffSignUp(req pb.StaffSignUpRequest) (pb.AuthResponse, error) {
	var res pb.AuthResponse

	return res, nil
}
