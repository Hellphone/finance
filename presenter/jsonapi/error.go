package jsonapi

import (
	"encoding/json"

	"github.com/hellphone/finance/domain"
)

type Error struct {
	Code  domain.ErrorCode `json:"code"`
	Title string           `json:"title"`
}

type ErrorResponse struct {
	Errors []*Error `json:"errors"`
}

func NewErrorResponse(errs ...*domain.Error) *ErrorResponse {
	resp := &ErrorResponse{
		Errors: make([]*Error, len(errs)),
	}

	for i, e := range errs {
		resp.Errors[i] = &Error{
			Code:  e.Code(),
			Title: e.Error(),
		}
	}

	return resp
}

func (e *ErrorResponse) ToBytes() []byte {
	bs, err := json.Marshal(e)
	if err != nil {
		return []byte(`{"errors":[{"title":"Unable to format error"}]}`)
	}
	return bs
}
