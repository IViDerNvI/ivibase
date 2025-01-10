// Copyright 2024 ividernvi <ividernvi@outlook.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package idutil

import (
	"github.com/google/uuid"
)

func UUID() string {
	return uuid.New().String()
}

func UUIDBytes() []byte {
	return []byte(uuid.New().String())
}
