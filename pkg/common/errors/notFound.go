package errors

type ErrNotFound struct {
	msg string
}

func (e *ErrNotFound) Error() string {
	return e.msg
}

func ThrowNotFound(msg string) *ErrNotFound {
	return &ErrNotFound{msg}
}
