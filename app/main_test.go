package main

import (
	"testing" // テストで使える関数・構造体が用意されているパッケージをimport
)

func TestInitTimeZone(t *testing.T) {
	_ = InitTimeZone("Asia/Tokyo")
}

//func TestInitDB(t *testing.T) {
//	_ = InitDB(mustGetEnv("POSTGRES_HOST"), mustGetEnv("POSTGRES_PORT"), mustGetEnv("POSTGRES_USER"),
//		mustGetEnv("POSTGRES_PASSWORD"), mustGetEnv("POSTGRES_DB"), mustGetEnv("POSTGRES_SSL_MODE"))
//}
