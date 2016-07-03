package myutils

import (
	"log"
)

// LogFatalError it's a snippet for
//     if err != nil {
//         log.Fatalf("Error: %s", err)
//     }
func LogFatalError(err error) {
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}
