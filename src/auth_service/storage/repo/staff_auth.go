package repo

import (
	"monorepo/src/auth_service/pkg/mappers"
	pb "monorepo/src/idl/auth_service"
)

// Defining Base interface for Authentication
type IAuthStorage interface {
	StaffLogin(*mappers.StaffLoginReq) (pb.AuthResponse, error)
	StaffSignUp(*mappers.StaffSignUpReq) (pb.AuthResponse, error)
}
