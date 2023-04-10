package main

import (
	"log"

	"github.com/glebtorbin/stock-rest-api/Internal/app/apiserver"
)

func main() {
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
