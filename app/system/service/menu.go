package service

import (
	"LdapAdmin/app/system/model"
	"LdapAdmin/common/util"
)

func AddMenuService(req *model.AddMenuReq) (int, error) {

	var parentId *int

	if req.ParentID != 0 {
		menu, errParent := model.GetMenuByID(req.ParentID)
		if errParent != nil {
			return 0, errParent
		}
		parentId = &menu.ID
	}

	return model.AddMenu(model.Menu{
		Active:          1,
		Status:          req.Status,
		Name:            req.Name,
		Path:            req.Path,
		IsDifferentPath: req.IsDifferentPath,
		Description:     req.Description,
		ParentID:        parentId,
	})
}

func DeleteMenuService(req *model.DeleteMenuReq) error {
	return model.DeleteMenu(req)
}

func GetMenuListService(req *model.GetMenuListReq) ([]model.Menu, int64, error) {
	req.Page, req.Size = util.FilterPageOption(req.Page, req.Size)
	if req.Active == 0 {
		req.Active = 1
	}
	return model.GetMenuList(req)
}

func ModifyMenuService(req *model.ModifyMenuReq) error {
	err := model.ModifyMenu(req.ID, model.Menu{
		Status:      req.Status,
		Name:        req.Name,
		Path:        req.Path,
		Description: req.Description,
	})
	if err != nil {
		return err
	}
	if req.Path != "" {

	}
	return nil
}
