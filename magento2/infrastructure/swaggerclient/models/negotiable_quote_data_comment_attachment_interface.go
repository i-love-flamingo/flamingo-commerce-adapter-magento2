// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NegotiableQuoteDataCommentAttachmentInterface Interface for quote comment attachment.
// swagger:model negotiable-quote-data-comment-attachment-interface
type NegotiableQuoteDataCommentAttachmentInterface struct {

	// Attachment ID.
	// Required: true
	AttachmentID *int64 `json:"attachment_id"`

	// Comment ID.
	// Required: true
	CommentID *int64 `json:"comment_id"`

	// extension attributes
	ExtensionAttributes NegotiableQuoteDataCommentAttachmentExtensionInterface `json:"extension_attributes,omitempty"`

	// File name.
	// Required: true
	FileName *string `json:"file_name"`

	// File path.
	// Required: true
	FilePath *string `json:"file_path"`

	// File type.
	// Required: true
	FileType *string `json:"file_type"`
}

// Validate validates this negotiable quote data comment attachment interface
func (m *NegotiableQuoteDataCommentAttachmentInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachmentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NegotiableQuoteDataCommentAttachmentInterface) validateAttachmentID(formats strfmt.Registry) error {

	if err := validate.Required("attachment_id", "body", m.AttachmentID); err != nil {
		return err
	}

	return nil
}

func (m *NegotiableQuoteDataCommentAttachmentInterface) validateCommentID(formats strfmt.Registry) error {

	if err := validate.Required("comment_id", "body", m.CommentID); err != nil {
		return err
	}

	return nil
}

func (m *NegotiableQuoteDataCommentAttachmentInterface) validateFileName(formats strfmt.Registry) error {

	if err := validate.Required("file_name", "body", m.FileName); err != nil {
		return err
	}

	return nil
}

func (m *NegotiableQuoteDataCommentAttachmentInterface) validateFilePath(formats strfmt.Registry) error {

	if err := validate.Required("file_path", "body", m.FilePath); err != nil {
		return err
	}

	return nil
}

func (m *NegotiableQuoteDataCommentAttachmentInterface) validateFileType(formats strfmt.Registry) error {

	if err := validate.Required("file_type", "body", m.FileType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NegotiableQuoteDataCommentAttachmentInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NegotiableQuoteDataCommentAttachmentInterface) UnmarshalBinary(b []byte) error {
	var res NegotiableQuoteDataCommentAttachmentInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
