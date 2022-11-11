package service

import (
	"context"
	"database/sql"
	"errors"
	"monorepo/src/auth_service/pkg/entity"
	"monorepo/src/auth_service/storage"
	pb "monorepo/src/idl/auth_service"
	"monorepo/src/libs/log"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

type AuthService struct {
	storage storage.IStorage
	logger  log.Factory
	pb.UnimplementedAuthServiceServer
	tracer opentracing.Tracer
}

// NewUserService ...
func New(storage storage.IStorage, logger log.Factory, tracer opentracing.Tracer) *AuthService {
	return &AuthService{
		storage: storage,
		logger:  logger,
		tracer:  tracer,
	}
}

func (as *AuthService) StaffLogin(ctx context.Context, req *pb.StaffLoginRequest) (*pb.AuthResponse, error) {

	as.logger.For(ctx).Info("StaffLogin started", zap.String("PhoneNumber", req.PhoneNumber), zap.String("Username", req.Username))

	res, err := as.storage.Authenitication().StaffLogin(*req)
	if err != nil {
		as.logger.For(ctx).Error("failed to staff login username or password didn't match", zap.Error(err))
		return nil, status.Error(codes.Unauthenticated, "username or password is incorrect")
	}
	as.logger.For(ctx).Info("StaffLogin finished", zap.String("PhoneNumber", req.PhoneNumber), zap.String("Username", req.Username))

	return &res, nil
}

func (as *AuthService) StaffSignUp(ctx context.Context, req *pb.StaffSignUpRequest) (*pb.AuthResponse, error) {

	as.logger.For(ctx).Info("StaffSignUp started", zap.String("PhoneNumber", req.PhoneNumber), zap.String("Username", req.Username))
	// if span := opentracing.SpanFromContext(ctx); span != nil {
	// 	span := as.tracer.StartSpan("Query database", opentracing.ChildOf(span.Context()))
	// 	span.SetTag("param.phoneNumber", req.PhoneNumber)
	// 	span.SetTag("param.username", req.Username)
	// 	ext.SpanKindRPCClient.Set(span)
	// 	defer span.Finish()
	// 	ctx = opentracing.ContextWithSpan(ctx, span)
	// }

	res, err := as.storage.Authenitication().StaffSignUp(*req)

	if errors.Is(err, bcrypt.ErrHashTooShort) {
		as.logger.For(ctx).Error("failed to generate default password hash", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	} else if errors.Is(err, sql.ErrNoRows) {
		as.logger.For(ctx).Error("failed to insert resource into db", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	} else if err != nil {
		as.logger.For(ctx).Error("username is already signed up", zap.Error(err))
		return nil, status.Error(codes.AlreadyExists, "user is already signed up")
	}
	as.logger.For(ctx).Info("StaffSignUp finished", zap.String("PhoneNumber", req.PhoneNumber), zap.String("Username", req.Username))

	return &res, nil
}

func (as *AuthService) StaffResetPassword(ctx context.Context, req *pb.StaffResetPasswordRequest) (*pb.Empty, error) {
	as.logger.For(ctx).Info("StaffResetPassword started", zap.String("StaffID", req.StaffID))

	err := as.storage.Authenitication().StaffResetPassword(ctx, entity.ReqResetPassword{
		StaffID:     req.StaffID,
		NewPassword: req.NewPassword,
	})

	if err != nil {
		as.logger.For(ctx).Error("failed to reset staff's password", zap.Error(err))
		return &pb.Empty{}, status.Error(codes.Internal, err.Error())
	}

	as.logger.For(ctx).Info("StaffResetPassword finished", zap.String("StaffID", req.StaffID))

	return &pb.Empty{}, nil
}
