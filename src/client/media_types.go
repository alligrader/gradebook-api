// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/alligrader/gradebook-api/designs
// --out=$(GOPATH)/src/github.com/alligrader/gradebook-api/src
// --version=v1.1.0-dirty
//
// API "GradebookAPI": Application Media Types
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

import (
	"github.com/goadesign/goa"
	"net/http"
	"unicode/utf8"
)

// GithubtokenMt media type (default view)
//
// Identifier: application/githubtoken.mt; view=default
type GithubtokenMt struct {
	Token *string `form:"token,omitempty" json:"token,omitempty" xml:"token,omitempty"`
}

// DecodeGithubtokenMt decodes the GithubtokenMt instance encoded in resp body.
func (c *Client) DecodeGithubtokenMt(resp *http.Response) (*GithubtokenMt, error) {
	var decoded GithubtokenMt
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// UserMt media type (Incoming view)
//
// Identifier: application/user.mt; view=Incoming
type UserMtIncoming struct {
	// User email
	Email    *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	Name     *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// Validate validates the UserMtIncoming media type instance.
func (mt *UserMtIncoming) Validate() (err error) {
	if mt.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, *mt.Email, goa.FormatEmail, err2))
		}
	}
	if mt.Email != nil {
		if utf8.RuneCountInString(*mt.Email) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.email`, *mt.Email, utf8.RuneCountInString(*mt.Email), 1, true))
		}
	}
	if mt.Name != nil {
		if utf8.RuneCountInString(*mt.Name) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, *mt.Name, utf8.RuneCountInString(*mt.Name), 1, true))
		}
	}
	if mt.Password != nil {
		if utf8.RuneCountInString(*mt.Password) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.password`, *mt.Password, utf8.RuneCountInString(*mt.Password), 2, true))
		}
	}
	return
}

// UserMt media type (default view)
//
// Identifier: application/user.mt; view=default
type UserMt struct {
	// User email
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	ID    *int    `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name  *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// Validate validates the UserMt media type instance.
func (mt *UserMt) Validate() (err error) {
	if mt.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, *mt.Email, goa.FormatEmail, err2))
		}
	}
	if mt.Email != nil {
		if utf8.RuneCountInString(*mt.Email) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.email`, *mt.Email, utf8.RuneCountInString(*mt.Email), 1, true))
		}
	}
	if mt.Name != nil {
		if utf8.RuneCountInString(*mt.Name) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, *mt.Name, utf8.RuneCountInString(*mt.Name), 1, true))
		}
	}
	return
}

// DecodeUserMtIncoming decodes the UserMtIncoming instance encoded in resp body.
func (c *Client) DecodeUserMtIncoming(resp *http.Response) (*UserMtIncoming, error) {
	var decoded UserMtIncoming
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeUserMt decodes the UserMt instance encoded in resp body.
func (c *Client) DecodeUserMt(resp *http.Response) (*UserMt, error) {
	var decoded UserMt
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
