package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func formatTime(localNow time.Time, globalNow time.Time) string {
	timeFormat := "15:04:05"
	return fmt.Sprintf("%s / %s", localNow.Format(timeFormat), globalNow.Format(timeFormat))
}

func main() {
	globalNow, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	localNow := time.Now()

	if err == nil {
		fmt.Println("Time now:", formatTime(localNow, globalNow))
	} else {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
		os.Exit(1)
	}
}
