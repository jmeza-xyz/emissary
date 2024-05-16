//go:build !disable_pgv

// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/network/thrift_proxy/v3/thrift_proxy.proto

package thrift_proxyv3

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

// Validate checks the field values on Trds with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Trds) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Trds with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TrdsMultiError, or nil if none found.
func (m *Trds) ValidateAll() error {
	return m.validate(true)
}

func (m *Trds) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetConfigSource() == nil {
		err := TrdsValidationError{
			field:  "ConfigSource",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetConfigSource()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TrdsValidationError{
					field:  "ConfigSource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TrdsValidationError{
					field:  "ConfigSource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetConfigSource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TrdsValidationError{
				field:  "ConfigSource",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for RouteConfigName

	if len(errors) > 0 {
		return TrdsMultiError(errors)
	}

	return nil
}

// TrdsMultiError is an error wrapping multiple validation errors returned by
// Trds.ValidateAll() if the designated constraints aren't met.
type TrdsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TrdsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TrdsMultiError) AllErrors() []error { return m }

// TrdsValidationError is the validation error returned by Trds.Validate if the
// designated constraints aren't met.
type TrdsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TrdsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TrdsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TrdsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TrdsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TrdsValidationError) ErrorName() string { return "TrdsValidationError" }

// Error satisfies the builtin error interface
func (e TrdsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTrds.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TrdsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TrdsValidationError{}

// Validate checks the field values on ThriftProxy with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ThriftProxy) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ThriftProxy with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ThriftProxyMultiError, or
// nil if none found.
func (m *ThriftProxy) ValidateAll() error {
	return m.validate(true)
}

func (m *ThriftProxy) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := TransportType_name[int32(m.GetTransport())]; !ok {
		err := ThriftProxyValidationError{
			field:  "Transport",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := ProtocolType_name[int32(m.GetProtocol())]; !ok {
		err := ThriftProxyValidationError{
			field:  "Protocol",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetStatPrefix()) < 1 {
		err := ThriftProxyValidationError{
			field:  "StatPrefix",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetRouteConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ThriftProxyValidationError{
					field:  "RouteConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ThriftProxyValidationError{
					field:  "RouteConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRouteConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ThriftProxyValidationError{
				field:  "RouteConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTrds()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ThriftProxyValidationError{
					field:  "Trds",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ThriftProxyValidationError{
					field:  "Trds",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTrds()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ThriftProxyValidationError{
				field:  "Trds",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetThriftFilters() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ThriftProxyValidationError{
						field:  fmt.Sprintf("ThriftFilters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ThriftProxyValidationError{
						field:  fmt.Sprintf("ThriftFilters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ThriftProxyValidationError{
					field:  fmt.Sprintf("ThriftFilters[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for PayloadPassthrough

	if all {
		switch v := interface{}(m.GetMaxRequestsPerConnection()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ThriftProxyValidationError{
					field:  "MaxRequestsPerConnection",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ThriftProxyValidationError{
					field:  "MaxRequestsPerConnection",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMaxRequestsPerConnection()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ThriftProxyValidationError{
				field:  "MaxRequestsPerConnection",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetAccessLog() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ThriftProxyValidationError{
						field:  fmt.Sprintf("AccessLog[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ThriftProxyValidationError{
						field:  fmt.Sprintf("AccessLog[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ThriftProxyValidationError{
					field:  fmt.Sprintf("AccessLog[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for HeaderKeysPreserveCase

	if len(errors) > 0 {
		return ThriftProxyMultiError(errors)
	}

	return nil
}

// ThriftProxyMultiError is an error wrapping multiple validation errors
// returned by ThriftProxy.ValidateAll() if the designated constraints aren't met.
type ThriftProxyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ThriftProxyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ThriftProxyMultiError) AllErrors() []error { return m }

// ThriftProxyValidationError is the validation error returned by
// ThriftProxy.Validate if the designated constraints aren't met.
type ThriftProxyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ThriftProxyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ThriftProxyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ThriftProxyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ThriftProxyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ThriftProxyValidationError) ErrorName() string { return "ThriftProxyValidationError" }

// Error satisfies the builtin error interface
func (e ThriftProxyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sThriftProxy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ThriftProxyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ThriftProxyValidationError{}

// Validate checks the field values on ThriftFilter with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ThriftFilter) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ThriftFilter with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ThriftFilterMultiError, or
// nil if none found.
func (m *ThriftFilter) ValidateAll() error {
	return m.validate(true)
}

func (m *ThriftFilter) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := ThriftFilterValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	switch v := m.ConfigType.(type) {
	case *ThriftFilter_TypedConfig:
		if v == nil {
			err := ThriftFilterValidationError{
				field:  "ConfigType",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetTypedConfig()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ThriftFilterValidationError{
						field:  "TypedConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ThriftFilterValidationError{
						field:  "TypedConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetTypedConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ThriftFilterValidationError{
					field:  "TypedConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return ThriftFilterMultiError(errors)
	}

	return nil
}

// ThriftFilterMultiError is an error wrapping multiple validation errors
// returned by ThriftFilter.ValidateAll() if the designated constraints aren't met.
type ThriftFilterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ThriftFilterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ThriftFilterMultiError) AllErrors() []error { return m }

// ThriftFilterValidationError is the validation error returned by
// ThriftFilter.Validate if the designated constraints aren't met.
type ThriftFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ThriftFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ThriftFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ThriftFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ThriftFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ThriftFilterValidationError) ErrorName() string { return "ThriftFilterValidationError" }

// Error satisfies the builtin error interface
func (e ThriftFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sThriftFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ThriftFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ThriftFilterValidationError{}

// Validate checks the field values on ThriftProtocolOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ThriftProtocolOptions) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ThriftProtocolOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ThriftProtocolOptionsMultiError, or nil if none found.
func (m *ThriftProtocolOptions) ValidateAll() error {
	return m.validate(true)
}

func (m *ThriftProtocolOptions) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := TransportType_name[int32(m.GetTransport())]; !ok {
		err := ThriftProtocolOptionsValidationError{
			field:  "Transport",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := ProtocolType_name[int32(m.GetProtocol())]; !ok {
		err := ThriftProtocolOptionsValidationError{
			field:  "Protocol",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ThriftProtocolOptionsMultiError(errors)
	}

	return nil
}

// ThriftProtocolOptionsMultiError is an error wrapping multiple validation
// errors returned by ThriftProtocolOptions.ValidateAll() if the designated
// constraints aren't met.
type ThriftProtocolOptionsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ThriftProtocolOptionsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ThriftProtocolOptionsMultiError) AllErrors() []error { return m }

// ThriftProtocolOptionsValidationError is the validation error returned by
// ThriftProtocolOptions.Validate if the designated constraints aren't met.
type ThriftProtocolOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ThriftProtocolOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ThriftProtocolOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ThriftProtocolOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ThriftProtocolOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ThriftProtocolOptionsValidationError) ErrorName() string {
	return "ThriftProtocolOptionsValidationError"
}

// Error satisfies the builtin error interface
func (e ThriftProtocolOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sThriftProtocolOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ThriftProtocolOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ThriftProtocolOptionsValidationError{}
