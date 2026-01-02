package utils

import (
	"fmt"
	"time"
)

func GenerateTimestampID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
