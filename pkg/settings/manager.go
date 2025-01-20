package settings

import "fmt"

// Manager handles Moodle settings operations
type Manager struct {
	config   interface{}
	settings map[string]string
}

// New creates a new settings manager instance
func New(config interface{}) *Manager {
	return &Manager{
		config: config,
		settings: map[string]string{
			"site_name":        "My Moodle Site",
			"maintenance_mode": "disabled",
			"theme":            "boost",
			"lang":             "en",
		},
	}
}

// GetSetting retrieves a setting value by name
func (m *Manager) GetSetting(name string) (string, error) {
	if value, exists := m.settings[name]; exists {
		return value, nil
	}
	return "", fmt.Errorf("setting %s not found", name)
}

// UpdateSetting updates a setting value
func (m *Manager) UpdateSetting(name, value string) error {
	m.settings[name] = value
	return nil
}
