package view

// RegisterAccount 註冊帳號
type RegisterAccount struct {
	NickName string  `json:"nickName"`
	Password string  `json:"password"`
	Age      uint8   `json:"age"`
	Height   float32 `json:"height"`
	Weight   float32 `json:"weight"`
	Email    string  `json:"email"`
	Name     string  `json:"name" validate:"required"`
	Type     int8    `json:"type" validate:"required"`
}
