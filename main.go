package main

import (
	"log"

	"github.com/IllyaYalovyy/learn-ds-log1/internal/server"
)

func main() {
	srv := server.New(":8080")
	log.Fatal(srv.ListenAndServe())
}
