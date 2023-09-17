package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
	_ "github.com/olebedev/when/rules/ru"
)

const (
	markName  = "GOLANG_CLI_REMINDER"
	markValue = "1"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Printf("Usage: %s <hh:mm> <text message>\n", os.Args[0])
		os.Exit(1)
	}

	now := time.Now()

	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	t, err := w.Parse(os.Args[1], now)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	if t == nil {
		fmt.Println("No reminder time found")
		os.Exit(3)
	}

	if t.Time.Before(now) {
		fmt.Println("Reminder time is in the past")
		os.Exit(4)
	}

	difference := t.Time.Sub(now)
	if os.Getenv(markName) == markValue {
		time.Sleep(difference)
		err := beeep.Alert("Reminder", strings.Join(os.Args[2:], ""), "assets/information.png")
		if err != nil {
			fmt.Println(err)
			os.Exit(5)
		}
	} else {
		cmd := exec.Command(os.Args[0], os.Args[1:]...)
		cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", markName, markValue))
		err := cmd.Start()
		if err != nil {
			fmt.Println(err)
			os.Exit(6)
		}

		fmt.Println("Reminder will be shown in", difference.Round(time.Second))
	}
}
