package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jhernandezj/moodle-settings-manager/internal/config"
	"github.com/jhernandezj/moodle-settings-manager/pkg/settings"
)

func main() {
	if err := run(); err != nil {
		log.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	fmt.Println("Moodle Settings Manager")

	cfg := config.New()
	manager := settings.New(cfg)

	// Example of getting a setting
	if _, err := manager.GetSetting("test_setting"); err != nil {
		return fmt.Errorf("failed to get setting: %w", err)
	}

	return nil
}
