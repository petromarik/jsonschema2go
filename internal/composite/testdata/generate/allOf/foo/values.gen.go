// Code generated by jsonschema2go. DO NOT EDIT.
package foo

import (
	"fmt"
)

// Bar gives you some dumb info
type Bar struct {
	Bar *int64  `json:"bar,omitempty"`
	Foo *string `json:"foo,omitempty"`
}

func (m *Bar) Validate() error {
	if m.Bar == nil {
		return &validationError{
			errType:  "required",
			message:  "field required",
			path:     []interface{}{"Bar"},
			jsonPath: []interface{}{"bar"},
		}
	}
	if m.Foo == nil {
		return &validationError{
			errType:  "required",
			message:  "field required",
			path:     []interface{}{"Foo"},
			jsonPath: []interface{}{"foo"},
		}
	}
	return nil
}

type valErr interface {
	ErrType() string
	JSONPath() []interface{}
	Path() []interface{}
	Message() string
}

type validationError struct {
	errType, message string
	jsonPath, path   []interface{}
}

func (e *validationError) ErrType() string {
	return e.errType
}

func (e *validationError) JSONPath() []interface{} {
	return e.jsonPath
}

func (e *validationError) Path() []interface{} {
	return e.path
}

func (e *validationError) Message() string {
	return e.message
}

func (e *validationError) Error() string {
	return fmt.Sprintf("%v: %v", e.path, e.message)
}

var _ valErr = new(validationError)
