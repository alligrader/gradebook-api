// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/alligrader/gradebook-api/designs
// --out=$(GOPATH)/src/github.com/alligrader/gradebook-api/src
// --version=v1.1.0-dirty
//
// API "GradebookAPI": Model Helpers
//
// The content of this file is auto-generated, DO NOT MODIFY

package models

import (
	"github.com/alligrader/gradebook-api/src/app"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"time"
)

// MediaType Retrieval Functions

// ListUserMtIncoming returns an array of view: Incoming.
func (m *UserDB) ListUserMtIncoming(ctx context.Context) []*app.UserMtIncoming {
	defer goa.MeasureSince([]string{"goa", "db", "userMt", "listuserMtincoming"}, time.Now())

	var native []*User
	var objs []*app.UserMtIncoming
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing User", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.UserToUserMtIncoming())
	}

	return objs
}

// UserToUserMtIncoming loads a User and builds the Incoming view of media type UserMt.
func (m *User) UserToUserMtIncoming() *app.UserMtIncoming {
	user := &app.UserMtIncoming{}
	user.Email = m.Email
	user.Name = m.Name
	user.Password = m.Password

	return user
}

// OneUserMtIncoming loads a User and builds the Incoming view of media type UserMt.
func (m *UserDB) OneUserMtIncoming(ctx context.Context, id int) (*app.UserMtIncoming, error) {
	defer goa.MeasureSince([]string{"goa", "db", "userMt", "oneuserMtincoming"}, time.Now())

	var native User
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting User", "error", err.Error())
		return nil, err
	}

	view := *native.UserToUserMtIncoming()
	return &view, err
}

// MediaType Retrieval Functions

// ListUserMtOnLogin returns an array of view: OnLogin.
func (m *UserDB) ListUserMtOnLogin(ctx context.Context) []*app.UserMtOnLogin {
	defer goa.MeasureSince([]string{"goa", "db", "userMt", "listuserMtonLogin"}, time.Now())

	var native []*User
	var objs []*app.UserMtOnLogin
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing User", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.UserToUserMtOnLogin())
	}

	return objs
}

// UserToUserMtOnLogin loads a User and builds the OnLogin view of media type UserMt.
func (m *User) UserToUserMtOnLogin() *app.UserMtOnLogin {
	user := &app.UserMtOnLogin{}
	user.Email = m.Email
	user.ID = &m.ID
	user.Name = m.Name

	return user
}

// OneUserMtOnLogin loads a User and builds the OnLogin view of media type UserMt.
func (m *UserDB) OneUserMtOnLogin(ctx context.Context, id int) (*app.UserMtOnLogin, error) {
	defer goa.MeasureSince([]string{"goa", "db", "userMt", "oneuserMtonLogin"}, time.Now())

	var native User
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting User", "error", err.Error())
		return nil, err
	}

	view := *native.UserToUserMtOnLogin()
	return &view, err
}

// MediaType Retrieval Functions

// ListUserMt returns an array of view: default.
func (m *UserDB) ListUserMt(ctx context.Context) []*app.UserMt {
	defer goa.MeasureSince([]string{"goa", "db", "userMt", "listuserMt"}, time.Now())

	var native []*User
	var objs []*app.UserMt
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing User", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.UserToUserMt())
	}

	return objs
}

// UserToUserMt loads a User and builds the default view of media type UserMt.
func (m *User) UserToUserMt() *app.UserMt {
	user := &app.UserMt{}
	user.Email = m.Email
	user.ID = &m.ID
	user.Name = m.Name

	return user
}

// OneUserMt loads a User and builds the default view of media type UserMt.
func (m *UserDB) OneUserMt(ctx context.Context, id int) (*app.UserMt, error) {
	defer goa.MeasureSince([]string{"goa", "db", "userMt", "oneuserMt"}, time.Now())

	var native User
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting User", "error", err.Error())
		return nil, err
	}

	view := *native.UserToUserMt()
	return &view, err
}
