package seeds

import (
	"github.com/fazaalexander/montirku-be/helper/password"
	ue "github.com/fazaalexander/montirku-be/modules/entity/user"
)

func CreateUser() []*ue.User {
	hashPasswordUser, _ := password.HashPassword("user1234")
	hashPasswordAdmin, _ := password.HashPassword("admin1234")
	hashPassworMitra, _ := password.HashPassword("mitra1234")
	users := []*ue.User{
		{
			Email:         "admin@gmail.com",
			Password:      string(hashPasswordAdmin),
			RoleId:        1,
			EmailVerified: true,
		},
		{
			Email:         "user@gmail.com",
			Password:      string(hashPasswordUser),
			RoleId:        2,
			EmailVerified: true,
		},
		{
			Email:         "mitra@gmail.com",
			Password:      string(hashPassworMitra),
			RoleId:        3,
			EmailVerified: true,
		},
		{
			Email:         "user1@gmail.com",
			Password:      string(hashPasswordUser),
			RoleId:        2,
			EmailVerified: true,
		},
		{
			Email:         "user2@gmail.com",
			Password:      string(hashPasswordUser),
			RoleId:        2,
			EmailVerified: true,
		},
		{
			Email:         "user3@gmail.com",
			Password:      string(hashPasswordUser),
			RoleId:        2,
			EmailVerified: true,
		},
	}

	return users
}
