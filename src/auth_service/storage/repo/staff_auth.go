package repo

import (
	"context"
	"monorepo/src/auth_service/pkg/entity"
	pb "monorepo/src/idl/auth_service"
)

// Defining Base interface for Authentication
type IAuthStorage interface {
	StaffLogin(context.Context, entity.StaffLoginReq) (pb.AuthResponse, error)
	StaffSignUp(context.Context, entity.StaffSignUpReq) (pb.AuthResponse, error)
	StaffResetPassword(ctx context.Context, req entity.ReqResetPassword) error
}
