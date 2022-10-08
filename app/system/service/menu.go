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
	var parentId int
	if req.ParentID != 0 {
		menu, err := model.GetMenuByIDAndPath(req.ParentID, "")
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return 0, constant.SqlError, err
			} else {
				return 0, constant.NotExistError, errors.New(fmt.Sprintf("The parent menu id %d is not found", req.ParentID))
			}
		}
		parentId = menu.ID
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
		return 0, constant.ExistError, errors.New(fmt.Sprintf("The path '%s' is exist", req.Path))
	}
}

func DeleteMenuService(req *model.DeleteMenuReq) error {
	return model.DeleteMenu(req)
}

func GetMenuListService(req *model.GetMenuListReq) ([]model.Menu, int64, error) {
	req.Page, req.Size = util.FilterPageOption(req.Page, req.Size)
	return model.GetMenuList(req)
}

func ModifyMenuService(req *model.ModifyMenuReq) error {
	var menu = model.Menu{
		Status:      req.Status,
		Name:        req.Name,
		Description: req.Description,
	}
	if err := model.ModifyMenu(req.ID, menu, req); err != nil {
		return err
	}
	if req.Path != "" {
		if req.OldPath != req.Path {
			if err := model.ModifyMenuPathBatch(req.OldPath, req.Path); err != nil {
				return err
			}
		}
	}
	return nil
}
