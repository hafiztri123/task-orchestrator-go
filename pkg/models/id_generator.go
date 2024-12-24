package models

import (
	"fmt"
	"time"
)

//NOTE: This generation method was for testing only and should not be used for production
func generateID() string {
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("task_%d", timestamp)
}