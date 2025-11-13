package errorsAPP

import "errors"

var (
	ErrInternalDB        = errors.New("error database")
	ErrDecodeJson        = errors.New("error decode json")
	ErrInvalidValidation = errors.New("validation error")
	ErrOutputPdf         = errors.New("error output pdf")
)
