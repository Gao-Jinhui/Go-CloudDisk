package e

const (
	Success            = 200
	Error              = 500
	InvalidParams      = 400
	ErrorValidator     = 600
	ErrorWrongPassword = 700

	ErrorExistUser       = 10001
	ErrorNotExistUser    = 10002
	ErrorNotExistArticle = 10003

	ErrorAuthCheckTokenFail    = 20001
	ErrorAuthCheckTokenTimeout = 20002
	ErrorAuthToken             = 20003
	ErrorAuth                  = 20004

	ErrorTokenGenerate = 30000
)
