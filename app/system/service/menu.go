package service

import (
	"LdapAdmin/app/system/model"
	"LdapAdmin/common/util"
)

func AddMenuService(req *model.AddMenuReq) (int, error) {
	return model.AddMenu(model.Menu{
		Active:      1,
		Name:        req.Name,
		Path:        req.Path,
		Description: req.Description,
	})
}

func DeleteMenuService(req *model.DeleteMenuReq) error {
	return model.DeleteMenu(req)
}

func GetMenusService(req *model.GetMenuListReq) ([]model.Menu, int64, error) {
	req.Page, req.Size = util.FilterPageOption(req.Page, req.Size)
	if req.Active == 0 {
		req.Active = 1
	}
	return model.GetMenuList(req)
}

func ModifyMenuService(req *model.ModifyMenuReq) error {
	return model.ModifyMenu(req.ID, model.Menu{
		Name:        req.Name,
		Path:        req.Path,
		Description: req.Description,
	})
}
