package view

// RegisterAccount 註冊帳號
type RegisterAccount struct {
	NickName string `json:"nickName"`
	Password string `json:"password"` // validate:"required"`
	Age      uint8  `json:"age"`
	Height   uint8  `json:"height"`
	Weight   uint8  `json:"weight"`
	Email    string `json:"email"` // validate:"required"`
	Name     string `json:"name"`  // validate:"required"`
	Type     int8   `json:"type"`  // validate:"required"`
}

// RegisterAccountResponse 註冊帳號
type RegisterAccountResponse struct {
	NickName string `json:"nickName"`
	Age      uint8  `json:"age"`
	Height   uint8  `json:"height"`
	Weight   uint8  `json:"weight"`
	Email    string `json:"email"` // validate:"required"`
	Name     string `json:"name"`  // validate:"required"`
	Type     int8   `json:"type"`  // validate:"required"`
	Error    string `json:"error"`
}

// LoginInfo 登入所需資訊
type LoginInfo struct {
	Name string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	//Type     int32  `json:"type" validate:"required"`
	//Links    struct {
	//	OTPBindURL   string `json:"otpBindURL"`
	//	OTPVerifyURL string `json:"otpVerifyURL"`
	//	ResourceURL  string `json:"resourceURL"`
	//} `json:"links"`
}

