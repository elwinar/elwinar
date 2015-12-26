// +build !docker

package main

import (
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("env")
}
