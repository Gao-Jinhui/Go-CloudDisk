package v1

type BaseResponse struct {
	//Id   string `json:"id"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CreateUserResponse struct {
	BaseResponse
}

type DeleteUserResponse struct {
	BaseResponse
}
