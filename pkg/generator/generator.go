package generator

import (
	"github.com/gophers-mx/just-go/config"
)

// Generator expose a method that can be implemented by different project generators.
type Generator interface {
	Run(cfg *config.ProjectConfig)
}
