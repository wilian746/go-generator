package errors

import "errors"

var ErrInitTypeEmpty = errors.New("{ERROR_COMMAND} Type of init is empty")
var ErrInitTypeInvalid = errors.New("{ERROR_COMMAND} Type of init is invalid")
var ErrDirectoryPathInvalid = errors.New("{ERROR_COMMAND} Directory path is invalid")
var ErrModuleNameInvalid = errors.New("{ERROR_COMMAND} Directory path is invalid")