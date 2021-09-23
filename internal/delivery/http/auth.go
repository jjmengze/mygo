package http

import (
	"fmt"
	"github.com/jjmengze/mygo/internal/delivery/http/view"
	"github.com/jjmengze/mygo/internal/dto"
	errMsg "github.com/jjmengze/mygo/pkg/errors"
	"github.com/labstack/echo/v4"

	"github.com/pkg/errors"
	"net/http"
)

func (h *Handler) loginEndpoint(c echo.Context) error {
	ctx := c.Request().Context()

	loginInfo := view.LoginInfo{}

	if err := c.Bind(&loginInfo); err != nil {
		return errors.Wrap(errMsg.ErrInvalidInput, err.Error())
	}

	req := dto.NewLoginDtoUser(&loginInfo)
	user, err := h.svc.Login(ctx, req)
	if err != nil {
		err = errors.New(fmt.Sprintf("登入失敗: %v", err))
		resp := dto.NewAccountResponseDto(user, err)
		return c.JSON(http.StatusUnauthorized, resp)
	}

	return c.JSON(http.StatusOK, dto.NewAccountResponseDto(user, err))
}

func (h *Handler) registerEndpoint(c echo.Context) (err error) {
	ctx := c.Request().Context()

	registerAcc := view.RegisterAccount{}
	if err := c.Bind(&registerAcc); err != nil {
		return errors.Wrap(errMsg.ErrInvalidInput, err.Error())
	}

	//namespace := ""
	//switch identity.AccountType(registerAcc.Type) {
	//case identity.Operator:
	//	namespace = pbiamdef.AccNamespaceOperator.String()
	//case identity.NormalUser:
	//	namespace = pbiamdef.AccNamespaceFenko.String()
	//default:
	//	return errors.WithMessage(errMsg.ErrInvalidInput, "request Type not allow")
	//}

	//create in identity
	req := dto.NewAccountDtoUser(&registerAcc)

	accountResp, err := h.svc.CreateUser(ctx, req)

	if err != nil {
		err = errors.New(fmt.Sprintf("註冊帳號失敗: %v", err))
		resp := dto.NewAccountResponseDto(accountResp, err)
		return c.JSON(http.StatusConflict, resp)
	}
	return c.JSON(http.StatusOK, dto.NewAccountResponseDto(accountResp, err))
}
