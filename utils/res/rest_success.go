package res

type Success struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func NewRestSuccess(code int, status string, data interface{}) *Success {
	return &Success{
		Code:   code,
		Status: status,
		Data:   data,
	}
}
