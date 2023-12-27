package schema

import (
	"fmt"
	"github.com/go-gormigrate/gormigrate/v2"
	"time"

	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"

	"test-git/migration/schema/versions"
)

func Run() (respErr error) {

	db, err := New(mysql.Open("test20:secret@tcp(localhost:3306)/test20?charset=utf8mb4&parseTime=true&loc=UTC&autocommit=true"), false)
	if err != nil {
		return err
	}

	defer func() {
		sqlDb, err := db.DB()
		if sqlDb != nil && err == nil {
			sqlDb.Close()
		}
	}()

	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				respErr = fmt.Errorf("%s", x)
			case error:
				respErr = x
			default:
				respErr = fmt.Errorf("unknown error: %+v", x)
			}
		}
	}()

	option := gormigrate.DefaultOptions
	option.TableName = "schema_migrations"

	//TODO: refactor version Version20230808080817, Version20230808080818
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID:       "20230808080810",
			Migrate:  versions.Version20230808080810,
			Rollback: versions.Rollback20230808080810,
		},
	})

	return m.Migrate()
}

func New(dialect gorm.Dialector, enableLog bool) (*gorm.DB, error) {
	orm, err := gorm.Open(dialect, &gorm.Config{})
	if nil != err {
		return nil, err
	}

	sqlDB, err := orm.DB()
	if nil != err {
		panic(err)
	}

	// TODO: use config for these values
	sqlDB.SetConnMaxLifetime(300 * time.Minute)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(15)

	return orm, nil
}
