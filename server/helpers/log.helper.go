package helpers

import (
	"log"
	"net/http"
	"strings"
)

func LogHTTPHeader(headers http.Header) {
	for key, value := range headers {
		if len(value) == 1 {
			log.Println(key, value[0])
			continue
		}
		log.Println(key, strings.Join(value, ","))

	}
}
