package seeds

import (
	"github.com/fazaalexander/montirku-be/helper/password"
	ue "github.com/fazaalexander/montirku-be/modules/entity/user"
)

func CreateUser() []*ue.User {
	hashPasswordUser, _ := password.HashPassword("user1234")
	hashPasswordAdmin, _ := password.HashPassword("admin1234")
	users := []*ue.User{
		{
			Email:         "admin@gmail.com",
			Username:      "admin",
			Password:      string(hashPasswordAdmin),
			RoleId:        1,
			EmailVerified: true,
		},
		{
			Email:         "user@gmail.com",
			Username:      "user",
			Password:      string(hashPasswordUser),
			RoleId:        2,
			EmailVerified: true,
		},
	}

	return users
}
