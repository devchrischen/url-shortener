package apires

type (
	Base struct {
		Code    int    `json:"code" description:"API Return Code"`
		Message string `json:"message" description:"API return Message"`
	}

	Data struct {
		Base
		Data interface{} `json:"data" description:"API return data"`
	}
)
