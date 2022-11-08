package repo

import (
	pb "monorepo/src/idl/auth_service"
)

// Defining Base interface for Authentication
type IAuthStorage interface {
	CustomerLogin(pb.CustomerLoginRequest) (pb.AuthResponse, error)
	CustomerSignUp(pb.CustomerSignUpRequest) (pb.AuthResponse, error)
	StaffLogin(pb.StaffLoginRequest) (pb.AuthResponse, error)
	StaffSignUp(pb.StaffSignUpRequest) (pb.AuthResponse, error)
}
