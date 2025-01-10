// Copyright 2024 ividernvi <ividernvi@outlook.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package mysql

import (
	"fmt"

	"github.com/ividernvi/ivibase/pkg/config/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New(options options.MysqlOptions) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", options.Username, options.Password, options.Host, options.Port, options.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: options.Logger,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(options.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(options.ConnMaxLifetime)
	sqlDB.SetMaxIdleConns(options.MaxIdleConns)

	return db, nil
}
