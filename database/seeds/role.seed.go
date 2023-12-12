package seeds

import (
	re "github.com/fazaalexander/montirku-be/modules/entity/role"
)

func CreateRoles() []*re.Role {
	roles := []*re.Role{
		{Role: "Admin"},
		{Role: "User"},
		{Role: "Mitra"},
	}

	return roles
}
