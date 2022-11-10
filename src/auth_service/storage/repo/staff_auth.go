package repo

import (
	pb "monorepo/src/idl/auth_service"
)

// Defining Base interface for Authentication
type IAuthStorage interface {
	StaffLogin(pb.StaffLoginRequest) (pb.AuthResponse, error)
	StaffSignUp(pb.StaffSignUpRequest) (pb.AuthResponse, error)
}
