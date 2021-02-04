package apis

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func (res *Result) SetCode(code int) *Result {
	res.Code = code
	return res
}

func (res *Result) SetMessage(message string) *Result {
	res.Message = message
	return res
}

func (res *Result) SetData(data string) *Result {
	res.Data = data
	return res
}
