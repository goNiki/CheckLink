package errorsAPP

import "errors"

var (
	ErrInternalDB        = errors.New("error database")
	ErrDecodeJson        = errors.New("error decode json")
	ErrInvalidValidation = errors.New("validation error")
	ErrOutputPdf         = errors.New("error output pdf")
	ErrNoTasks           = errors.New("no tasks")
	ErrNoLinks           = errors.New("no links")
)

var (
	ErrFailedCreateRequest = errors.New("failed to create request")
	ErrFailedRequest       = errors.New("request failed")
	ErrHTTPStatusInvalid   = errors.New("htttp status <200 or >400 ")
)

var (
	ErrMarshalIndent = errors.New("error marshal indent")
	ErrWriteFile     = errors.New("error write file")
)

var (
	ErrSaveTasksToFile   = errors.New("error save tasks to file")
	ErrSaveLinksToFile   = errors.New("error save links to file")
	ErrLoadLinks         = errors.New("error load links")
	ErrLoadTasks         = errors.New("error load tasks")
	ErrCreateDIContainer = errors.New("error when creating the container")
	ErrGetAndCleanTasks  = errors.New("error get and clean tasks")
	ErrCheckLink         = errors.New("error check link")
)
