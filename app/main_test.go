package main

import (
	"testing" // テストで使える関数・構造体が用意されているパッケージをimport
)

func TestExampleSuccess(t *testing.T) {
	_ = InitTimeZone("Asia/Tokyo")
}

func TestExampleFailed(t *testing.T) {
	_ = InitTimeZone("Asia/Tokyo")
}