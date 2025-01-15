package main

import (
	"log"

	"github.com/margarote/approve_request_golang/approverequest"
)

func main() {
	rsult, _ := approverequest.SendValidationPost("123", "google.com", 2432223)

	log.Println(rsult)
}
