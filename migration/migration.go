package migration

import (
	"fmt"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"os"
	"test/connection"
	"test/migration/versions"
)

func Migrate(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID:      "20221110000009",
			Migrate: versions.Version20221110000009,
		},
	})

	return m.Migrate()
}

func MigrateDB() {
	fmt.Println("mysql connection established")
	err := Migrate(connection.Conn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("migration successful")
	os.Exit(0)
}
