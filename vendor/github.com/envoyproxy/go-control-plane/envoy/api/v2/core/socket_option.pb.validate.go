// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/api/v2/core/socket_option.proto

package core

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on SocketOption with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SocketOption) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SocketOption with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SocketOptionMultiError, or
// nil if none found.
func (m *SocketOption) ValidateAll() error {
	return m.validate(true)
}

func (m *SocketOption) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Description

	// no validation rules for Level

	// no validation rules for Name

	if _, ok := SocketOption_SocketState_name[int32(m.GetState())]; !ok {
		err := SocketOptionValidationError{
			field:  "State",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	oneofValuePresent := false
	switch v := m.Value.(type) {
	case *SocketOption_IntValue:
		if v == nil {
			err := SocketOptionValidationError{
				field:  "Value",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofValuePresent = true
		// no validation rules for IntValue
	case *SocketOption_BufValue:
		if v == nil {
			err := SocketOptionValidationError{
				field:  "Value",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofValuePresent = true
		// no validation rules for BufValue
	default:
		_ = v // ensures v is used
	}
	if !oneofValuePresent {
		err := SocketOptionValidationError{
			field:  "Value",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SocketOptionMultiError(errors)
	}

	return nil
}

// SocketOptionMultiError is an error wrapping multiple validation errors
// returned by SocketOption.ValidateAll() if the designated constraints aren't met.
type SocketOptionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SocketOptionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SocketOptionMultiError) AllErrors() []error { return m }

// SocketOptionValidationError is the validation error returned by
// SocketOption.Validate if the designated constraints aren't met.
type SocketOptionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SocketOptionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SocketOptionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SocketOptionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SocketOptionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SocketOptionValidationError) ErrorName() string { return "SocketOptionValidationError" }

// Error satisfies the builtin error interface
func (e SocketOptionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSocketOption.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SocketOptionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SocketOptionValidationError{}
