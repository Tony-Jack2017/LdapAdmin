package service

import (
	"LdapAdmin/app/system/model"
	model2 "LdapAdmin/app/user/model"
	"LdapAdmin/common/constant"
	"LdapAdmin/common/middleware"
	"LdapAdmin/common/util"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func LoginService(req *model.LoginReq, ip string) (interface{}, int, error) {
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
	plaintextSql, errDecryptSql := util.DecryptPassword([]byte(userInfo.Password))
	plaintextReq, errDecryptReq := util.DecryptPassword([]byte(req.Password))
	if errDecryptSql != nil || errDecryptReq != nil {
		return nil, constant.DecryptError, errors.New(fmt.Sprintf("Decrypt password is failed !!!"))
	}
	if string(plaintextReq) != string(plaintextSql) {
		return nil, constant.PasswordError, errors.New(fmt.Sprintf("Wrong password !!!"))
	}
	token, errCode, errToken := tokenHandler(userInfo.Account, ip, userInfo.Password, "active", "admin")
	if errToken != nil {
		return nil, errCode, errToken
	}
	data := map[string]interface{}{
		"user_info": userInfo,
		"token":     token,
	}
	return data, 0, nil
}

func tokenHandler(account string, ip string, password string, active string, role string) (string, int, error) {
	token, err := model.GetToken(&model.GetTokenReq{
		Account: account,
	})
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return "", constant.SqlError, errors.New(fmt.Sprintf("get token failed !!! \n error: %v", err))
		} else {
			tokenString, errGen := middleware.GenerateToken(account, password, active, role)
			if errGen != nil {
				return "", constant.TokenError, errors.New(fmt.Sprintf("generate token failed !!! \n error: %v", errGen))
			}
			errSql := model.AddToken(model.Token{
				Account:     account,
				IP:          ip,
				TokenString: tokenString,
			})
			if errSql != nil {
				return "", constant.SqlError, errors.New(fmt.Sprintf("add token failed !!! \n error: %v", errSql))
			}
			return tokenString, 0, nil
		}
	}
	claims, errParse := middleware.ParseToken(token.TokenString)
	// when token is expired
	if errParse != nil {
		tokenString, errGen := middleware.GenerateToken(claims.Account, claims.Password, claims.Active, claims.Role)
		if errGen != nil {
			return "", constant.TokenError, errors.New(fmt.Sprintf("generate token failed !!! \n error: %v", errGen))
		}
		errSql := model.ModifyToken(token.ID, model.Token{
			TokenString: tokenString,
		})
		if errSql != nil {
			return "", constant.SqlError, errors.New(fmt.Sprintf("modify token failed !!! \n error: %v", errSql))
		}
		return tokenString, 0, nil
	} else {
		return token.TokenString, 0, nil
	}
}
