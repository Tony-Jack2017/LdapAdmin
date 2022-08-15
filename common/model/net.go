package model

type Data struct {
	Total int32       `json:"total"`
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
	Msg  int         `json:"msg"`
	Data interface{} `json:"data"`
}

type ResponseData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Data   `json:"data"`
}
