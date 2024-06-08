package core

import "slices"

type Role struct {
	Name   string
	Scopes []string
}

type RoleInterface interface {
	HasPermission(permission string) bool
}

type Roles []Role

type RolesInterface interface {
	GetRole(name string) Role
}

// HasPermission returns true if the Role has the given permission.
func (r Role) HasPermissions(permission string) bool {
	return slices.Contains(r.Scopes, permission)
}

func (r Roles) GetRole(name string) Role {
	for _, role := range r {
		if role.Name == name {
			return role
		}
	}
	return Role{}
}
