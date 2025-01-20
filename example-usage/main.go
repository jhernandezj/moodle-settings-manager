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

	// Try to get all predefined settings
	settingNames := []string{"site_name", "maintenance_mode", "theme", "lang"}

	fmt.Println("Current settings:")
	for _, name := range settingNames {
		value, err := manager.GetSetting(name)
		if err != nil {
			log.Printf("Error getting %s: %v\n", name, err)
			continue
		}
		fmt.Printf("%s: %s\n", name, value)
	}

	// Update a setting
	err := manager.UpdateSetting("maintenance_mode", "enabled")
	if err != nil {
		log.Fatalf("Error updating setting: %v", err)
	}
	fmt.Println("\nAfter update:")
	value, _ := manager.GetSetting("maintenance_mode")
	fmt.Printf("maintenance_mode: %s\n", value)
}
