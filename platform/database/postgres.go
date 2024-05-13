package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	C "github.com/amirrstm/go-clinique/pkg/config"
)

var PostgresConn *sql.DB

func GetPostgresURL() string {
	dbHost := C.DBCfg().Host
	dbPort := C.DBCfg().Port
	dbUser := C.DBCfg().User
	dbName := C.DBCfg().Name
	dbPass := C.DBCfg().Password

	if C.DBCfg().SslMode == "disable" {
		return fmt.Sprintf("host=%s port=%v user=%s "+
			"password=%s dbname=%s sslmode=disable",
			dbHost, dbPort, dbUser, dbPass, dbName)
	} else {
		return fmt.Sprintf("host=%s port=%v user=%s "+
			"password=%s dbname=%s sslmode=%s",
			dbHost, dbPort, dbUser, dbPass, dbName, C.DBCfg().SslMode)
	}
}

func Init() error {
	var err error
	PostgresConn, err = sql.Open("postgres", GetPostgresURL())
	if err != nil {
		return fmt.Errorf("error opening database connection: %w", err)
	}

	err = PostgresConn.Ping()
	if err != nil {
		return fmt.Errorf("error pinging database: %w", err)
	}

	PostgresConn.SetMaxOpenConns(C.DBCfg().MaxOpenConn)
	PostgresConn.SetMaxIdleConns(C.DBCfg().MaxIdleConn)

	return nil
}

func PGTransaction(ctx context.Context) (*sql.Tx, error) {
	tx, err := PostgresConn.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func Close() {
	PostgresConn.Close()
}
