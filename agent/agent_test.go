package agent

import (
	"testing"
)

func TestGetEndPoints(t *testing.T) {
	cfg := &Config{
		Port:      "5000",
		Websocket: "ws://www.example.com",
		RestURI:   "getsmth",
	}
	res := cfg.GetEndPoints()
	if res != "ws://www.example.com:5000/getsmth" {
		t.Errorf("Expected: ws://www.example.com:5000/getsmth, got: %s", res)
	}
}
