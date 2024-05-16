//go:build !disable_pgv

// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/path/match/uri_template/v3/uri_template_match.proto

package uri_templatev3

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

// Validate checks the field values on UriTemplateMatchConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UriTemplateMatchConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UriTemplateMatchConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UriTemplateMatchConfigMultiError, or nil if none found.
func (m *UriTemplateMatchConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *UriTemplateMatchConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetPathTemplate()); l < 1 || l > 256 {
		err := UriTemplateMatchConfigValidationError{
			field:  "PathTemplate",
			reason: "value length must be between 1 and 256 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UriTemplateMatchConfigMultiError(errors)
	}

	return nil
}

// UriTemplateMatchConfigMultiError is an error wrapping multiple validation
// errors returned by UriTemplateMatchConfig.ValidateAll() if the designated
// constraints aren't met.
type UriTemplateMatchConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UriTemplateMatchConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UriTemplateMatchConfigMultiError) AllErrors() []error { return m }

// UriTemplateMatchConfigValidationError is the validation error returned by
// UriTemplateMatchConfig.Validate if the designated constraints aren't met.
type UriTemplateMatchConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UriTemplateMatchConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UriTemplateMatchConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UriTemplateMatchConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UriTemplateMatchConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UriTemplateMatchConfigValidationError) ErrorName() string {
	return "UriTemplateMatchConfigValidationError"
}

// Error satisfies the builtin error interface
func (e UriTemplateMatchConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUriTemplateMatchConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UriTemplateMatchConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UriTemplateMatchConfigValidationError{}
