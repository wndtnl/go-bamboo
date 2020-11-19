package bamboo

import "testing"

func TestNewSimpleClient(t *testing.T) {
	bambooClient := NewSimpleClient(nil, "admin", "admin")
	if bambooClient == nil {
		t.Fatalf("Could not initialize client")
	}
}
