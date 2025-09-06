package internal

import (
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"sigs.k8s.io/yaml"
)

var (
	configDir = os.Getenv("HOME") + "/.config/bookworm/"
)

type Config struct {
	DbPath     string               `json:"dbpath"`
	BookMarks  map[string]*BookMark `json:"bookmarks"`
	LastOpened string               `json:"lastopened"`
}

func (c *Config) writeConfig() error {
	path := getConfigPath()
	if path == "" {
		return errors.New("the config path is not there!")
	}
	yamlBytes, err := yaml.Marshal(&c)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(path, yamlBytes, 0777)
	if err != nil {
		return err
	}
	return nil
}

// Get config from system
// Returns an error if the Config cannot be found
func getConfig() (*Config, error) {
	var cfg Config
	path := getConfigPath()
	_, err := os.Stat(path)
	// Create the config files if they don't exist
	if err != nil {
		return nil, err
	}
	// Read in Config
	yamlBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlBytes, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

// Returns the absolute path to the config file
func getConfigPath() string {
	return configDir + "config.yml"
}

// Returns the absolute path to the db file
func getDbPath() string {
	return configDir + "worm.db"
}

// Writes a new config & returns an *os.File
// This will write to ~/.config/bookworm/config.yml
// .. or it will blow up
func initConfig() (*Config, error) {
	configInfo, err := os.Stat(configDir)
	// Create the config.yml if it's not there
	if os.IsNotExist(err) {
		errr := os.Mkdir(configDir, 0666)
		if errr != nil {
			return nil, errr
		}
	} else if err != nil {
		return nil, err
	}
	if !configInfo.IsDir() {
		return nil, errors.New("~/.config/bookworm is not a directory!")
	}
	_, err = os.Create(getConfigPath())
	if err != nil {
		return nil, err
	}
	cfg := &Config{
		DbPath:     getDbPath(),
		BookMarks:  make(map[string]*BookMark),
		LastOpened: "nothing... yet",
	}
	err = cfg.writeConfig()
	return cfg, err
}

// openURL opens the specified URL in the default browser of the user.
// From https://gist.github.com/sevkin/9798d67b2cb9d07cb05f89f14ba682f8
func OpenURL(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default:
		if isWSL() {
			cmd = "cmd.exe"
			args = []string{"/c", "start", url}
		} else {
			cmd = "xdg-open"
			args = []string{url}
		}
	}
	if len(args) > 1 {
		// args[0] is used for 'start' command argument, to prevent issues with URLs starting with a quote
		args = append(args[:1], append([]string{""}, args[1:]...)...)
	}
	return exec.Command(cmd, args...).Start()
}

// isWSL checks if the Go program is running inside Windows Subsystem for Linux
// From https://gist.github.com/sevkin/9798d67b2cb9d07cb05f89f14ba682f8
func isWSL() bool {
	releaseData, err := exec.Command("uname", "-r").Output()
	if err != nil {
		return false
	}
	return strings.Contains(strings.ToLower(string(releaseData)), "microsoft")
}

// This is going to need a lot of work
func IsValidUrl(url string) bool {
	return strings.Contains(url, "https://") ||
		strings.Contains(url, "http://")
}

func bytesToBookMark(buf []byte) (*BookMark, error) {
	bmToReturn := &BookMark{}
	err := json.Unmarshal(buf, bmToReturn)
	if err != nil {
		return nil, err
	}
	return bmToReturn, nil
}
