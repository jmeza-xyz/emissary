//go:build !disable_pgv

// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/network/sni_dynamic_forward_proxy/v3/sni_dynamic_forward_proxy.proto

package sni_dynamic_forward_proxyv3

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

// Validate checks the field values on FilterConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FilterConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FilterConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FilterConfigMultiError, or
// nil if none found.
func (m *FilterConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *FilterConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetDnsCacheConfig() == nil {
		err := FilterConfigValidationError{
			field:  "DnsCacheConfig",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetDnsCacheConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, FilterConfigValidationError{
					field:  "DnsCacheConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, FilterConfigValidationError{
					field:  "DnsCacheConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDnsCacheConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FilterConfigValidationError{
				field:  "DnsCacheConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch v := m.PortSpecifier.(type) {
	case *FilterConfig_PortValue:
		if v == nil {
			err := FilterConfigValidationError{
				field:  "PortSpecifier",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if val := m.GetPortValue(); val <= 0 || val > 65535 {
			err := FilterConfigValidationError{
				field:  "PortValue",
				reason: "value must be inside range (0, 65535]",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return FilterConfigMultiError(errors)
	}

	return nil
}

// FilterConfigMultiError is an error wrapping multiple validation errors
// returned by FilterConfig.ValidateAll() if the designated constraints aren't met.
type FilterConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FilterConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FilterConfigMultiError) AllErrors() []error { return m }

// FilterConfigValidationError is the validation error returned by
// FilterConfig.Validate if the designated constraints aren't met.
type FilterConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FilterConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FilterConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FilterConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FilterConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FilterConfigValidationError) ErrorName() string { return "FilterConfigValidationError" }

// Error satisfies the builtin error interface
func (e FilterConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFilterConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FilterConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FilterConfigValidationError{}
