# just-go

## Descripción

Este proyecto expone la funcionalidad para crear herramientas CLI para generar proyectos personalizados. No necesariamente de Golang, puede implementarse para cualquier tipo de proyecto.

## Uso

Instala `just-go` como biblioteca para tu generador.

```shell
$ go get github.com/gophers-mx/just-go
```

Crea una estructura con los campos necesarios para tus archivos template.

```golang
type Details struct {
    Name string
}
```

Crea un directorio de assets con tus templates de proyecto para utilizarlo en tu proyecto.

```golang
//go:embed assets/*
var assets embed.FS
```

Crea una estructura que implemente la interfaz `Generator` con el m'etodo `Run(cfg *config.Cfg)` y utiliza las funciones del paquete `files` para poder realizar las copias de archivos, directorios y generaci'on de templates.

```golang
package cli

import (
	"github.com/gophers-mx/just-go/config"
	"github.com/gophers-mx/just-go/pkg/files"
)

type CustomGenerator struct {}

func (lg *CustomGenerator) Run(cfg *config.Cfg) {
	files.GenFromTemplate("assets", "", "README.md")
	files.CopyDirectory("assets/cmd", "cmd")
	files.CopyFile("assets", "", "hello.md")
}
```

Crea una instancia de la estructura `Generathor` y ejecuta su m'etodo `Run()`.

```golang
func main() {
	generathor := &justgo.Generathor{
		Assets:      assets,
		ProjectName: "hello world",
		Generator:   cli.NewGeneratorCLI(),
		TemplateDetails: Details{
			Name: "hello world",
		},
	}

	generathor.Run()
}
```

Puedes encontrar un ejemplo completo [aqui](https://github.com/4k1k0/gen-cli).


## Contribuir

Si quieres contribuir al proyecto, puedes enviar tu pull request. Las sugerencias son bienvenidas.

## Licencia

[![GPLv3](https://www.gnu.org/graphics/gplv3-127x51.png "Licencia GNU General Public License")](https://www.gnu.org/licenses/gpl-3.0.html)

## Contacto

[Gohpers MX en Telegram](https://t.me/golangmx)
