package main

import (
	"context"
	"fmt"
	"time"

	"bytes"
	"os/exec"
	"strconv"
	"strings"
)

var monitoring = false

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) StartMonitoring() {
	if !monitoring {
		monitoring = true
		go monitorBattery()
	}
}

func (a *App) StopMonitoring() {
	monitoring = false	
}

func monitorBattery() {
	for monitoring {
		batteryLevel, isCharging := getBatteryStatus()
		if batteryLevel == -1 {
			fmt.Println("Failed to get battery status.")
			return
		}
		fmt.Printf("Current battery level: %d%%, Charging: %t\n", batteryLevel, isCharging)

		if batteryLevel >= 80 && isCharging {
			showNotification(batteryLevel)
			time.Sleep(1 * time.Hour) // Adjust based on preference
		} else {
			time.Sleep(2 * time.Second) // Check every 5 seconds
		}
	}
}

func getBatteryStatus() (int, bool) {
	cmd := exec.Command("pmset", "-g", "batt")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return 0, false
	}

	output := out.String()
	lines := strings.Split(output, "\n")
	batteryInfo := lines[1]

	start := strings.Index(batteryInfo, "\t") + 1
	end := strings.Index(batteryInfo, "%")
	batteryLevel, err := strconv.Atoi(batteryInfo[start:end])
	if err != nil {
		fmt.Println("Error parsing battery level:", err)
		return 0, false
	}

	isCharging := strings.Contains(output, "AC Power")

	return batteryLevel, isCharging
}

func showNotification(batteryLevel int) {
	notification := fmt.Sprintf("display notification \"Battery level is at %d%%\" with title \"Battery Notification\"", batteryLevel)
	soundFilePath := "UIBeep.aac"
	exec.Command("afplay", soundFilePath).Run()
	cmd := exec.Command("osascript", "-e", notification)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error showing notification:", err)
	}
}
