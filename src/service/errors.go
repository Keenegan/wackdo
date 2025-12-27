package service

import (
	"fmt"
)

type InvalidParamError struct {
	Reason string
}

func (e *InvalidParamError) Error() string {
	return fmt.Sprintf("invalid parameter : %s'", e.Reason)
}
