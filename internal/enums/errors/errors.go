package errors

import "errors"

var ErrInitTypeInvalid = errors.New("{ERROR_COMMAND} Type of init is invalid")
var ErrInitDBCmdInvalid = errors.New("{ERROR_COMMAND} Type of database command is invalid")
var ErrDirectoryPathInvalid = errors.New("{ERROR_COMMAND} Directory path is invalid")
var ErrModuleNameInvalid = errors.New("{ERROR_COMMAND} Directory path is invalid")
