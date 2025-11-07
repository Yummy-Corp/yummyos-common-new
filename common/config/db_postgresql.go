package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectWithConnector() (*sql.DB, error) {
	mustGetenv := func(k string) string {
		v := os.Getenv(k)
		if v == "" {
			log.Fatalf("Warning: %s environment variable not set.\n", k)
		}
		return v
	}
	// Note: Saving credentials in environment variables is convenient, but not
	// secure - consider a more secure solution such as
	// Cloud Secret Manager (https://cloud.google.com/secret-manager) to help
	// keep secrets safe.
	var (
		// Either a DB_USER or a DB_IAM_USER should be defined. If both are
		// defined, DB_IAM_USER takes precedence.
		dbUser                 = os.Getenv("DB_USER")      // e.g. 'my-db-user'
		dbPwd                  = mustGetenv("DB_PASS")     // e.g. 'my-db-password'
		dbName                 = mustGetenv("DB_NAME")     // e.g. 'my-database'
		instanceConnectionName = mustGetenv("DB_INSTANCE") // e.g. 'project:region:instance'
	)
	if dbUser == "" {
		log.Fatal("DB_USER undefined")
	}

	dsn := fmt.Sprintf("user=%s password=%s database=%s", dbUser, dbPwd, dbName)
	config, err := pgx.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}
	config.DialFunc = func(ctx context.Context, network, instance string) (net.Conn, error) {
		// Use the Cloud SQL connector to handle connecting to the instance.
		// This approach does *NOT* require the Cloud SQL proxy.
		d, err := cloudsqlconn.NewDialer(ctx)
		if err != nil {
			return nil, err
		}
		return d.Dial(ctx, instanceConnectionName)
	}
	dbURI := stdlib.RegisterConnConfig(config)
	dbPool, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}
	return dbPool, nil
}

func InitPostgreSQL() *gorm.DB {
	var err error

	log.Printf("DB_INSTANCE: %s", os.Getenv("DB_INSTANCE"))
	if os.Getenv("DB_INSTANCE") == "" {
		return nil
	}

	sqlDB, err := connectWithConnector()
	if sqlDB == nil || err != nil {
		log.Fatalf("Error connecting db: %s", err)
		return nil
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Error pinging db: %s", err)
		return nil
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error initializing db: %s", err)
		return nil
	}

	return db
}
