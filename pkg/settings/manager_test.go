package settings

import "testing"

func TestNew(t *testing.T) {
	config := struct{}{}
	manager := New(config)
	if manager == nil {
		t.Error("Expected non-nil manager")
	}
}

func TestGetSetting(t *testing.T) {
	manager := New(struct{}{})
	_, err := manager.GetSetting("test")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
