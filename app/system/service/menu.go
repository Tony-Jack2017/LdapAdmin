package service

import (
	"LdapAdmin/app/system/model"
	"LdapAdmin/common/constant"
	"LdapAdmin/common/util"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func AddMenuService(req *model.AddMenuReq) (int, int, error) {
	var parentId *int
	if req.ParentID != 0 {
		menu, err := model.GetMenuByIDAndPath(req.ParentID, "")
		if err != nil {
			return 0, constant.SqlError, err
		}
		parentId = &menu.ID
	}
	if _, err := model.GetMenuByIDAndPath(0, req.Path); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, constant.SqlError, err
		}
		id, errAdd := model.AddMenu(model.Menu{
			Active:      1,
			Status:      req.Status,
			Name:        req.Name,
			Path:        req.Path,
			Description: req.Description,
			ParentID:    parentId,
		})
		if errAdd != nil {
			return 0, constant.SqlError, errAdd
		}
		return id, 0, nil
	} else {
		return 0, constant.SqlError, errors.New(fmt.Sprintf("The path '%s' is exist", req.Path))
	}

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
	if err := model.ModifyMenu(req.ID, model.Menu{
		Status:      req.Status,
		Name:        req.Name,
		Description: req.Description,
	}); err != nil {
		return err
	}
	if req.NewPath != "" {
		if req.OldPath != req.NewPath {
			if err := model.ModifyMenuPathBatch(req.OldPath, req.NewPath); err != nil {
				return err
			}
		}
	}
	return nil
}
