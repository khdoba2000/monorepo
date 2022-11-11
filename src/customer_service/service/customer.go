package service

import (
	"context"
	"fmt"
	"monorepo/src/customer_service/storage"
	pb "monorepo/src/idl/customer_service"
	"monorepo/src/libs/log"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CustomerService struct {
	storage storage.IStorage
	logger  log.Factory
	tracer  opentracing.Tracer
}

// NewUserService ...
func New(logger log.Factory, tracer opentracing.Tracer, storage storage.IStorage) *CustomerService {
	fmt.Println("Service new")
	customerServer := &CustomerService{
		storage: storage,
		logger:  logger,
		tracer:  tracer,
	}
	return customerServer
}

func (as *CustomerService) Create(ctx context.Context, req *pb.CreateRequest) (*emptypb.Empty, error) {

	as.logger.For(ctx).Info("StaffSignUp started", zap.String("PhoneNumber", req.PhoneNumber), zap.String("name", req.Name))

	as.logger.For(ctx).Info("StaffSignUp finished", zap.String("PhoneNumber", req.PhoneNumber), zap.String("name", req.Name))

	return &emptypb.Empty{}, nil
}
