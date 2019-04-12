package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

var db *sql.DB

func main() {
	// Timezoneを東京に
	InitTimeZone(mustGetEnv("TIME_ZONE"))

	// HTTPルーターを初期化
	router := httprouter.New()

	InitDB(mustGetEnv("POSTGRES_HOST"), mustGetEnv("POSTGRES_PORT"), mustGetEnv("POSTGRES_USER"),
		mustGetEnv("POSTGRES_PASSWORD"), mustGetEnv("POSTGRES_DB"), mustGetEnv("POSTGRES_SSL_MODE"))

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

func InitDB(postgresHost, postgresPort, postgresUser, postgresPassword, postgresDB, postgresSSLMode string) bool {
	dataSourceName := "host=" + postgresHost + " port=" + postgresPort + " user=" + postgresUser +
		" password=" + postgresPassword + " dbname=" + postgresDB + " sslmode=" + postgresSSLMode
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		fmt.Println("sql open failed")
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		fmt.Println("response failed")
		log.Fatal(err)
	}
	return true
}

func InitTimeZone(timeZone string) bool{
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		panic(err)
	}
	time.Local = loc
	return true
}

func mustGetEnv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("%s environment variable not set.", k)
	}
	return v
}