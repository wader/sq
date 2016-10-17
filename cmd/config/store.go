package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/neilotoole/sq/libsq/drvr"
	"gopkg.in/yaml.v2"
)

// Store saves and loads config.
type Store interface {
	// Save writes config to the store.
	Save(cfg *Config) error
	// Load reads config from the store.
	Load() (*Config, error)
}

// FileStore provides file-based persistence of config.
type FileStore struct {
	Path string
}

func NewFileStore(path string) (*FileStore, error) {

	fs := &FileStore{path}
	err := fs.checkFile()
	if err != nil {
		return nil, err
	}

	return fs, nil
}

func (f *FileStore) String() string {
	return fmt.Sprintf("config filestore: %v", f.Path)
}

// Load reads config from disk.
func (f *FileStore) Load() (*Config, error) {

	bytes, err := ioutil.ReadFile(f.Path)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = yaml.Unmarshal(bytes, cfg)
	if err != nil {
		return nil, err
	}

	applyDefaults(cfg)

	if cfg.SourceSet == nil {
		cfg.SourceSet = drvr.NewSourceSet()
	}

	return cfg, nil
}

// Save writes config to disk.
func (f *FileStore) Save(cfg *Config) error {

	bytes, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(f.Path, bytes, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// checkFile verifies that the file at f.Path exists, and if not, creates it etc.
func (f *FileStore) checkFile() error {
	_, err := os.Stat(f.Path)
	if err == nil {
		// File exists
		return nil
	}

	if !os.IsNotExist(err) {
		// some other kind of error, return it
		return fmt.Errorf("config: error with backing file '%v': %v", f.Path, err)
	}

	// File doesn't exist, create it.

	// Create the path
	parent := filepath.Dir(f.Path)
	err = os.MkdirAll(parent, os.ModePerm)
	if err != nil {
		return fmt.Errorf("config: backing file not created '%v': %v", f.Path, err)
	}

	conf := NewConfig()
	return f.Save(conf)
}

// InMemoryStore is a memory-based impl of Store. Useful for testing.
type InMemoryStore struct {
}

func (f *InMemoryStore) String() string {
	return "in-memory config store"
}

// Load returns a new config
func (f *InMemoryStore) Load() (*Config, error) {

	return NewConfig(), nil
}

// Save is a no-op
func (f *InMemoryStore) Save(cfg *Config) error {

	return nil
}
