package errors

import "fmt"

type CustomErr struct {
	HTTPCode int
	Code     int
	MsgCode  *int
	Verbs    map[string]string
	Data     interface{}
	Err      error
	Errs     map[string][]interface{}
}

func NewErr(httpCode, code int) *CustomErr {
	return &CustomErr{
		HTTPCode: httpCode,
		Code:     code,
	}
}

func (e *CustomErr) Reset() {
	e.MsgCode = nil
	e.Verbs = map[string]string{}
	e.Data = nil
	e.Err = nil
	e.Errs = map[string][]interface{}{}
}

func (e *CustomErr) GetHttpCode() int {
	return e.HTTPCode
}

func (e *CustomErr) GetCode() int {
	return e.Code
}

func (e *CustomErr) GetMsgCode() int {
	if e.MsgCode != nil {
		return *e.MsgCode
	}
	return e.Code
}

func (e *CustomErr) GetVerbs() map[string]string {
	return e.Verbs
}

func (e *CustomErr) SetMsgCode(code int) *CustomErr {
	e.MsgCode = &code
	return e
}

func (e *CustomErr) SetVerbs(verbs map[string]string) *CustomErr {
	e.Verbs = verbs
	return e
}

func (e *CustomErr) SetError(err error) *CustomErr {
	e.Err = err
	return e
}

func (e *CustomErr) SetErrors(errs map[string][]interface{}) *CustomErr {
	e.Errs = errs
	return e
}

func (e *CustomErr) HasErrors() bool {
	return len(e.Errs) > 0
}

func (e *CustomErr) GetErrors() map[string][]interface{} {
	return e.Errs
}

func (e *CustomErr) SetData(data interface{}) *CustomErr {
	e.Data = data
	return e
}

func (e *CustomErr) GetData() interface{} {
	return e.Data
}

func (e *CustomErr) Error() string {
	if e.Err == nil {
		return GetMessage(e.Code)
	}
	return fmt.Sprintf("%+v", e.Err)
}
