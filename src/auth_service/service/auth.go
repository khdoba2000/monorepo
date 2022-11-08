package service

import (
	"context"
	"monorepo/src/auth_service/configs"
	"monorepo/src/auth_service/storage"
	pb "monorepo/src/idl/auth_service"
	l "monorepo/src/libs/logger"
)

type AuthService struct {
	storage storage.IStorage
	logger  l.Logger
	pb.UnimplementedAuthServiceServer
}

// NewUserService ...
func New(storage storage.IStorage, config configs.Configuration, log l.Logger) *AuthService {
	return &AuthService{
		storage: storage,
		logger:  log,
	}
}

func (as *AuthService) CustomerLogin(ctx context.Context, req *pb.CustomerLoginRequest) (*pb.AuthResponse, error) {
	var res pb.AuthResponse

	return &res, nil
}

func (as *AuthService) CustomerSignUp(ctx context.Context, req *pb.CustomerSignUpRequest) (*pb.AuthResponse, error) {
	var res pb.AuthResponse

	return &res, nil
}

func (as *AuthService) StaffLogin(ctx context.Context, req *pb.StaffLoginRequest) (*pb.AuthResponse, error) {
	var res pb.AuthResponse

	return &res, nil
}

func (as *AuthService) StaffSignUp(ctx context.Context, req *pb.StaffSignUpRequest) (*pb.AuthResponse, error) {
	var res pb.AuthResponse

	return &res, nil
}

// CheckField ...
// func (s *AuthService) CheckField(ctx context.Context, req *pb.CustomerLoginRequest) (*pb.AuthResponse, error) {
// 	resp, err := s.storage.User().CheckField(req.Field, req.Value)
// 	if err == sql.ErrNoRows {
// 		s.logger.Error("Error while checking field, Not Found", l.Any("req", req))
// 		return nil, status.Error(codes.NotFound, "Not Found")
// 	} else if err == repo.ErrInvalidField {
// 		s.logger.Error("Error while checking field, Invalid input", l.Error(err), l.Any("req", req))
// 		return nil, status.Error(codes.InvalidArgument, "Invalid argument")
// 	} else if err != nil {
// 		s.logger.Error("Error while checking field", l.Error(err), l.Any("req", req))
// 		return nil, status.Error(codes.Internal, "Internal Server Error")
// 	}

// 	return &pb.CheckFieldResponse{
// 		Exists: resp,
// 	}, nil
// }
