package config

type Result struct {
	Status int         `json:"response_status"`
	Data   interface{} `json:"data"`
	Error  interface{} `json:"error"`
}

func BaseResult(status int, data interface{}, err ...interface{}) Result {
	return Result{
		Status: status,
		Data:   data,
		Error:  err,
	}
}
