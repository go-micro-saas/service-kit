// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: testdata/ping-service/api/ping-service/v1/enums/ping.enum.v1.proto

package enumv1

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

// Validate checks the field values on PingInitEnum with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PingInitEnum) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PingInitEnum with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PingInitEnumMultiError, or
// nil if none found.
func (m *PingInitEnum) ValidateAll() error {
	return m.validate(true)
}

func (m *PingInitEnum) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return PingInitEnumMultiError(errors)
	}

	return nil
}

// PingInitEnumMultiError is an error wrapping multiple validation errors
// returned by PingInitEnum.ValidateAll() if the designated constraints aren't met.
type PingInitEnumMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PingInitEnumMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PingInitEnumMultiError) AllErrors() []error { return m }

// PingInitEnumValidationError is the validation error returned by
// PingInitEnum.Validate if the designated constraints aren't met.
type PingInitEnumValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PingInitEnumValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PingInitEnumValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PingInitEnumValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PingInitEnumValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PingInitEnumValidationError) ErrorName() string { return "PingInitEnumValidationError" }

// Error satisfies the builtin error interface
func (e PingInitEnumValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPingInitEnum.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PingInitEnumValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PingInitEnumValidationError{}
