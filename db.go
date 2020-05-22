package pagination

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

// InitDB initializes the database instance
func InitDB() *sql.DB {
	host := MustHaveEnv("POSTGRES_HOST")
	port := MustHaveEnvInt("POSTGRES_PORT")
	user := MustHaveEnv("POSTGRES_USER")
	password := MustHaveEnv("POSTGRES_PASSWORD")
	dbname := MustHaveEnv("POSTGRES_DATABASE")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		logrus.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		logrus.Fatal(err)
	}

	// Set Max Connection Life time
	maxConLifeTime := MustHaveEnvInt("DB_MAX_CONN_LIFE_TIME_S")
	if maxConLifeTime <= 0 {
		logrus.Fatal(err, "DB_MAX_CONN_LIFE_TIME_S is not well set ")
	}
	db.SetConnMaxLifetime(time.Second * time.Duration(maxConLifeTime))

	// Set Max Open Connection
	maxOpenConn := MustHaveEnvInt("DB_MAX_OPEN_CONNECTION")
	if maxOpenConn <= 0 {
		logrus.Fatal(err, "DB_MAX_OPEN_CONNECTION is not well set ")
	}
	db.SetMaxOpenConns(maxOpenConn)

	// Set max Idle Connection that can be reused
	maxIdleConn := MustHaveEnvInt("DB_MAX_IDLE_CONNECTION")
	if maxIdleConn <= 0 {
		logrus.Fatal(err, "DB_MAX_IDLE_CONNECTION is not well set ")
	}
	db.SetMaxIdleConns(maxIdleConn)

	return db
}
