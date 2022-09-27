package service

import (
	"LdapAdmin/app/system/model"
	"LdapAdmin/common/util"
	"fmt"
)

/* $ Api Service */

func AddApiService(req *model.AddApiReq) (int, error) {
	return model.AddApi(model.Api{
		Active:      1,
		Name:        req.Name,
		Path:        req.Path,
		Description: req.Description,
		ApiGroupID:  req.ApiGroupID,
	})
}

func DeleteApiService(req *model.DeleteApiReq) error {
	return model.DeleteApi(req.IDS)
}

func GetApiListService(req *model.GetApiListReq) ([]model.Api, int64, error) {
	req.Page, req.Size = util.FilterPageOption(req.Page, req.Size)
	if req.Active == 0 {
		req.Active = 1
	}
	return model.GetApiList(req)
}

func ModifyApiService(req *model.ModifyApiReq) error {
	return model.ModifyApi(req.ID, model.Api{
		Name:        req.Name,
		Path:        req.Path,
		Description: req.Description,
		ApiGroupID:  req.ApiGroupID,
	}, req)
}

/* $ ApiGroup Service */

func AddApiGroupService(req *model.AddApiGroupReq) (int, error) {
	var apiList []model.Api
	if len(req.ApiList) != 0 {
		for _, item := range req.ApiList {
			apiList = append(apiList, model.Api{
				Active:      1,
				Name:        item.Name,
				Path:        item.Path,
				Description: item.Description,
			})
		}
	}
	return model.AddApiGroup(model.ApiGroup{
		Active:      1,
		Name:        req.Name,
		Prefix:      req.Prefix,
		Description: req.Description,
		ApiList:     apiList,
	})
}

func DeleteApiGroupService(req *model.DeleteApiGroupReq) error {
	if errGroup := model.DeleteApiGroup(req.IDS); errGroup != nil {
		return errGroup
	}
	for _, item := range req.IDS {
		if errApi := model.DeleteApiByGroupID(item); errApi != nil {
			return errApi
		}
	}
	return nil
}

func GetApiGroupListService(req *model.GetApiGroupListReq) ([]model.ApiGroup, int64, error) {
	req.Page, req.Size = util.FilterPageOption(req.Page, req.Size)
	fmt.Println(req.Page, req.Size)
	if req.Active == 0 {
		req.Active = 1
	}
	if req.Type == 0 {
		req.Type = 1
	}
	return model.GetApiGroupList(req)
}

func ModifyApiGroupService(req *model.ModifyApiGroupReq) error {
	return model.ModifyApiGroup(req.ID, model.ApiGroup{})
}
