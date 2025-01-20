package main

import (
	"fmt"
	"log"

	"github.com/jhernandezj/moodle-settings-manager/pkg/settings"
)

func main() {
	// Create a simple config struct
	config := struct {
		DBHost string
		DBName string
	}{
		DBHost: "localhost",
		DBName: "moodle",
	}

	// Create a new settings manager
	manager := settings.New(config)

	// Example: Get a setting
	value, err := manager.GetSetting("site_name")
	if err != nil {
		log.Fatalf("Error getting setting: %v", err)
	}
	fmt.Printf("Site name: %s\n", value)

	// Example: Update a setting
	err = manager.UpdateSetting("maintenance_mode", "enabled")
	if err != nil {
		log.Fatalf("Error updating setting: %v", err)
	}
	fmt.Println("Successfully updated maintenance mode setting")
}
