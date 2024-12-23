package drivers

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"vending_machine/config"
)

type MySQLClient struct {
	*sql.DB
}

func MySQLConn(opt config.DBConfig) (*MySQLClient, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", opt.DBUser, opt.DBPassword, opt.DBHost, opt.DBPort, opt.DBName)
	var db *sql.DB
	var err error

	for i := 0; i < 5; i++ {
		db, err = sql.Open("mysql", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				break
			}
		}
		log.Printf("Failed to connect to database. Retrying in 2 seconds... (%d/5)\n", i+1)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		return nil, err
	}

	log.Println("Database connected")

	return &MySQLClient{DB: db}, nil
}
