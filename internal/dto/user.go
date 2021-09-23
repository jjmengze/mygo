package dto

import (
	"github.com/jjmengze/mygo/internal/delivery/http/view"
	"github.com/jjmengze/mygo/internal/model"
)

// NewAccountDto ...
func NewAccountDto(input *view.RegisterAccount) *model.User {
	enable := true
	return &model.User{
		Password: input.Password,
		Email:    input.Email,
		Name:     input.Name,
		NikeName: input.NickName,
		Age:      uint8(input.Age),
		Height:   uint8(input.Height),
		Weight:   uint8(input.Weight),
		IsEnable: &enable,
	}
}

// NewAccountResponseDto ...
func NewAccountResponseDto(input *model.User, err error) *view.RegisterAccountResponse {
	if input != nil {
		return &view.RegisterAccountResponse{
			Email:    input.Email,
			Name:     input.Name,
			NickName: input.NikeName,
			Age:      input.Age,
			Height:   input.Height,
			Weight:   input.Weight,
		}
	}
	return &view.RegisterAccountResponse{
		Error: err.Error(),
	}
}
