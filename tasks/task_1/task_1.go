package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

type Time struct {
	Hour    int
	Minute  int
	Seconds int
}

func (t Time) String() string {
	return fmt.Sprintf("%02d:%02d:%02d", t.Hour, t.Minute, t.Seconds)
}

func main() {
	now, err := timeNow()

	if err == nil {
		fmt.Println("Time now:", now)
	} else {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
		os.Exit(1)
	}
}

func timeNow() (Time, error) {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err == nil {
		return Time{
			Hour:    time.Hour(),
			Minute:  time.Minute(),
			Seconds: time.Second(),
		}, nil
	} else {
		return Time{}, err
	}
}
