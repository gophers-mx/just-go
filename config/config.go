package config

import (
	"embed"
	"sync"
)

type ProjectConfig struct {
	FS   embed.FS
	Name string
	Type string
}

const (
	ProjectTypeCLI = "CLI"
)

var config *ProjectConfig
var once sync.Once

// ValidTypes holds the project types that this tool can generate.
var ValidTypes = map[string]string{
	"cli": ProjectTypeCLI,
}

// SetConfig sets the project config values using a singleton.
func SetConfig(cfg ProjectConfig) *ProjectConfig {
	once.Do(func() {
		config = &ProjectConfig{
			FS:   cfg.FS,
			Name: cfg.Name,
			Type: cfg.Type,
		}
	})
	return config
}

// GetConfig retrieves the project's config.
func GetConfig() *ProjectConfig {
	return config
}
