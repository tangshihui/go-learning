package model

type HttpResult struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func OK(data interface{}) HttpResult {
	return HttpResult{
		Code:    "200",
		Message: "OK",
		Data:    data,
	}
}

func NotFound() HttpResult {
	return HttpResult{
		Code:    "200",
		Message: "Not found",
		Data:    nil,
	}
}
