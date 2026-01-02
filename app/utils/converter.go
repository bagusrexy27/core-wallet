package utils

import (
	"fmt"
	"time"
)

func StringToDateConvert(string) time.Time {
	dateStr := "2025-10-16"
	layout := "2006-01-02" // reference layout (must use this exact format)

	t, err := time.Parse(layout, dateStr)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return t
	}
	return t
}

func StringToInt(str string) (int, error) {
	var result int
	_, err := fmt.Sscanf(str, "%d", &result)
	if err != nil {
		return 0, err
	}
	return result, nil
}
