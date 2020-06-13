package errors

import "errors"

var ErrInitTypeInvalid = errors.New("{ERROR_COMMAND} Type of init is invalid")
var ErrInitArgsInvalid = errors.New("{ERROR_COMMAND} Type of args is invalid, is expected 2 arguments")
var ErrArgsRepositoryOrCommandInvalid = errors.New(
	"{ERROR_COMMAND} Type of args of the [REPOSITORY] or [GENERATE_TYPE] is invalid")
var ErrDirectoryPathInvalid = errors.New("{ERROR_COMMAND} Directory path is invalid")
var ErrModuleNameInvalid = errors.New("{ERROR_COMMAND} Directory path is invalid")
