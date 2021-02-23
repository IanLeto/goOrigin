package errors

var (
	// init errors
	ErrInitMySQL = &Errno{Code: 20001, Message: "Error occurred while init mysql backend."}
)

