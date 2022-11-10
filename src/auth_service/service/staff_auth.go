package service

import (
	"context"
	"database/sql"
	"errors"
	"monorepo/src/auth_service/storage"
	pb "monorepo/src/idl/auth_service"
	"monorepo/src/libs/log"
	"time"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
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
	}
}

func (as *AuthService) StaffLogin(ctx context.Context, req *pb.StaffLoginRequest) (*pb.AuthResponse, error) {

	as.logger.For(ctx).Info("StaffLogin req", zap.String("PhoneNumber", req.PhoneNumber), zap.String("Username", req.Username))
	if span := opentracing.SpanFromContext(ctx); span != nil {
		span := as.tracer.StartSpan("Query database", opentracing.ChildOf(span.Context()))
		span.SetTag("param.phoneNumber", req.PhoneNumber)
		span.SetTag("param.username", req.Username)
		ext.SpanKindRPCClient.Set(span)
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}
	//simulate signup reg
	time.Sleep(1 * time.Second)

	res, err := as.storage.Authenitication().StaffLogin(*req)
	if err != nil {
		as.logger.For(ctx).Error("failed to staff login username or password didn't match", zap.Error(err))
		return nil, status.Error(codes.Unauthenticated, "username or password is incorrect")
	}
	return &res, nil
}

func (as *AuthService) StaffSignUp(ctx context.Context, req *pb.StaffSignUpRequest) (*pb.AuthResponse, error) {

	as.logger.For(ctx).Info("StaffSignUp req", zap.String("PhoneNumber", req.PhoneNumber), zap.String("Username", req.Username))
	if span := opentracing.SpanFromContext(ctx); span != nil {
		span := as.tracer.StartSpan("Query database", opentracing.ChildOf(span.Context()))
		span.SetTag("param.phoneNumber", req.PhoneNumber)
		span.SetTag("param.username", req.Username)
		ext.SpanKindRPCClient.Set(span)
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}
	//simulate signup reg
	time.Sleep(1 * time.Second)

	var res pb.AuthResponse
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

	return &res, nil
}
