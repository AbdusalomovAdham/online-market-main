package user

import (
	"main/internal/entity"
)

type Create struct {
	Username  *string `form:"username"`
	Email     string  `form:"email"`
	Avatar    *entity.File
	Role      *int `form:"role"`
	CreatedBy int
	UpdatedBy int
}

type Update struct {
	Username  *string `json:"username" form:"username"`
	Avatar    *entity.File
	Email     *string `json:"email" form:"email" `
	Birthdate *string `json:"birth_date" form:"birth_date" `
}

type Delete struct {
	ID int
}
