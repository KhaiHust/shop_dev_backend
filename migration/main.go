package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"shop_dev/datasource/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DatabaseName,
	)
	db, _ := sql.Open("mysql", url)
	driver, _ := mysql.WithInstance(db, &mysql.Config{})

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)
	fmt.Printf("err: %v \n", err)

	m.Steps(2)
}
