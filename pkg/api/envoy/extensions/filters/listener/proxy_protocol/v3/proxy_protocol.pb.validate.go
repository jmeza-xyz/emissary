//go:build !disable_pgv

// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/listener/proxy_protocol/v3/proxy_protocol.proto

package proxy_protocolv3

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

	v3 "github.com/emissary-ingress/emissary/v3/pkg/api/envoy/config/core/v3"
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

	_ = v3.ProxyProtocolConfig_Version(0)
)

// Validate checks the field values on ProxyProtocol with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ProxyProtocol) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProxyProtocol with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ProxyProtocolMultiError, or
// nil if none found.
func (m *ProxyProtocol) ValidateAll() error {
	return m.validate(true)
}

func (m *ProxyProtocol) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetRules() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ProxyProtocolValidationError{
						field:  fmt.Sprintf("Rules[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ProxyProtocolValidationError{
						field:  fmt.Sprintf("Rules[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProxyProtocolValidationError{
					field:  fmt.Sprintf("Rules[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for AllowRequestsWithoutProxyProtocol

	if all {
		switch v := interface{}(m.GetPassThroughTlvs()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProxyProtocolValidationError{
					field:  "PassThroughTlvs",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProxyProtocolValidationError{
					field:  "PassThroughTlvs",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPassThroughTlvs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProxyProtocolValidationError{
				field:  "PassThroughTlvs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ProxyProtocolMultiError(errors)
	}

	return nil
}

// ProxyProtocolMultiError is an error wrapping multiple validation errors
// returned by ProxyProtocol.ValidateAll() if the designated constraints
// aren't met.
type ProxyProtocolMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProxyProtocolMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProxyProtocolMultiError) AllErrors() []error { return m }

// ProxyProtocolValidationError is the validation error returned by
// ProxyProtocol.Validate if the designated constraints aren't met.
type ProxyProtocolValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProxyProtocolValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProxyProtocolValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProxyProtocolValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProxyProtocolValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProxyProtocolValidationError) ErrorName() string { return "ProxyProtocolValidationError" }

// Error satisfies the builtin error interface
func (e ProxyProtocolValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProxyProtocol.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProxyProtocolValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProxyProtocolValidationError{}

// Validate checks the field values on ProxyProtocol_KeyValuePair with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ProxyProtocol_KeyValuePair) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProxyProtocol_KeyValuePair with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProxyProtocol_KeyValuePairMultiError, or nil if none found.
func (m *ProxyProtocol_KeyValuePair) ValidateAll() error {
	return m.validate(true)
}

func (m *ProxyProtocol_KeyValuePair) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for MetadataNamespace

	if utf8.RuneCountInString(m.GetKey()) < 1 {
		err := ProxyProtocol_KeyValuePairValidationError{
			field:  "Key",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ProxyProtocol_KeyValuePairMultiError(errors)
	}

	return nil
}

// ProxyProtocol_KeyValuePairMultiError is an error wrapping multiple
// validation errors returned by ProxyProtocol_KeyValuePair.ValidateAll() if
// the designated constraints aren't met.
type ProxyProtocol_KeyValuePairMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProxyProtocol_KeyValuePairMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProxyProtocol_KeyValuePairMultiError) AllErrors() []error { return m }

// ProxyProtocol_KeyValuePairValidationError is the validation error returned
// by ProxyProtocol_KeyValuePair.Validate if the designated constraints aren't met.
type ProxyProtocol_KeyValuePairValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProxyProtocol_KeyValuePairValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProxyProtocol_KeyValuePairValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProxyProtocol_KeyValuePairValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProxyProtocol_KeyValuePairValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProxyProtocol_KeyValuePairValidationError) ErrorName() string {
	return "ProxyProtocol_KeyValuePairValidationError"
}

// Error satisfies the builtin error interface
func (e ProxyProtocol_KeyValuePairValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProxyProtocol_KeyValuePair.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProxyProtocol_KeyValuePairValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProxyProtocol_KeyValuePairValidationError{}

// Validate checks the field values on ProxyProtocol_Rule with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ProxyProtocol_Rule) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProxyProtocol_Rule with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProxyProtocol_RuleMultiError, or nil if none found.
func (m *ProxyProtocol_Rule) ValidateAll() error {
	return m.validate(true)
}

func (m *ProxyProtocol_Rule) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetTlvType() >= 256 {
		err := ProxyProtocol_RuleValidationError{
			field:  "TlvType",
			reason: "value must be less than 256",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetOnTlvPresent()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProxyProtocol_RuleValidationError{
					field:  "OnTlvPresent",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProxyProtocol_RuleValidationError{
					field:  "OnTlvPresent",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOnTlvPresent()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProxyProtocol_RuleValidationError{
				field:  "OnTlvPresent",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ProxyProtocol_RuleMultiError(errors)
	}

	return nil
}

// ProxyProtocol_RuleMultiError is an error wrapping multiple validation errors
// returned by ProxyProtocol_Rule.ValidateAll() if the designated constraints
// aren't met.
type ProxyProtocol_RuleMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProxyProtocol_RuleMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProxyProtocol_RuleMultiError) AllErrors() []error { return m }

// ProxyProtocol_RuleValidationError is the validation error returned by
// ProxyProtocol_Rule.Validate if the designated constraints aren't met.
type ProxyProtocol_RuleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProxyProtocol_RuleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProxyProtocol_RuleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProxyProtocol_RuleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProxyProtocol_RuleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProxyProtocol_RuleValidationError) ErrorName() string {
	return "ProxyProtocol_RuleValidationError"
}

// Error satisfies the builtin error interface
func (e ProxyProtocol_RuleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProxyProtocol_Rule.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProxyProtocol_RuleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProxyProtocol_RuleValidationError{}
