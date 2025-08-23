package internal

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
	"sigs.k8s.io/yaml"
	"strings"
)

type Config struct {
	DbPath     string               `json:"dbpath"`
	BookMarks  map[string]*BookMark `json:"bookmarks"`
	LastOpened string               `json:"lastopened"`
}

func (c Config) writeConfig() error {
	path := getConfigPath()
	if path == "" {
		_, err := writeNewConfig()
		if err != nil {
			return err
		}
		// no need to do anything else if we just made the config, it should be empty
		return nil
	}
	yamlBytes, err := yaml.Marshal(c)
	err = os.WriteFile(path, yamlBytes, 0600)
	if err != nil {
		return err
	}
	return nil
}

// Get config from system
func getConfig() (*Config, error) {
	var cfg Config
	path := getConfigPath()
	if path == "" {
		_, err := writeNewConfig()
		if err != nil {
			return nil, err
		}
		// no need to do anything else if we just made the config, it should be empty
		return &cfg, nil
	}
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

func getConfigPath() string {
	homedir := os.Getenv("HOME")
	suspects := []string{
		homedir + "/.config/bookworm/config.yml",
		homedir + "/.config/bookworm/config.yaml",
		"/etc/bookworm/config.yml",
		"/etc/bookworm/config.yaml",
	}
	for _, s := range suspects {
		_, try := os.Stat(s)
		if os.IsExist(try) {
			return s
		}
	}
	return ""
}

// Writes a new config & returns an *os.File
// This will write to ~/.config/bookworm/config.yml
func writeNewConfig() (*os.File, error) {
	homedir := os.Getenv("HOME")
	if getConfigPath() != "" {
		return nil, errors.New("config is already there!")
	}
	configInfo, err := os.Stat(homedir + "/.config/bookworm")
	if err != nil {
		return nil, err
	}
	if !configInfo.IsDir() {
		return nil, errors.New("~/.config/bookworm is not a directory!")
	}
	f, err := os.Create(homedir + "/.config/bookworm/config.yml")
	return f, nil
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
