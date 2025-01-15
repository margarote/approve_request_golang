package main

import (
	"log"

	"github.com/margarote/approve_request_golang/approverequest"
)

func main() {
	secretKey := ""

	data := approverequest.GenerateCodeWithTimestamp(secretKey, 30)

	rsult, _ := approverequest.SendValidationPost(data.Code, "google.com", data.Timestamp)

	log.Println(rsult)
}
