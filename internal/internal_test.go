package internal

import (
	"os"
	"strings"
	"testing"
)

func TestPaths(t *testing.T) {
	pathTo := os.TempDir()
	t.Log("Using Config Dir " + pathTo)
	defer os.Remove(pathTo)
	cfgPath := strings.Split(getConfigPath(pathTo), "/")
	dbPath := strings.Split(getDbPath(pathTo), "/")
	t.Log(cfgPath, dbPath)
	if dbPath[len(dbPath)-1] != "worm.db" {
		t.Fatalf("DB path is wrong")
	}
	if cfgPath[len(cfgPath)-1] != "config.yml" {
		t.Fatalf("Config path is wrong")
	}
	if getConfigDir(pathTo) != pathTo {
		t.Fatalf("Config dir is wrong")
	}
}

func TestInitConfig(t *testing.T) {
	// pathTo := os.TempDir()
	// t.Log("Using Config Dir " + pathTo)
	// defer os.Remove(pathTo)
	// cfg, err := initConfig(pathTo)
	// if err != nil {
	// 	t.Fatalf("initConfig returned an error %s", err)
	// }
	// if cfg == nil {
	// 	t.Fatalf("initConfig returned a nil object")
	// }
}

func TestIsValidUrl(t *testing.T) {}
