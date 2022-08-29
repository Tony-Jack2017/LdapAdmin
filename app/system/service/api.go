package service

import "LdapAdmin/app/system/model"

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

func GetApiListService() {

}

func ModifyApiService(req *model.ModifyApiReq) error {
	return model.ModifyApi(req.ID, model.Api{
		Name:        req.Name,
		Path:        req.Path,
		Description: req.Description,
		ApiGroupID:  req.ApiGroupID,
	})
}

/* $ ApiGroup Service */

func AddApiGroup(req *model.AddApiGroupReq) (int, error) {
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

func DeleteApiGroup(req *model.DeleteApiGroupReq) error {
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

func GetApiGroupList(req *model.GetApiGroupListReq) (int64, error) {
	return 0, nil
}

func ModifyApiGroup(req *model.ModifyApiGroupReq) error {
	return nil
}
