package util

var defaultPage = 0
var defaultSize = 20

func FilterPageOption(page int, pageSize int) (int, int) {
	if !(pageSize > 0 && pageSize <= 1000) || !(page > 0) {
		return defaultPage, defaultSize
	}
	return page, pageSize
}
