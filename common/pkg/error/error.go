package error

import "fmt"

type WrappedError struct {
	Context string
	ErrType string
	ErrCode string
	Err     error
}

func (w *WrappedError) Error() string {
	return fmt.Sprintf("%s: %v", w.Context, w.Err)
}

func Wrap(err error, info, errtype, code string) *WrappedError {
	return &WrappedError{
		Context: info,
		Err:     err,
		ErrType: errtype,
		ErrCode: code,
	}
}
