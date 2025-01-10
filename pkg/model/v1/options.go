// Copyright 2024 ividernvi <ividernvi@outlook.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package v1

type OptionMeta struct {
	Kind       string `json:"kind,omitempty"`
	APIVersion string `json:"apiVersion,omitempty"`
}

type ListOptions struct {
	OptionMeta     `json:",inline"`
	Selector       string `json:"selector,omitempty" form:"selector"`
	TimeoutSeconds int64  `json:"timeoutSeconds,omitempty"`
	Offset         int64  `json:"offset,omitempty" form:"offset"`
	Limit          int64  `json:"limit,omitempty" form:"limit"`
}

type GetOptions struct {
	OptionMeta `json:",inline"`
}

type DeleteOptions struct {
	OptionMeta `json:",inline"`
}

type CreateOptions struct {
	OptionMeta `json:",inline"`
}

type UpdateOptions struct {
	OptionMeta `json:",inline"`
}
