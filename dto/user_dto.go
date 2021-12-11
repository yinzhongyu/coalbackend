package dto

import "ginVue/model"

/**
dto是数据传输对象
*/

//传输到前端
type UserDto struct {
	Name      string `json:name`
	Telephone string `json:telephone`
	PublicKey string `json:pubclicKey`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
		PublicKey: user.PublicKey,
	}
}
