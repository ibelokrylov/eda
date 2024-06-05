package config

type Result struct {
	Status int         `json:"response_status"`
	Data   interface{} `json:"data"`
}

func BaseResult(status int, data interface{}) Result {
	return Result{
		Status: status,
		Data:   data,
	}
}
