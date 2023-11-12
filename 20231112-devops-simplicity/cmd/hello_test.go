package main

import (
	"testing"
	"time"
)

func TestGetCurrentTime(t *testing.T) {
	tests := []struct {
		name string
	}{
		// Test case 1: Testing the current time format
		{name: "Test current time format"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getCurrentTime()
			// Check if the current time format matches the expected format
			if _, err := time.Parse("02.01.2006 - 15:04:05", got); err != nil {
				t.Errorf("getCurrentTime() = %v, expected time in format 02.01.2006 - 15:04:05", got)
			}
		})
	}
}
