package app

type exception struct {
	msg string
}

func (e exception) Error() string {
	return e.msg
}

// NewError 新建一个错误对象
func NewError(msg string) error {
	e := new(exception)
	e.SetErrMsg(msg)
	return e
}
func (e exception) SetErrMsg(msg string) {
	e.msg = msg
}
