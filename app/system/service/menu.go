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

func DeleteMenuService() {

}

func GetMenusService() {

}

func ModifyMenuService() {

}
