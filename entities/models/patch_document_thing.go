//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PatchDocumentThing Either a JSONPatch document as defined by RFC 6902 (from, op, path, value), or a merge document (RFC 7396).
// swagger:model PatchDocumentThing
type PatchDocumentThing struct {

	// A string containing a JSON Pointer value.
	From string `json:"from,omitempty"`

	// merge
	Merge *Thing `json:"merge,omitempty"`

	// The operation to be performed.
	// Required: true
	// Enum: [add remove replace move copy test]
	Op *string `json:"op"`

	// A JSON-Pointer.
	// Required: true
	Path *string `json:"path"`

	// The value to be used within the operations.
	Value interface{} `json:"value,omitempty"`
}

// Validate validates this patch document thing
func (m *PatchDocumentThing) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMerge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchDocumentThing) validateMerge(formats strfmt.Registry) error {

	if swag.IsZero(m.Merge) { // not required
		return nil
	}

	if m.Merge != nil {
		if err := m.Merge.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("merge")
			}
			return err
		}
	}

	return nil
}

var patchDocumentThingTypeOpPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["add","remove","replace","move","copy","test"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchDocumentThingTypeOpPropEnum = append(patchDocumentThingTypeOpPropEnum, v)
	}
}

const (

	// PatchDocumentThingOpAdd captures enum value "add"
	PatchDocumentThingOpAdd string = "add"

	// PatchDocumentThingOpRemove captures enum value "remove"
	PatchDocumentThingOpRemove string = "remove"

	// PatchDocumentThingOpReplace captures enum value "replace"
	PatchDocumentThingOpReplace string = "replace"

	// PatchDocumentThingOpMove captures enum value "move"
	PatchDocumentThingOpMove string = "move"

	// PatchDocumentThingOpCopy captures enum value "copy"
	PatchDocumentThingOpCopy string = "copy"

	// PatchDocumentThingOpTest captures enum value "test"
	PatchDocumentThingOpTest string = "test"
)

// prop value enum
func (m *PatchDocumentThing) validateOpEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, patchDocumentThingTypeOpPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PatchDocumentThing) validateOp(formats strfmt.Registry) error {

	if err := validate.Required("op", "body", m.Op); err != nil {
		return err
	}

	// value enum
	if err := m.validateOpEnum("op", "body", *m.Op); err != nil {
		return err
	}

	return nil
}

func (m *PatchDocumentThing) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchDocumentThing) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchDocumentThing) UnmarshalBinary(b []byte) error {
	var res PatchDocumentThing
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}