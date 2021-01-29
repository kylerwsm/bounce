package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dsn string
	db  *gorm.DB
)

func dieIfAbsent(boolean bool, envKey string) {
	if boolean {
		log.Fatalf("Required environment variable \"%s\" is not present", envKey)
	}
}

func init() {
	const (
		Host     = "PG_HOST"
		User     = "PG_USER"
		Password = "PG_PASSWORD"
		DBName   = "PG_DBNAME"
		Port     = "PG_PORT"
		SslMode  = "PG_SSL_MODE"
		TimeZone = "PG_TIMEZONE"
	)

	var (
		host     string
		user     string
		password string
		dbname   string
		port     string
		sslmode  string
		timezone string
		ok       bool
	)

	host, ok = os.LookupEnv(Host)
	dieIfAbsent(ok, Host)
	user, ok = os.LookupEnv(User)
	dieIfAbsent(ok, User)
	password, ok = os.LookupEnv(Password)
	dieIfAbsent(ok, Password)
	dbname, ok = os.LookupEnv(DBName)
	dieIfAbsent(ok, DBName)
	port, ok = os.LookupEnv(Port)
	dieIfAbsent(ok, Port)
	sslmode, ok = os.LookupEnv(SslMode)
	dieIfAbsent(ok, SslMode)
	timezone, ok = os.LookupEnv(TimeZone)
	dieIfAbsent(ok, TimeZone)

	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone)
}

// SetDSN allows a default dsn manual override.
func SetDSN(customDSN string) {
	dsn = customDSN
}

// GetDatabase returns the current database.
func GetDatabase() *gorm.DB {
	var err error
	if db == nil {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalln(err)
		}
	}
	return db
}
