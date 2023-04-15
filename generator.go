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

func Run(assets embed.FS, projectName, version *string, generator pkg.Generator) {
	checkFlags(projectName, version)
	cfg := New(projectName, version, assets)

	pkg.CreateDirectory()
	generator.Run(cfg)
}

func checkFlags(projectName, version *string) {
	if *projectName == "" {
		fmt.Println("name cannot be blank")
		os.Exit(1)
	}

	if *version == "" {
		*version = runtime.Version()
	}

	*version = cleanVersion(*version)
}

func cleanVersion(v string) (res string) {
	defer func() {
		if recover() != nil {
			res = "1.18"
		}
	}()

	vn := strings.Split(v, "go")[1]
	vs := strings.Split(vn, ".")

	res = fmt.Sprintf("%s.%s\n", vs[0], vs[1])

	return
}

func New(projectName, version *string, assets embed.FS) *config.ProjectConfig {
	cfg := config.ProjectConfig{
		FS:      assets,
		Name:    *projectName,
		Version: *version,
	}

	return config.New(cfg)
}
