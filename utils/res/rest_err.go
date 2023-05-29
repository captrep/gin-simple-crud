package res

type Err struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Error  interface{} `json:"errors"`
}

func NewRestErr(code int, status string, err interface{}) *Err {
	return &Err{
		Code:   code,
		Status: status,
		Error:  err,
	}
}
