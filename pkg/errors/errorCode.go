package errors

import "fmt"

var (
	// ErrInitMySQL init errors
	ErrInitMySQL = func(err error) *Errno {
		return &Errno{Code: 20001, Message: fmt.Sprintf("Error occurred while init mysql backend. Error Detail:%v", err)}
	}

	// ErrBind handler errors
	ErrBind = func(err error) *Errno {
		return &Errno{Code: 30001, Message: fmt.Sprintf("Error occurred while bind data to struct. Error Detail:%v", err)}
	}
)
