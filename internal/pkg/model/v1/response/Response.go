package response

type BaseResponse struct {
	//Id   string `json:"id"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
