package justgo

import (
	"embed"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/gophers-mx/just-go/config"
	"github.com/gophers-mx/just-go/pkg"
)

type Generathor struct {
	Assets      embed.FS
	ProjectName *string
	Version     *string
	Generator   pkg.Generator
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

	if *g.Version == "" {
		*g.Version = runtime.Version()
	}

	*g.Version = g.cleanVersion()
}

func (g *Generathor) cleanVersion() (res string) {
	defer func() {
		if recover() != nil {
			res = "1.18"
		}
	}()

	vn := strings.Split(*g.Version, "go")[1]
	vs := strings.Split(vn, ".")

	res = fmt.Sprintf("%s.%s\n", vs[0], vs[1])

	return
}

func (g *Generathor) cfg() *config.ProjectConfig {
	cfg := config.ProjectConfig{
		FS:      g.Assets,
		Name:    *g.ProjectName,
		Version: *g.Version,
	}

	return config.New(cfg)
}
