package main

import (
	"fmt"
	pluralize "github.com/gertd/go-pluralize"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"path"
	"runtime"
	"strings"
)

var cliCmd = &cobra.Command{
	Use:  "crud",
	Args: cobra.MinimumNArgs(1),
	RunE: crud,
}

func crud(cmd *cobra.Command, args []string) error {
	name := args[0]

	hasDir, _ := afero.DirExists(afero.NewOsFs(), "pkg")
	if !hasDir {
		afero.NewOsFs().MkdirAll("pkg", 0755)
	}

	fs := afero.NewBasePathFs(afero.NewOsFs(), "pkg/")

	createFolders(fs, name)
	createFiles(fs, name)
	
	return nil
}

func createFolders(fs afero.Fs, name string) {
	fs.MkdirAll(name, 0755)
	fs.MkdirAll(name+"/controllers", 0755)
	fs.MkdirAll(name+"/models", 0755)
	fs.MkdirAll(name+"/routes", 0755)
	fs.MkdirAll(name+"/services", 0755)
}

func createFiles(fs afero.Fs, name string) {
	createFile(fs, name, "stubs/controller.stub", name+"/controllers/"+name+"_controller.go")
	createFile(fs, name, "stubs/model.stub", name+"/models/"+name+".go")
	createFile(fs, name, "stubs/route.stub", name+"/routes/api.go")
	createFile(fs, name, "stubs/service.stub", name+"/services/"+name+"_service.go")
}

func createFile(fs afero.Fs, name string, stubPath, filePath string) {
	fs.Create(filePath)

	_, filename, _, _ := runtime.Caller(1)
	stubPath = path.Join(path.Dir(filename), stubPath)

	contents, _ := fileContents(stubPath)
	contents = replaceStub(contents, name)

	err := overwrite("pkg/"+filePath, contents)
	if err != nil {
		fmt.Println(err)
	}
}

func fileContents(file string) (string, error) {
	a := afero.NewOsFs()
	contents, err := afero.ReadFile(a, file)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}

func overwrite(file string, message string) error {
	a := afero.NewOsFs()
	return afero.WriteFile(a, file, []byte(message), 0666)
}

func replaceStub(content string, name string) string {
	content = strings.Replace(content, "{{TitleName}}", Title(name), -1)
	content = strings.Replace(content, "{{PluralLowerName}}", Lower(Plural(name)), -1)
	content = strings.Replace(content, "{{SingularLowerName}}", Lower(Singular(name)), -1)
	return content
}

func Plural(name string) string {
	pluralize := pluralize.NewClient()

	return pluralize.Plural(name)
}

func Singular(name string) string {
	pluralize := pluralize.NewClient()
	return pluralize.Singular(name)
}

func Lower(name string) string {
	return strings.ToLower(name)
}

func Title(name string) string {
	return strings.Title(Lower(name))
}
