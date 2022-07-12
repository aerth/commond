package commond

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

// Write toml config to filename, creating subdirectories if necessary.
func WriteConfig(filename string, config any) error {
	var (
		buf = &bytes.Buffer{}
		enc = toml.NewEncoder(buf)
	)

	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		return err
	}
	if err := enc.Encode(config); err != nil {
		return err
	}
	if _, err := os.Stat(filename); err == nil {
		return fmt.Errorf("config file exists, not overwriting")
	}
	return ioutil.WriteFile(filename, buf.Bytes(), 0600)
}

// Read toml config from file.
// If file doesn't exist and defaultValue is non-nil, write default config and read the newly created file.
func ReadConfigFile(filename string, configptr any, writeNewIfNotExist bool) error {
	_, err := toml.DecodeFile(filename, configptr)
	if err == nil {
		return nil
	}
	if !writeNewIfNotExist || !strings.Contains(err.Error(), "no such file") {
		return fmt.Errorf("reading toml file %s: %v", filename, err)
	}
	if err := WriteConfig(filename, configptr); err != nil {
		return fmt.Errorf("writing default toml config: %v", err)
	}
	_, err = toml.DecodeFile(filename, configptr)
	if err != nil {
		return fmt.Errorf("reading new toml file %s: %v", filename, err)
	}
	return nil
}
