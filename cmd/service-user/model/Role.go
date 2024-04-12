package model

type Role struct {
	Id          uint
	Name        string
	Description string
}

type RolePermission struct {
	Role       Role
	Permission Permission
}
