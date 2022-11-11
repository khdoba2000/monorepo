package mappers

import (
	pb "monorepo/src/idl/auth_service"
)

type StaffLoginReq struct {
	Username    string
	Password    string
	PhoneNumber string
}

type StaffSignUpReq struct {
	Name        string
	Username    string
	Password    string
	PhoneNumber string
	Role        string
	BranchId    string
}

func (s *StaffLoginReq) MapProtoLoginReq(req *pb.StaffLoginRequest) {
	s.Username = req.Username
	s.Password = req.Password
	s.PhoneNumber = req.PhoneNumber
}

func (s *StaffSignUpReq) MapProtoSignUpReq(req *pb.StaffSignUpRequest) {

	s.Name = req.Name
	s.Username = req.Username
	s.Password = req.Password
	s.PhoneNumber = req.PhoneNumber
	s.Role = req.Role
	s.BranchId = req.BranchId

}
