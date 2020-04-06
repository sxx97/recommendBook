package tools

type Response struct {
	Status string `json:"status"`
	Msg interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

type Lists struct {
	Data interface{} `json:"data"`
	Total int `json:total`
}

func ApiResource(status string, objects interface{}, msg string) (r *Response) {
	r = &Response{
		status,
		msg,
		objects,
	}
	return
}