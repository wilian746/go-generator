package errors

import "errors"

var ErrInitTypeInvalid = errors.New("{ERROR_COMMAND} Type of init is invalid")
var ErrInitArgsInvalid = errors.New("{ERROR_COMMAND} Type of args is invalid, is expected 2 arguments")
var ErrArgsRepositoryInvalid = errors.New("{ERROR_COMMAND} Type of args of the repository invalid")
var ErrArgsGenerateInvalid = errors.New("{ERROR_COMMAND} Type of args to generate invalid")
var ErrInitDBCmdInvalid = errors.New("{ERROR_COMMAND} Type of database command is invalid")
var ErrDirectoryPathInvalid = errors.New("{ERROR_COMMAND} Directory path is invalid")
var ErrModuleNameInvalid = errors.New("{ERROR_COMMAND} Directory path is invalid")
