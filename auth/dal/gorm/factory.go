package gorm

import (
	"github.com/smbody/anonym/config"
	"github.com/smbody/anonym/errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func Init() *Users {
	return initUsers(initGorm())
}

func initGorm() *gorm.DB {
	db, err := gorm.Open(dialect(config.Database(), config.DataSourceName()), &gorm.Config{})
	if err != nil {
		errors.DatabaseError(err)
	}
	return db

}

func dialect(database string, dsn string) gorm.Dialector {
	switch database {
	case "mysql":
		return mysql.Open(dsn)
	case "postgres":
		return postgres.Open(dsn)
	case "sqlserver":
		return sqlserver.Open(dsn)
	default:
		errors.UnknownDatabase(database)
		return nil
	}
}
