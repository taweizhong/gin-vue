package dto

import "awesomeProject1/model"

type UserDto struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func GetUserDto(user model.User) UserDto {
	return UserDto{Name: user.Name, Phone: user.Iphone}
}
