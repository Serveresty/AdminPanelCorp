package models

type RoleAction struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
}
