package response

type UniversalReturn struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type LogRespModel struct {
	Total int64         `json:"total"`
	List  []interface{} `json:"data"`
}
