package service

import "LdapAdmin/app/system/model"

func AddMenuService(req *model.AddMenuReq) (int, error) {
	return model.AddMenu(model.Menu{
		Active:      1,
		Name:        req.Name,
		Path:        req.Path,
		Description: req.Description,
	})
}

func DeleteMenuService(req *model.DeleteMenuReq) error {
	return model.DeleteMenu(req.IDS)
}

func GetMenusService() {

}

func ModifyMenuService(req *model.ModifyMenuReq) error {
	return model.ModifyMenu(req.ID, model.Menu{
		Name:        req.Name,
		Path:        req.Path,
		Description: req.Description,
	})
}
