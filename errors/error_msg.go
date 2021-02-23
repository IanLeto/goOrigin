package errors

type Errno struct {
	Code    int
	Message string
}

func (e Errno) Error() string {
	return e.Message
}

type Err struct {
	Code    int
	Message string
	Err     error
}
