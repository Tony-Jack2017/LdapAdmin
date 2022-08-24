package service

import (
	"LdapAdmin/app/system/model"
	model2 "LdapAdmin/app/user/model"
	"LdapAdmin/common/constant"
	"LdapAdmin/common/middleware"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func LoginService(req *model.LoginReq) (interface{}, int, error) {

	userInfo, err := model2.GetUserInfo(&model2.GetUserInfoReq{
		Account: req.Account,
	})

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constant.SqlError, errors.New(fmt.Sprintf("Get user info failed !!! \n error: %v", err))
		} else {
			return nil, constant.NotExistError, errors.New(fmt.Sprintf("Account %v not found", req.Account))
		}
	}

	fmt.Println(userInfo)

	return nil, 0, nil
}

func tokenHandler(account string, ip string, password string, active string, role string) (error, int, string) {
	token, err := model.GetToken(&model.GetTokenReq{
		Account: account,
	})
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(fmt.Sprintf("get token failed !!! \n error: %v", err)), constant.SqlError, ""
		} else {
			tokenString, errGen := middleware.GenerateToken(account, password, active, role)
			if errGen != nil {
				return errors.New(fmt.Sprintf("generate token failed !!! \n error: %v", errGen)), constant.TokenError, ""
			}
			errSql := model.AddToken(model.Token{
				Account:     account,
				IP:          ip,
				TokenString: tokenString,
			})
			if errSql != nil {
				return errors.New(fmt.Sprintf("add token failed !!! \n error: %v", errSql)), constant.SqlError, ""
			}
			return nil, 0, tokenString
		}
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
		return nil, 0, token.TokenString
	}
}
