package internal

import (
	"os"
	"strings"
	"testing"
)

func TestPaths(t *testing.T) {
	cfgPath := strings.Split(getConfigPath("/tmp/"), "/")
	dbPath := strings.Split(getDbPath("/tmp/"), "/")
	if dbPath[len(dbPath)-1] != "worm.db" {
		t.Fatalf("DB path is wrong")
	}
	if cfgPath[len(cfgPath)-1] != "config.yml" {
		t.Fatalf("Config path is wrong")
	}
}

func TestInitConfig(t *testing.T) {
	pathTo := os.TempDir()
	defer os.Remove(pathTo)
	cfg, err := initConfig(pathTo)
	if err != nil {
		t.Fatalf("initConfig returned an error %s", err)
	}
	if cfg == nil {
		t.Fatalf("initConfig returned a nil object")
	}
}

func TestIsValidUrl(t *testing.T) {}
