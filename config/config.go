package config

import (
	"embed"
	"sync"
)

type ProjectConfig struct {
	FS              embed.FS
	Name            string
	TemplateDetails interface{}
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

// New sets the project config values using a singleton.
func New(cfg ProjectConfig) *ProjectConfig {
	once.Do(func() {
		config = &ProjectConfig{
			FS:              cfg.FS,
			Name:            cfg.Name,
			TemplateDetails: cfg.TemplateDetails,
		}
	})
	return config
}

// Config retrieves the project's config.
func Config() *ProjectConfig {
	return config
}
