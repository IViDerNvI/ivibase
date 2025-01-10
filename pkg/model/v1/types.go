// Copyright 2024 ividernvi <ividernvi@outlook.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package v1

import (
	"time"
)

// ObjectMeta is a base model for all models.
type ObjectMeta struct {
	ID        uint       `json:"id,omitempty" gorm:"primary_key;auto_increment"`
	TypeName  string     `json:"type_name,omitempty" gorm:"column:type_name" validate:"required"`
	UUID      string     `json:"uuid,omitempty" gorm:"column:uuid" validate:"uuid;required"`
	CreatedAt time.Time  `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt time.Time  `json:"updated_at,omitempty" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"column:deleted_at" sql:"index"`
}

// ListMeta is a model for list meta.
type ListMeta struct {
	Total int `json:"total"`
}
