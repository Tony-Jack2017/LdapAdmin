package model

type Response struct {
}

type ResponseErr struct {
	Code   int    `json:"code"`
	ErrMsg string `json:"errMsg"`
}
