package mappers

import (
	"monorepo/src/auth_service/pkg/entity"
	pb "monorepo/src/idl/auth_service"
)

func MapProtoLoginReq(req *pb.StaffLoginRequest) entity.StaffLoginReq {
	r := entity.StaffLoginReq{
		Username:    req.Username,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
	}
	return r
}

func MapProtoSignUpReq(req *pb.StaffSignUpRequest) entity.StaffSignUpReq {
	r := entity.StaffSignUpReq{
		Name:        req.Name,
		Username:    req.Username,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
		Role:        req.Role,
		BranchId:    req.BranchId,
	}
	return r
}
