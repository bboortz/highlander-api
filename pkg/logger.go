package pkg

import (
	"log"
)

// LogStartup is logging out at startup
func LogStartup(c Conf) {
	// startup logs
	log.Printf("Starting up highlander API")
	c.PrintConf()
	log.Printf("ready for incoming requests ...")
}
