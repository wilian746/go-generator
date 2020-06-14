package errors

import "errors"

var ErrInitTypeInvalid = errors.New("{ERROR_COMMAND} Type of init is invalid")
var ErrInitArgsInvalid = errors.New("{ERROR_COMMAND} Type of args is invalid, is expected 2 arguments")
var ErrArgsRepositoryOrCommandInvalid = errors.New(
	"{ERROR_COMMAND} Type of args of the [REPOSITORY] or [GENERATE_TYPE] is invalid")
var ErrDirectoryPathInvalid = errors.New("{ERROR_COMMAND} Directory path is invalid")
var ErrHostNameInvalid = errors.New("{ERROR_COMMAND} Host name is invalid")
var ErrUsernameInvalid = errors.New("{ERROR_COMMAND} Username is invalid")
var ErrProjectNameInvalid = errors.New("{ERROR_COMMAND} Project name is invalid")
