package repo

import (
	"context"
	"monorepo/src/auth_service/pkg/entity"
	pb "monorepo/src/idl/auth_service"
)

// Defining Base interface for Authentication
type IAuthStorage interface {
	StaffLogin(*entity.StaffLoginReq) (pb.AuthResponse, error)
	StaffSignUp(*entity.StaffSignUpReq) (pb.AuthResponse, error)
	StaffResetPassword(ctx context.Context, req entity.ReqResetPassword) error
}
