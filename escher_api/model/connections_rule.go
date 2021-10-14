// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	strfmt "github.com/go-openapi/strfmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConnectionsRule A rule that states what apps are allowed to communicate with each other.
// swagger:model ConnectionsRule
type ConnectionsRule struct {

	// action
	Action ConnectionRuleAction `json:"action,omitempty"`

	destinationField ConnectionRulePart

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	GroupName string `json:"groupName,omitempty"`

	sourceField ConnectionRulePart

	// status
	// Enum: [ENABLED DISABLED DELETED]
	Status string `json:"status,omitempty"`
}

// Destination gets the destination of this base type
func (m *ConnectionsRule) Destination() ConnectionRulePart {
	return m.destinationField
}

// SetDestination sets the destination of this base type
func (m *ConnectionsRule) SetDestination(val ConnectionRulePart) {
	m.destinationField = val
}

// Source gets the source of this base type
func (m *ConnectionsRule) Source() ConnectionRulePart {
	return m.sourceField
}

// SetSource sets the source of this base type
func (m *ConnectionsRule) SetSource(val ConnectionRulePart) {
	m.sourceField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *ConnectionsRule) UnmarshalJSON(raw []byte) error {
	var data struct {
		Action ConnectionRuleAction `json:"action,omitempty"`

		Destination json.RawMessage `json:"destination,omitempty"`

		ID strfmt.UUID `json:"id,omitempty"`

		Name string `json:"name,omitempty"`

		GroupName string `json:"groupName,omitempty"`

		Source json.RawMessage `json:"source,omitempty"`

		Status string `json:"status,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var propDestination ConnectionRulePart
	if string(data.Destination) != "null" {
		destination, err := UnmarshalConnectionRulePart(bytes.NewBuffer(data.Destination), runtime.JSONConsumer())
		if err != nil && err != io.EOF {
			return err
		}
		propDestination = destination
	}

	var propSource ConnectionRulePart
	if string(data.Source) != "null" {
		source, err := UnmarshalConnectionRulePart(bytes.NewBuffer(data.Source), runtime.JSONConsumer())
		if err != nil && err != io.EOF {
			return err
		}
		propSource = source
	}

	var result ConnectionsRule

	// action
	result.Action = data.Action

	// destination
	result.destinationField = propDestination

	// id
	result.ID = data.ID

	// name
	result.Name = data.Name

	result.GroupName = data.GroupName

	// source
	result.sourceField = propSource

	// status
	result.Status = data.Status

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m ConnectionsRule) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		Action ConnectionRuleAction `json:"action,omitempty"`

		ID strfmt.UUID `json:"id,omitempty"`

		Name string `json:"name,omitempty"`

		GroupName string `json:"groupName,omitempty"`

		Status string `json:"status,omitempty"`
	}{

		Action: m.Action,

		ID: m.ID,

		Name:      m.Name,
		GroupName: m.GroupName,

		Status: m.Status,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Destination ConnectionRulePart `json:"destination,omitempty"`

		Source ConnectionRulePart `json:"source,omitempty"`
	}{

		Destination: m.destinationField,

		Source: m.sourceField,
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this connections rule
func (m *ConnectionsRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConnectionsRule) validateAction(formats strfmt.Registry) error {

	if swag.IsZero(m.Action) { // not required
		return nil
	}

	if err := m.Action.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("action")
		}
		return err
	}

	return nil
}

func (m *ConnectionsRule) validateDestination(formats strfmt.Registry) error {

	if swag.IsZero(m.Destination()) { // not required
		return nil
	}

	if err := m.Destination().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("destination")
		}
		return err
	}

	return nil
}

func (m *ConnectionsRule) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConnectionsRule) validateSource(formats strfmt.Registry) error {

	if swag.IsZero(m.Source()) { // not required
		return nil
	}

	if err := m.Source().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("source")
		}
		return err
	}

	return nil
}

var connectionsRuleTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENABLED","DISABLED","DELETED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		connectionsRuleTypeStatusPropEnum = append(connectionsRuleTypeStatusPropEnum, v)
	}
}

const (

	// ConnectionsRuleStatusENABLED captures enum value "ENABLED"
	ConnectionsRuleStatusENABLED string = "ENABLED"

	// ConnectionsRuleStatusDISABLED captures enum value "DISABLED"
	ConnectionsRuleStatusDISABLED string = "DISABLED"

	// ConnectionsRuleStatusDELETED captures enum value "DELETED"
	ConnectionsRuleStatusDELETED string = "DELETED"
)

// prop value enum
func (m *ConnectionsRule) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, connectionsRuleTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ConnectionsRule) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConnectionsRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConnectionsRule) UnmarshalBinary(b []byte) error {
	var res ConnectionsRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
