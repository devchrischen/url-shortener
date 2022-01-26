package apires

type (
	Base struct {
		Message string `json:"message"	description:"API return Message"`
	}

	Data struct {
		Base
		Data interface{} `json:"data" description:"API return data"`
	}
)
