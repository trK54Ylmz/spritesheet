package util

import (
	"fmt"

	"time"
)

type LogOutput struct{}

func (o *LogOutput) Write(b []byte) (int, error) {
	d := time.Now().Format("2006-01-02 15:04:05.000")

	return fmt.Printf("[%s] %s", d, string(b))
}
