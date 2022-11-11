package entity

type ReqResetPassword struct {
	StaffID     string
	NewPassword string
}

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
