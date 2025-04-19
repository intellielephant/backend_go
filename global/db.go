/*
 * @Author: Will
 * @Date: 2022-11-16 10:53:33
 * @LastEditors: will-liu will-liu@live.com
 * @LastEditTime: 2023-02-16 11:26:33
 * @Description: 请填写简介
 */
package global

import (
	"fmt"
	//"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DBAI         *gorm.DB
	DBLittleFish *gorm.DB
	DBTour       *gorm.DB
)

func SetupDBAI() error {
	var err error
	userName := "root"
	password := "nupaeer"
	dbAddress := "47.120.22.151"
	//dbAddress := os.Getenv("mysqlAddress")
	dbName := "toolbox"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", userName, password, dbAddress, dbName)

	DBAI, err = gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN:        dsn,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err == nil {
		return nil
	} else {
		return err
	}
}

func SetupDBFish() error {
	var err error
	userName := "root"
	password := "nupaeer"
	dbAddress := "47.120.22.151"
	//dbAddress := os.Getenv("mysqlAddress")
	dbName := "fish"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", userName, password, dbAddress, dbName)

	DBLittleFish, err = gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN:        dsn,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err == nil {
		return nil
	} else {
		return err
	}
}

func SetupDBTour() error {
	var err error
	userName := "root"
	password := "nupaeer"
	dbAddress := "47.120.22.151"
	// dbAddress := os.Getenv("mysqlAddress")
	dbName := "travel"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", userName, password, dbAddress, dbName)

	DBTour, err = gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN:        dsn,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err == nil {
		return nil
	} else {
		return err
	}
}
