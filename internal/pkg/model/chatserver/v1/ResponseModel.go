package models

type BaseResponse struct {
	//Id   string `json:"id"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CreateUserResponse struct {
	BaseResponse
	Data string `json:"data"`
}
