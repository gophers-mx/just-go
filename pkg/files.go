package pkg

import (
	"fmt"
	"html/template"
	"os"

	"github.com/gophers-mx/just-go/config"
)

// CleanUp deletes the generated files in case of an error.
func CleanUp(err error) {
	fmt.Printf("There was an error running the application: %v\n", err)
	if err := os.RemoveAll(config.GetConfig().Name); err != nil {
		fmt.Printf("There was an error removing the project directory: %v\n", err)
	}
	os.Exit(1)
}

// CreateDirectory creates the project directory on the current path.
func CreateDirectory() {
	cfg := config.GetConfig()
	fmt.Printf("Creating base directory %q\n", cfg.Name)
	if _, err := os.Stat(cfg.Name); os.IsNotExist(err) {
		if errMkdir := os.MkdirAll(cfg.Name, os.ModePerm); errMkdir != nil {
			CleanUp(errMkdir)
		}
	}
}

// CopyFile copies a file to its new location without modify it.
func CopyFile(origin, destination, filename string) {
	cfg := config.GetConfig()
	newDirectoryName := fmt.Sprintf("%s/%s", cfg.Name, destination)

	if _, err := os.Stat(newDirectoryName); os.IsNotExist(err) {
		errMkdir := os.MkdirAll(newDirectoryName, os.ModePerm)
		if errMkdir != nil {
			CleanUp(errMkdir)
		}
	}

	fileContent, err := cfg.FS.ReadFile(fmt.Sprintf("%s/%s", origin, filename))
	if err != nil {
		CleanUp(err)
	}

	newFilename := fmt.Sprintf("%s/%s/%s", cfg.Name, destination, filename)
	if destination == "" {
		newFilename = fmt.Sprintf("%s/%s", cfg.Name, filename)
	}

	f, err := os.Create(newFilename)
	if err != nil {
		CleanUp(err)
	}

	defer func() {
		err := f.Close()
		if err != nil {
			CleanUp(err)
		}
	}()

	if _, err := f.Write(fileContent); err != nil {
		CleanUp(err)
	}
}

// GenFromTemplate copies a file to its new location, fills it's content and save it on the final destination.
// E.g.
// internal.GenFromTemplate("assets/shared", "/a/b/c", "README.md").
func GenFromTemplate(templatePath, destination, filename string) {
	cfg := config.GetConfig()
	newDirectoryName := fmt.Sprintf("%s%s", cfg.Name, destination)
	fmt.Printf("Generating %q from template\n", filename)

	tmpl, err := template.ParseFS(cfg.FS, fmt.Sprintf("%s/%s.tmpl", templatePath, filename))
	if err != nil {
		CleanUp(err)
	}

	if _, err := os.Stat(newDirectoryName); os.IsNotExist(err) {
		errMkdir := os.MkdirAll(newDirectoryName, os.ModePerm)
		if errMkdir != nil {
			CleanUp(errMkdir)
		}
	}

	file, err := os.Create(fmt.Sprintf("%s/%s", newDirectoryName, filename))
	if err != nil {
		CleanUp(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			CleanUp(err)
		}
	}()

	if err := tmpl.Execute(file, cfg); err != nil {
		CleanUp(err)
	}
}

// CopyDirectory copies an entire directory to a new location.
// E.g.
// internal.CopyDirectory)("assets/foo", "foo")
func CopyDirectory(origin, dirName string) {
	cfg := config.GetConfig()
	fmt.Printf("Copying %q\n", dirName)
	files, err := cfg.FS.ReadDir(origin)
	if err != nil {
		CleanUp(err)
	}

	newDirectoryName := fmt.Sprintf("%s/%s", cfg.Name, dirName)

	if _, err := os.Stat(newDirectoryName); os.IsNotExist(err) {
		if errMkdir := os.MkdirAll(newDirectoryName, os.ModePerm); errMkdir != nil {
			CleanUp(errMkdir)
		}
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filename := fmt.Sprintf("%s/%s", origin, file.Name())

		fileContent, err := cfg.FS.ReadFile(filename)
		if err != nil {
			CleanUp(err)
		}

		newFilename := fmt.Sprintf("%s/%s/%s", cfg.Name, dirName, file.Name())
		f, err := os.Create(newFilename)
		if err != nil {
			CleanUp(err)
		}

		if _, err := f.Write(fileContent); err != nil {
			CleanUp(err)
		}

		err = f.Close()
		if err != nil {
			CleanUp(err)
		}
	}
}
