package app

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"

	"git.jiaxianghudong.com/go/goproject/assets"

	"github.com/spf13/cobra"
)

var (
	basePath    string
	projectPath string
	projectName string

	templateVariable = map[string]string{
		"#Year#":  time.Now().Format("2006"),
		"#Month#": time.Now().Format("01"),
		"#Day#":   time.Now().Format("02"),
	}
)

func Run() {
	var rootCmd = &cobra.Command{Use: "goproject"}
	rootCmd.AddCommand(newCmd())
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func newCmd() *cobra.Command {
	cmdNew := &cobra.Command{
		Use:   "new <project name>",
		Short: "Create folder struct of new project",
		Long:  "",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			projectPath = args[0]
			projectName = path.Base(projectPath)
			if err := newProject(); err != nil {
				fmt.Println(err.Error())
			}
			fmt.Printf("done, the new project path at %s/%s\n",
				basePath, projectPath)
		},
	}
	cmdNew.Flags().StringVarP(
		&basePath, "path", "p", ".", "directory path of new project")
	return cmdNew
}

// newProject create folder struct of new project at specified path
func newProject() error {
	if basePath == "" {
		return fmt.Errorf("directory path error")
	}

	if !strings.HasPrefix(basePath, "/") {
		var err error
		basePath, err = getAbsolutePath()
		if err != nil {
			return err
		}
	}

	if err := createAllDir(); err != nil {
		return err
	}

	err := writeFiles()
	if err != nil {
		return err
	}

	cmd := exec.Command("go", "mod", "init", projectName)
	cmd.Dir = basePath + "/" + projectPath
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(string(out))

	return nil
}

// getAbsolutePath get absolute path of project
func getAbsolutePath() (string, error) {
	cur, err := os.Getwd()
	if err != nil {
		return "", err
	}
	if basePath == "." {
		return cur, nil
	}
	return filepath.Abs(path.Join(cur, path.Clean(basePath)))
}

// createAllDir creating all folder for project
func createAllDir() error {
	dirs := []string{
		"cmd/" + strings.ToLower(projectName),
		"internal/app",
		"internal/pkg/manager",
		"configs/default/dev/" + strings.ToLower(projectName),
		"pkg",
		"api",
		"_bin",
		"docs",
		"build",
		"build/default/dev/" + strings.ToLower(projectName),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(
			path.Join(basePath, projectPath, dir),
			os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}

// writeFiles write template files for project
// has @ prefix files content via https://github.com/go-bindata/go-bindata generating
// shell cmd (at root path of this project):
// 		go-bindata -o assets/assets.go -pkg=assets -ignore="\\.DS_Store|README.md|golangci_ymls" scripts/...
func writeFiles() (err error) {
	files := map[string]string{
		// Base
		"README.md":     "# " + projectName,
		"Makefile":      "@scripts/Makefile",
		".gitignore":    "@scripts/.gitignore",
		".golangci.yml": "@scripts/.golangci.yml",

		// Source
		"cmd/" + strings.ToLower(projectName) + "/main.go": "@scripts/source/cmd/#ProjectName#/main.go_",
		"internal/pkg/manager/manager.go":                  "@scripts/source/internal/pkg/manager/manager.go",

		// CI/CD
		".gitlab-ci.yml": "@scripts/.gitlab-ci.yml",
		"build/default/dev/" + strings.ToLower(projectName) + "/config.sh": "@scripts/build/default/dev/#ProjectName#/config.sh",
	}

	templateVariable["#ProjectName#"] = projectName
	templateVariable["#ProjectPath#"] = projectPath
	for fPath, content := range files {
		if strings.HasPrefix(content, "@") {
			content, err = getFileContent(content[1:])
			if err != nil {
				return err
			}
		}
		err = ioutil.WriteFile(
			path.Join(basePath, projectPath, fPath),
			[]byte(content),
			os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// getFileContent return specified file content,
// and replace all template variable for it
func getFileContent(fileName string) (string, error) {
	b, err := assets.Asset(fileName)
	if err != nil {
		return "", err
	}

	s := string(b)
	for k, v := range templateVariable {
		s = strings.ReplaceAll(s, k, v)
	}
	return s, nil
}
