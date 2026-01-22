package configs

import (
	"log"
	"time"
)

var Location *time.Location

func InitTimezone() {
	var err error
	Location, err = time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Fatalf("❌ Gagal load timezone: %v", err)
	}
}
