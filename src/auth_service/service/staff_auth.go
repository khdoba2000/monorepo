package service

import (
	"context"
	"database/sql"
	"errors"
	"monorepo/src/auth_service/configs"
	"monorepo/src/auth_service/storage"
	pb "monorepo/src/idl/auth_service"
	l "monorepo/src/libs/logger"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (as *AuthService) StaffLogin(ctx context.Context, req *pb.StaffLoginRequest) (*pb.AuthResponse, error) {

	res, err := as.storage.Authenitication().StaffLogin(*req)
	if err != nil {
		as.logger.Error("failed to staff login username or password didn't match", l.Error(err))
		return nil, status.Error(codes.Unauthenticated, "username or password is incorrect")
	}
	return &res, nil
}

func (as *AuthService) StaffSignUp(ctx context.Context, req *pb.StaffSignUpRequest) (*pb.AuthResponse, error) {
	var res pb.AuthResponse
	res, err := as.storage.Authenitication().StaffSignUp(*req)

	if errors.Is(err, bcrypt.ErrHashTooShort) {
		as.logger.Error("failed to generate default password hash", l.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	} else if errors.Is(err, sql.ErrNoRows) {
		as.logger.Error("failed to insert resource into db", l.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	} else if err != nil {
		as.logger.Error("username is already signed up", l.Error(err))
		return nil, status.Error(codes.AlreadyExists, "user is already signed up")
	}

	return &res, nil
}
