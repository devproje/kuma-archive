package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type ConfigRef struct {
	Port int `json:"port"`
}

var (
	ROOT_DIRECTORY string
	CONFIG_FILE    string
	INDEX_DIR      string
)

func init() {
	ROOT_DIRECTORY = "ka_data"
	if os.Getenv("KA_PATH") != "" {
		ROOT_DIRECTORY = os.Getenv("KA_PATH")
	}

	if _, err := os.ReadDir(ROOT_DIRECTORY); err != nil {
		raw := ConfigRef{
			Port: 8080,
		}
		buf, _ := json.MarshalIndent(raw, "", "    ")

		_ = os.Mkdir(ROOT_DIRECTORY, 0755)
		_ = os.WriteFile(filepath.Join(ROOT_DIRECTORY, "config.json"), []byte(buf), 0644)
	}

	CONFIG_FILE = filepath.Join(ROOT_DIRECTORY, "config.json")

	INDEX_DIR = filepath.Join(ROOT_DIRECTORY, "index")
	if _, err := os.ReadDir(INDEX_DIR); err != nil {
		_ = os.Mkdir(INDEX_DIR, 0755)
	}
}

func Get() *ConfigRef {
	buf, err := os.ReadFile(CONFIG_FILE)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		return nil
	}

	var config ConfigRef
	if err = json.Unmarshal(buf, &config); err != nil {
		return nil
	}

	return &config
}
