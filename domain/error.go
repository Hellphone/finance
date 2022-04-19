package domain

const (
	UnexpectedError ErrorCode = "unexpected_error"
	NotFoundError             = "not_found_error"
	ValidateError             = "validate_error"
)

type ErrorCode string

type Error struct {
	msg  string
	code ErrorCode
}

func (e *Error) Error() string {
	return e.msg
}

func (e *Error) Code() ErrorCode {
	return e.code
}

func (e *Error) IsNotFoundError() bool {
	return e.code == NotFoundError
}

func (e *Error) IsValidateError() bool {
	return e.code == ValidateError
}

func NewError(msg string, code ErrorCode) *Error {
	return &Error{
		msg:  msg,
		code: code,
	}
}

func NewNotFoundError(msg string) *Error {
	return NewError(msg, NotFoundError)
}

func NewValidateError(msg string) *Error {
	return NewError(msg, ValidateError)
}

func ToDomainError(e error) *Error {
	if e == nil {
		return &Error{
			msg:  "Unexpected error",
			code: UnexpectedError,
		}
	}

	if domainError, ok := e.(*Error); ok {
		return domainError
	}

	return &Error{
		msg:  e.Error(),
		code: UnexpectedError,
	}
}
