package helpers

import (
	"log"
)

func CheckErrorAndPanic(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
