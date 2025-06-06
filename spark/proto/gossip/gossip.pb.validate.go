// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: gossip.proto

package gossip

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

// Validate checks the field values on GossipMessage with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GossipMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GossipMessage with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GossipMessageMultiError, or
// nil if none found.
func (m *GossipMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *GossipMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for MessageId

	switch v := m.Message.(type) {
	case *GossipMessage_CancelTransfer:
		if v == nil {
			err := GossipMessageValidationError{
				field:  "Message",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetCancelTransfer()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GossipMessageValidationError{
						field:  "CancelTransfer",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GossipMessageValidationError{
						field:  "CancelTransfer",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCancelTransfer()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GossipMessageValidationError{
					field:  "CancelTransfer",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return GossipMessageMultiError(errors)
	}

	return nil
}

// GossipMessageMultiError is an error wrapping multiple validation errors
// returned by GossipMessage.ValidateAll() if the designated constraints
// aren't met.
type GossipMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GossipMessageMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GossipMessageMultiError) AllErrors() []error { return m }

// GossipMessageValidationError is the validation error returned by
// GossipMessage.Validate if the designated constraints aren't met.
type GossipMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GossipMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GossipMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GossipMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GossipMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GossipMessageValidationError) ErrorName() string { return "GossipMessageValidationError" }

// Error satisfies the builtin error interface
func (e GossipMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGossipMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GossipMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GossipMessageValidationError{}

// Validate checks the field values on GossipMessageCancelTransfer with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GossipMessageCancelTransfer) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GossipMessageCancelTransfer with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GossipMessageCancelTransferMultiError, or nil if none found.
func (m *GossipMessageCancelTransfer) ValidateAll() error {
	return m.validate(true)
}

func (m *GossipMessageCancelTransfer) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TransferId

	if len(errors) > 0 {
		return GossipMessageCancelTransferMultiError(errors)
	}

	return nil
}

// GossipMessageCancelTransferMultiError is an error wrapping multiple
// validation errors returned by GossipMessageCancelTransfer.ValidateAll() if
// the designated constraints aren't met.
type GossipMessageCancelTransferMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GossipMessageCancelTransferMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GossipMessageCancelTransferMultiError) AllErrors() []error { return m }

// GossipMessageCancelTransferValidationError is the validation error returned
// by GossipMessageCancelTransfer.Validate if the designated constraints
// aren't met.
type GossipMessageCancelTransferValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GossipMessageCancelTransferValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GossipMessageCancelTransferValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GossipMessageCancelTransferValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GossipMessageCancelTransferValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GossipMessageCancelTransferValidationError) ErrorName() string {
	return "GossipMessageCancelTransferValidationError"
}

// Error satisfies the builtin error interface
func (e GossipMessageCancelTransferValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGossipMessageCancelTransfer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GossipMessageCancelTransferValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GossipMessageCancelTransferValidationError{}
