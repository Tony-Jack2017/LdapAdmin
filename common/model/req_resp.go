package model

/*
	$ Request Model
*/

type RangeTime struct {
	StartTime string `form:"start_time" time_format:"2006-01-02 15:04:05" json:"start_time"` // start of range time
	EndTime   string `form:"end_time" time_format:"2006-01-02 15:04:05" json:"end_time"`     // end of range time
}

type PaginationOption struct {
	Page int `form:"page" json:"page"` // Paging query page
	Size int `form:"size" json:"size"` // Paging query size
}

/*
	$ Response Model
*/

type Data struct {
	Total int64       `json:"total"`
	List  interface{} `json:"list"`
}

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ResponseErr struct {
	Code   int    `json:"code"`
	ErrMsg string `json:"errMsg"`
}

type ResponseCommon struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ResponseData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Data   `json:"data"`
}
