package main

import (
	"log"

	"github.com/margarote/approve_request_golang/approverequest"
)

func main() {
	rsult, _ := approverequest.SendValidationPost(".s", "google.com", 1736972931)

	log.Println(rsult)
}
