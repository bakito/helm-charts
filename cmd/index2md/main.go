// from https://github.com/onosproject/build-tools/tree/master/build/cmd/index2md

package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"

	"github.com/spf13/viper"
	htmltemplate "html/template"
	texttemplate "text/template"
)

var (
	//go:embed yamlAppsTemplate.md
	yamlAppsTemplateMarkdown string

	//go:embed yamlAppsTemplate.html
	yamlAppsTemplateHTML string
)

// Chart :
type Chart struct {
	APIVersion  string   `yaml:"apiVersion"`
	AppVersion  string   `yaml:"appVersion"`
	Version     string   `yaml:"version"`
	Created     string   `yaml:"created"`
	Description string   `yaml:"description"`
	Urls        []string `yaml:"urls"`
}

// IndexYaml :
type IndexYaml struct {
	Title      string             `yaml:"-"`
	APIVersion string             `yaml:"apiVersion"`
	Entries    map[string][]Chart `yaml:"entries"`
	Generated  string             `yaml:"generated"`
}

/**
 * A simple application that takes the generated index.yaml and outputs it in
 * a Markdown format - usually we pipe this to README.md when in the gh-pages branch
 */
func main() {
	file := flag.String("file", "docs/index", "name of YAML file to parse (without extension or path)")
	title := flag.String("title", "bakito Helm Chart Releases", "title for the output")
	htmlout := flag.Bool("html", false, "output HTML instead of Markdown")
	flag.Parse()
	indexYaml, err := getIndexYaml(*file, *title)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to load %s.yaml %s\n", *file, err)
		if err != nil {
			return
		}
		os.Exit(1)
	}

	if !*htmlout {
		tmplAppsList, _ := htmltemplate.New("yamlAppsTemplateMarkdown").Parse(yamlAppsTemplateMarkdown)
		err = tmplAppsList.Execute(os.Stdout, indexYaml)
	} else {
		tmplAppsList, _ := texttemplate.New("yamlAppsTemplateHtml").Parse(yamlAppsTemplateHTML)
		err = tmplAppsList.Execute(os.Stdout, indexYaml)
	}

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to execute %v\n", err)
		os.Exit(1)
	}
}

func getIndexYaml(location string, title string) (IndexYaml, error) {
	indexYaml := &IndexYaml{}
	viper.SetConfigName(location)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return IndexYaml{}, err
	}

	if err := viper.Unmarshal(indexYaml); err != nil {
		return IndexYaml{}, err
	}

	indexYaml.Title = title

	return *indexYaml, nil
}
