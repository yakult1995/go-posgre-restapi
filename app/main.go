package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

var db *sql.DB

func main() {
	// Timezoneを東京に
	InitTimeZone(mustGetenv("TIME_ZONE"))

	// HTTPルーターを初期化
	router := httprouter.New()

	InitDB(mustGetenv("POSTGRES_HOST"), mustGetenv("POSTGRES_PORT"), mustGetenv("POSTGRES_USER"),
		mustGetenv("POSTGRES_PASSWORD"), mustGetenv("POSTGRES_DB"), mustGetenv("POSTGRES_SSL_MODE"))

	// routing
	router.GET("/", Root)
	router.GET("/users", ListUsers)
	router.GET("/users/:id", DetailUser)
	router.POST("/users", CreateUser)
	router.PUT("/users/:id", UpdateUser)
	router.DELETE("/users/:id", DeleteUser)

	// Webサーバーを8080ポートで立ち上げる
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

func InitDB(postgresHost, postgresPort, postgresUser, postgresPassword, postgresDB, postgresSSLMode string) {
	dataSourceName := "host=" + postgresHost + " port=" + postgresPort + " user=" + postgresUser +
		" password=" + postgresPassword + " dbname=" + postgresDB + " sslmode=" + postgresSSLMode
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}

func InitTimeZone(timeZone string) {
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		panic(err)
	}
	time.Local = loc
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("%s environment variable not set.", k)
	}
	return v
}