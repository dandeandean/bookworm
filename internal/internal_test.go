package internal

import (
	"strings"
	"testing"
)

func TestPaths(t *testing.T) {
	cfgPath := strings.Split(getConfigPath(), "/")
	dbPath := strings.Split(getDbPath(), "/")
	if dbPath[len(dbPath)-1] != "worm.db" {
		t.Fatalf("DB path is wrong")
	}
	if cfgPath[len(cfgPath)-1] != "config.yml" {
		t.Fatalf("Config path is wrong")
	}
	t.Log("OK")
}

func TestInitConfig(t *testing.T) {
	t.Log("OK")
}

func TestIsValidUrl(t *testing.T) {
	t.Log("OK")
}
