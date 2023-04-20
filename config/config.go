package config

import (
	"embed"
	"sync"
)

var config *Cfg
var once sync.Once

type Cfg struct {
	FS              embed.FS
	Name            string
	TemplateDetails interface{}
}

// New sets the project config values using a singleton.
func New(cfg Cfg) *Cfg {
	once.Do(func() {
		config = &Cfg{
			FS:              cfg.FS,
			Name:            cfg.Name,
			TemplateDetails: cfg.TemplateDetails,
		}
	})
	return config
}

// Config retrieves the project's config.
func Config() *Cfg {
	return config
}
