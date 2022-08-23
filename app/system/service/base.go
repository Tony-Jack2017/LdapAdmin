package service

import (
	"LdapAdmin/app/system/model"
	"LdapAdmin/common/constant"
	"LdapAdmin/common/middleware"
	"errors"
	"fmt"
)

func LoginService(req *model.LoginReq) (interface{}, int, error) {

	return nil, 0, nil
}

func tokenHandler(account string, password string, role string) (error, int, string) {
	token, err := model.GetToken(&model.GetTokensReq{
		Account: account,
	})
	if err != nil {
		return errors.New(fmt.Sprintf("get token failed !!! \n error: %v", err)), constant.SqlError, ""
	}
	claims, errParse := middleware.ParseToken(token.TokenString)
	// when token is expired
	if errParse != nil {
		tokenString, errGen := middleware.GenerateToken(claims.Account, claims.Password, claims.Active, claims.Role)
		if errGen != nil {
			return errors.New(fmt.Sprintf("generate token failed !!! \n error: %v", errGen)), constant.TokenError, ""
		}
		errSql := model.ModifyToken(token.ID, model.Token{
			TokenString: tokenString,
		})
		if errSql != nil {
			return errors.New(fmt.Sprintf("modify token failed !!! \n error: %v", errSql)), constant.SqlError, ""
		}
		return nil, 0, tokenString
	} else {

	}

}
