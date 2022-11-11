package models

type ReqResetPassword struct {
	NewPassword string `json:"new_password"`
}
