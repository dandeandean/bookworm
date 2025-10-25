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
	configDir      = os.Getenv("HOME") + "/.config/bookworm/"
	configFileName = "config.yml"
	dbFileName     = "worm.db"
	dirPerms       = os.FileMode(0700)
	configPerms    = os.FileMode(0666)
	dbPerms        = os.FileMode(0600)
	verboseMode    = os.Getenv("BW_VERBOSE") == "true"
)

type Config struct {
	DbPath         string `json:"dbpath"`
	LastOpened     string `json:"lastopened"`
	FzfIntegration bool   `json:"fzf"`
}

func (c *Config) writeConfig(pathTo string) error {
	path := getConfigPath(pathTo)
	if path == "" {
		return errors.New("the config path is not there!")
	}
	yamlBytes, err := yaml.Marshal(&c)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(path, yamlBytes, configPerms)
	if err != nil {
		return err
	}
	return nil
}

// Get config from system
// Returns an error if the Config cannot be found
func getConfig(pathTo string) (*Config, error) {
	var cfg Config
	path := getConfigPath(pathTo)
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

// Adds a slash at the end if it's not there
// does nothing if it is there
func withSlash(pathTo string) string {
	if pathTo[len(pathTo)-1] != '/' &&
		pathTo != "" {
		return pathTo + "/"
	}
	return pathTo
}

// Returns the absolute path to the config file
func getConfigPath(pathTo string) string {
	if pathTo == "" {
		return withSlash(configDir) + configFileName
	}
	return withSlash(pathTo) + configFileName
}

// Returns the absolute path to the db file
func getDbPath(pathTo string) string {
	if pathTo == "" {
		return withSlash(configDir) + dbFileName
	}
	return withSlash(pathTo) + dbFileName
}

// Returns the config dir
// defaults to ~/.config/
func getConfigDir(pathTo string) string {
	if pathTo == "" {
		return withSlash(configDir)
	}
	return withSlash(pathTo)
}

// Writes a new config & returns an *os.File
// This will write to $pathTo+config.yml
func initConfig(pathTo string) (*Config, error) {
	pathTo = getConfigDir(pathTo)
	configDirInfo, err := os.Stat(pathTo)
	// Create the configDir if it's not there
	if os.IsNotExist(err) {
		err = os.Mkdir(pathTo, dirPerms)
		if err != nil {
			return nil, err
		}
		configDirInfo, err = os.Stat(pathTo)
	}
	if err != nil {
		return nil, err
	}
	if !configDirInfo.IsDir() {
		return nil, errors.New(pathTo + " is not a directory!")
	}
	_, err = os.Create(
		getConfigPath(pathTo),
	)
	if err != nil {
		return nil, err
	}
	var fzfThere = false
	fzf, _ := exec.LookPath("fzf")
	if fzf != "" {
		fzfThere = true
	}
	cfg := &Config{
		DbPath:         getDbPath(pathTo),
		LastOpened:     "nothing... yet",
		FzfIntegration: fzfThere,
	}
	err = cfg.writeConfig(pathTo)
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
