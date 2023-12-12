package seeds

import (
	ue "github.com/fazaalexander/montirku-be/modules/entity/user"
)

func CreateUserDetail() []*ue.UserDetail {
	userDetails := []*ue.UserDetail{
		{
			FirstName:    "Saiful",
			LastName:     "Jamil",
			Phone:        "08917283129283",
			ProfilePhoto: "",
			UserId:       1,
		},
		{
			FirstName:    "Saeful",
			LastName:     "Anwar",
			Phone:        "08917283109283",
			ProfilePhoto: "",
			UserId:       2,
		},
		{
			FirstName:    "Anwar",
			LastName:     "Saipul",
			Phone:        "08917283109283",
			ProfilePhoto: "",
			UserId:       3,
		},
	}

	return userDetails
}
