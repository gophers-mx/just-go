package justgo

import (
	"embed"
	"fmt"
	"os"

	"github.com/gophers-mx/just-go/config"
	"github.com/gophers-mx/just-go/pkg"
)

type Generathor struct {
	Assets          embed.FS
	ProjectName     *string
	Generator       pkg.Generator
	TemplateDetails interface{}
}

func (g *Generathor) Run() {
	g.checkFlags()
	cfg := g.cfg()

	pkg.CreateDirectory()
	g.Generator.Run(cfg)
}

func (g *Generathor) checkFlags() {
	if *g.ProjectName == "" {
		fmt.Println("name cannot be blank")
		os.Exit(1)
	}
}

func (g *Generathor) cfg() *config.ProjectConfig {
	cfg := config.ProjectConfig{
		FS:              g.Assets,
		Name:            *g.ProjectName,
		TemplateDetails: g.TemplateDetails,
	}

	return config.New(cfg)
}
