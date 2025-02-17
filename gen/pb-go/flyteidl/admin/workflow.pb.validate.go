// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: flyteidl/admin/workflow.proto

package admin

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _workflow_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on WorkflowCreateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *WorkflowCreateRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkflowCreateRequestValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetSpec()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkflowCreateRequestValidationError{
				field:  "Spec",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// WorkflowCreateRequestValidationError is the validation error returned by
// WorkflowCreateRequest.Validate if the designated constraints aren't met.
type WorkflowCreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkflowCreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkflowCreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkflowCreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkflowCreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkflowCreateRequestValidationError) ErrorName() string {
	return "WorkflowCreateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e WorkflowCreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkflowCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkflowCreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkflowCreateRequestValidationError{}

// Validate checks the field values on WorkflowCreateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *WorkflowCreateResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// WorkflowCreateResponseValidationError is the validation error returned by
// WorkflowCreateResponse.Validate if the designated constraints aren't met.
type WorkflowCreateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkflowCreateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkflowCreateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkflowCreateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkflowCreateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkflowCreateResponseValidationError) ErrorName() string {
	return "WorkflowCreateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e WorkflowCreateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkflowCreateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkflowCreateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkflowCreateResponseValidationError{}

// Validate checks the field values on Workflow with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Workflow) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkflowValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetClosure()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkflowValidationError{
				field:  "Closure",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// WorkflowValidationError is the validation error returned by
// Workflow.Validate if the designated constraints aren't met.
type WorkflowValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkflowValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkflowValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkflowValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkflowValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkflowValidationError) ErrorName() string { return "WorkflowValidationError" }

// Error satisfies the builtin error interface
func (e WorkflowValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkflow.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkflowValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkflowValidationError{}

// Validate checks the field values on WorkflowList with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *WorkflowList) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetWorkflows() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return WorkflowListValidationError{
					field:  fmt.Sprintf("Workflows[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Token

	return nil
}

// WorkflowListValidationError is the validation error returned by
// WorkflowList.Validate if the designated constraints aren't met.
type WorkflowListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkflowListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkflowListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkflowListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkflowListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkflowListValidationError) ErrorName() string { return "WorkflowListValidationError" }

// Error satisfies the builtin error interface
func (e WorkflowListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkflowList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkflowListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkflowListValidationError{}

// Validate checks the field values on WorkflowSpec with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *WorkflowSpec) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTemplate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkflowSpecValidationError{
				field:  "Template",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetSubWorkflows() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return WorkflowSpecValidationError{
					field:  fmt.Sprintf("SubWorkflows[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// WorkflowSpecValidationError is the validation error returned by
// WorkflowSpec.Validate if the designated constraints aren't met.
type WorkflowSpecValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkflowSpecValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkflowSpecValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkflowSpecValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkflowSpecValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkflowSpecValidationError) ErrorName() string { return "WorkflowSpecValidationError" }

// Error satisfies the builtin error interface
func (e WorkflowSpecValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkflowSpec.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkflowSpecValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkflowSpecValidationError{}

// Validate checks the field values on WorkflowClosure with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *WorkflowClosure) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetCompiledWorkflow()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkflowClosureValidationError{
				field:  "CompiledWorkflow",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkflowClosureValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// WorkflowClosureValidationError is the validation error returned by
// WorkflowClosure.Validate if the designated constraints aren't met.
type WorkflowClosureValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkflowClosureValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkflowClosureValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkflowClosureValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkflowClosureValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkflowClosureValidationError) ErrorName() string { return "WorkflowClosureValidationError" }

// Error satisfies the builtin error interface
func (e WorkflowClosureValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkflowClosure.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkflowClosureValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkflowClosureValidationError{}

// Validate checks the field values on WorkflowErrorExistsDifferentStructure
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *WorkflowErrorExistsDifferentStructure) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkflowErrorExistsDifferentStructureValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// WorkflowErrorExistsDifferentStructureValidationError is the validation error
// returned by WorkflowErrorExistsDifferentStructure.Validate if the
// designated constraints aren't met.
type WorkflowErrorExistsDifferentStructureValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkflowErrorExistsDifferentStructureValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkflowErrorExistsDifferentStructureValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkflowErrorExistsDifferentStructureValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkflowErrorExistsDifferentStructureValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkflowErrorExistsDifferentStructureValidationError) ErrorName() string {
	return "WorkflowErrorExistsDifferentStructureValidationError"
}

// Error satisfies the builtin error interface
func (e WorkflowErrorExistsDifferentStructureValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkflowErrorExistsDifferentStructure.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkflowErrorExistsDifferentStructureValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkflowErrorExistsDifferentStructureValidationError{}

// Validate checks the field values on WorkflowErrorExistsIdenticalStructure
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *WorkflowErrorExistsIdenticalStructure) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkflowErrorExistsIdenticalStructureValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// WorkflowErrorExistsIdenticalStructureValidationError is the validation error
// returned by WorkflowErrorExistsIdenticalStructure.Validate if the
// designated constraints aren't met.
type WorkflowErrorExistsIdenticalStructureValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkflowErrorExistsIdenticalStructureValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkflowErrorExistsIdenticalStructureValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkflowErrorExistsIdenticalStructureValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkflowErrorExistsIdenticalStructureValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkflowErrorExistsIdenticalStructureValidationError) ErrorName() string {
	return "WorkflowErrorExistsIdenticalStructureValidationError"
}

// Error satisfies the builtin error interface
func (e WorkflowErrorExistsIdenticalStructureValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkflowErrorExistsIdenticalStructure.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkflowErrorExistsIdenticalStructureValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkflowErrorExistsIdenticalStructureValidationError{}

// Validate checks the field values on CreateWorkflowFailureReason with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateWorkflowFailureReason) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Reason.(type) {

	case *CreateWorkflowFailureReason_ExistsDifferentStructure:

		if v, ok := interface{}(m.GetExistsDifferentStructure()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateWorkflowFailureReasonValidationError{
					field:  "ExistsDifferentStructure",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *CreateWorkflowFailureReason_ExistsIdenticalStructure:

		if v, ok := interface{}(m.GetExistsIdenticalStructure()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateWorkflowFailureReasonValidationError{
					field:  "ExistsIdenticalStructure",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CreateWorkflowFailureReasonValidationError is the validation error returned
// by CreateWorkflowFailureReason.Validate if the designated constraints
// aren't met.
type CreateWorkflowFailureReasonValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateWorkflowFailureReasonValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateWorkflowFailureReasonValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateWorkflowFailureReasonValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateWorkflowFailureReasonValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateWorkflowFailureReasonValidationError) ErrorName() string {
	return "CreateWorkflowFailureReasonValidationError"
}

// Error satisfies the builtin error interface
func (e CreateWorkflowFailureReasonValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateWorkflowFailureReason.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateWorkflowFailureReasonValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateWorkflowFailureReasonValidationError{}
