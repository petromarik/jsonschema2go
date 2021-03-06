// Code generated by jsonschema2go. DO NOT EDIT.
package foo

import (
	"fmt"
)

// Bar is generated from https://example.com/testdata/generate/expanded/foo/bar.json
type Bar struct {
	Foo
	BarAllOf1
}

// Validate returns an error if this value is invalid according to rules defined in https://example.com/testdata/generate/expanded/foo/bar.json
func (m *Bar) Validate() error {
	if err := m.Foo.Validate(); err != nil {
		return err
	}
	if err := m.BarAllOf1.Validate(); err != nil {
		return err
	}
	return nil
}

// BarAllOf1 is generated from https://example.com/testdata/generate/expanded/foo/bar.json#/allOf/1
type BarAllOf1 struct {
	Parent *Foo `json:"parent,omitempty"`
}

// Validate returns an error if this value is invalid according to rules defined in https://example.com/testdata/generate/expanded/foo/bar.json#/allOf/1
func (m *BarAllOf1) Validate() error {
	if m.Parent != nil {
		if err := m.Parent.Validate(); err != nil {
			if err, ok := err.(valErr); ok {
				return &validationError{
					errType:  err.ErrType(),
					message:  err.Message(),
					path:     append([]interface{}{"Parent"}, err.Path()...),
					jsonPath: append([]interface{}{"parent"}, err.JSONPath()...),
				}
			}
			return err
		}
	}
	return nil
}

// Foo is generated from https://example.com/testdata/generate/expanded/foo/foo.json
type Foo struct {
	Bar
}

// Validate returns an error if this value is invalid according to rules defined in https://example.com/testdata/generate/expanded/foo/foo.json
func (m *Foo) Validate() error {
	if err := m.Bar.Validate(); err != nil {
		return err
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
