// Copyright 2024 ividernvi <ividernvi@outlook.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package v1

import (
	"time"

	"github.com/ividernvi/ivibase/pkg/util/idutil"
	"gorm.io/gorm"
)

// User is a model for user.
type User struct {
	// Base information
	ObjectMeta `json:",inline"`

	// Account information
	Status   string `json:"status" gorm:"column:status" validate:"oneof=1 2 3 4 5"`
	Username string `json:"username" gorm:"column:username" validate:"required"`
	Password string `json:"password" gorm:"column:password" validate:"required"`

	// Login information
	LoginAt time.Time `json:"login_at" gorm:"column:login_at"`

	// Personal information
	FirstName string `json:"first_name" gorm:"column:first_name" validate:"min=2;max=32"`
	LastName  string `json:"last_name" gorm:"column:last_name" validate:"min=2;max=32"`
	Profile   string `json:"profile" gorm:"column:profile" validate:"max=255"`
	Email     string `json:"email" gorm:"column:email" validate:"email;required"`
	Phone     string `json:"phone" gorm:"column:phone" validate:"phone"`
	Sexy      string `json:"sexy" gorm:"column:sexy" validate:"oneof=0 1 2"`
	Locale    string `json:"locale" gorm:"column:locale"`

	// Other information
}

// UserList is a model for user list.
type UserList struct {
	ListMeta `json:",inline"`
	Items    []User `json:"items"`
}

// TableName returns the table name.
func (u *User) TableName() string {
	return "user"
}

// BeforeCreate will set a UUID rather than numeric ID.
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// ObjectMeta fields change.
	u.UUID = idutil.UUID()
	u.TypeName = u.TableName()

	// User fields change.
	u.TypeName = u.TableName()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	return
}

// BeforeUpdate will set the updated time.
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}

// BeforeDelete will set the deleted time.
func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	now := time.Now()
	u.DeletedAt = &now
	return
}

// AfterCreate will reset the ID to 0.
// Preventing the ID from being returned to the client.
func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	u.ID = 0
	u.Password = ""
	return
}

// AfterUpdate will do nothing for now.
func (u *User) AfterUpdate(tx *gorm.DB) (err error) {
	return
}

// AfterDelete will do nothing for now.
func (u *User) AfterDelete(tx *gorm.DB) (err error) {
	return
}

// AfterFind will do nothing for now.
func (u *User) AfterFind(tx *gorm.DB) (err error) {
	return
}
