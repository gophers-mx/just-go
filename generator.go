package justgo

import (
	"embed"
	"fmt"
	"os"
	"strings"

	"github.com/gophers-mx/just-go/config"
	"github.com/gophers-mx/just-go/pkg"
)

func Run(assets embed.FS, projectName, projectType *string, generator pkg.Generator) {
	checkFlags(projectName, projectType)
	cfg := setConfig(projectName, projectType, assets)

	pkg.CreateDirectory()
	generator.Run(cfg)
}

func checkFlags(projectName, projectType *string) {
	if *projectName == "" {
		fmt.Println("name cannot be blank")
		os.Exit(1)
	}

	if *projectType == "" {
		fmt.Println("type cannot be blank")
		os.Exit(1)
	}

	checkedType, ok := config.ValidTypes[strings.ToLower(*projectType)]
	if !ok {
		fmt.Println("type is not valid")
		os.Exit(1)
	}
	*projectType = checkedType
}

func setConfig(projectName, projectType *string, assets embed.FS) *config.ProjectConfig {
	cfg := config.ProjectConfig{
		FS:   assets,
		Name: *projectName,
		Type: *projectType,
	}

	return config.SetConfig(cfg)
}
