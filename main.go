package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
	_ "github.com/olebedev/when/rules/ru"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <hh:mm> <text message>\n", os.Args[0])
		os.Exit(1)
	}

	reminderTime, err := parseReminderTime(os.Args[1])
	if err != nil {
		handleError(err, 2)
	}

	if reminderTime.Before(time.Now()) {
		handleError(fmt.Errorf("reminder time is in the past"), 4)
	}

	message := strings.Join(os.Args[2:], " ")

	timeDifference := reminderTime.Sub(time.Now())
	showCountdown("Reminder will be shown in", timeDifference)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			timeDifference = reminderTime.Sub(time.Now())
			if timeDifference <= 0 {
				showReminder("Reminder", message)
				return
			}
			showCountdown("Reminder will be shown in", timeDifference)
		}
	}
}

func parseReminderTime(timeStr string) (time.Time, error) {
	currentTime := time.Now()
	parser := when.New(nil)
	parser.Add(en.All...)
	parser.Add(common.All...)

	t, err := parser.Parse(timeStr, currentTime)
	if err != nil {
		return time.Time{}, err
	}

	if t == nil {
		return time.Time{}, fmt.Errorf("no reminder time found")
	}

	return t.Time, nil
}

func handleError(err error, exitCode int) {
	fmt.Println(err)
	os.Exit(exitCode)
}

func showReminder(title, message string) {
	err := beeep.Alert(title, message, "assets/information.png")
	if err != nil {
		handleError(err, 5)
	}
}

func showCountdown(message string, duration time.Duration) {
	fmt.Printf("%s %s\n", message, formatDuration(duration))
}

func formatDuration(duration time.Duration) string {
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
