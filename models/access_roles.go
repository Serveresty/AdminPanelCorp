package models

type AccessRoles struct {
	Role        string   `json:"role"`
	AccessRoles []string `json:"access_roles"`
}
