package service

import "LdapAdmin/app/system/model"

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
