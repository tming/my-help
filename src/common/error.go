package common

import "fmt"

var (
	ErrorUnknown = fmt.Errorf("unknown error")
	ErrorTarget  = fmt.Errorf("input parameter string with -s is empty")
)

var errorCodeMap map[error]int

func init() {
	errorCodeMap = map[error]int{
		ErrorUnknown: 99,
	}
}

// GetErrorCode return error code by error
func GetErrorCode(err error) int {
	if v, ok := errorCodeMap[err]; ok {
		return v
	}

	// default error code
	return 99
}
