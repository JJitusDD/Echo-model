package error

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrInvalid  = errors.New("input invalid")
)

type ErrCodeMesg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e ErrCodeMesg) Error() string {
	return e.Msg
}
