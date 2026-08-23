package dto

type UserUpdateStatusReq struct {
	Status string `json:"status" binding:"required,oneof=active blocked"`
}

type UserUpdateRoleReq struct {
	Role string `json:"role" binding:"required,oneof=student admin"`
}
