package utils

import (
	"fmt"
	"os"
	"time"
)

func LogToFile(msg string) {
	f, err := os.OpenFile("log_husk.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()
	timestamp := time.Now().Format("02-01-2006 15:04:05")
	fmt.Fprintf(f, "[%s] %s\n", timestamp, msg)
}