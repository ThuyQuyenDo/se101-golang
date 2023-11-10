package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Commom app error
var (
	ErrRecordNotFound     = errors.New("record not found")
	ErrPasswordIsNotValid = errors.New("password must have from 8 to 30 characters")

	ErrEmailIsNotValid = errors.New("email is not valid")
	ErrEmailHasExisted = errors.New("email has existed")

	ErrPhoneIsNotValid     = errors.New("phone is not valid")
	ErrPhoneHasNotVerified = errors.New("phone has not been verified")
	ErrPhoneHasExisted     = errors.New("phone has existed")

	ErrLoginFailed = errors.New("email and password are not valid")

	ErrFullNameIsEmpty = errors.New("full name can not be blank")
	ErrFullNameTooLong = errors.New("full name too long, max character is 30")

	ErrEmailOrPasswordInvalid = errors.New("email or password invalid")
)

// Custom app error
type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorized(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Key:        key,
	}
}

func NewCustomError(root error, msg, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}

	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

// Find root error (recursive alg)
func (e *AppError) RootError() error {
	// Type assertions (means: check property `RootErr` of `e` is of type AppError)
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}
	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func ErrDB(err error) *AppError {
	return NewErrorResponse(err, "something went wrong with DB", err.Error(), "DB_ERROR")
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "invalid request", err.Error(), "ERROR_INVALID_REQUEST")
}

func ErrInternal(err error) *AppError {
	return NewErrorResponse(err, "internal error", err.Error(), "ERROR_INVALID_REQUEST")
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ERR_CANNOT_DELETE_%s", entity),
	)
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot create %s", strings.ToLower(entity)),
		fmt.Sprintf("ERR_CANNOT_CREATE_%s", entity),
	)
}

func ErrCannotGetEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot get %s", strings.ToLower(entity)),
		fmt.Sprintf("ERR_CANNOT_GET_%s", entity),
	)
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot update %s", strings.ToLower(entity)),
		fmt.Sprintf("ERR_CANNOT_UPDATE_%s", entity),
	)
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ERR_CANNOT_DELETE_%s", entity),
	)
}

func ErrEntityDeleted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s deleted", strings.ToLower(entity)),
		fmt.Sprintf("ERR_%s_DELETED", entity),
	)
}

func ErrEntityExisted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s already exists", strings.ToLower(entity)),
		fmt.Sprintf("ERR_%s_ALREADY_EXISTS", entity),
	)
}

func ErrEntityNotFound(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s not found", strings.ToLower(entity)),
		fmt.Sprintf("ERR_%s_NOT_FOUND", entity),
	)
}

func ErrNoPermission(err error) *AppError {
	return NewCustomError(
		err,
		"you have no permission",
		"ERR_NO_PERMISSION",
	)
}
