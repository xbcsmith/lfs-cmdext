// Copyright © 2019 Brett Smith <xbcsmith@gmail.com>, . All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"

	"flag"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/yaml.v3"
)

// PackageInformation struct for packageinformation
type PackageInformation struct {
	Commands     []Command    `json:"commands" yaml:"commands"`
	Dependencies Dependencies `json:"dependencies" yaml:"dependencies"`
	Description  string       `json:"description" yaml:"description"`
	Name         string       `json:"name" yaml:"name"`
	Sources      []Source     `json:"sources" yaml:"sources"`
	Version      string       `json:"version" yaml:"version"`
}

// Command struct for command
type Command struct {
	Cmd   string `json:"cmd" yaml:"cmd"`
	Index int    `json:"index" yaml:"index"`
}

// Dependencies struct for dependencies
type Dependencies struct {
	Optional    []string `json:"optional" yaml:"optional"`
	Recommended []string `json:"recommended" yaml:"recommended"`
	Requires    []string `json:"requires" yaml:"requires"`
}

// Source struct for source
type Source struct {
	Archive   string `json:"archive" yaml:"archive"`
	BuildTime string `json:"build_time" yaml:"build_time"`
	MD5Sum    string `json:"md5sum" yaml:"md5sum"`
	OnDisk    string `json:"ondisk" yaml:"ondisk"`
	Size      string `json:"size" yaml:"size"`
}

// Application struct for application
type Application struct {
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
	Version     string `json:"version" yaml:"version"`
}

// ToYAML func takes no input and returns []byte, error
func (p *PackageInformation) ToYAML() ([]byte, error) {
	content, err := yaml.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to yaml : %v", err)
	}
	return content, nil
}

// ToJSON func takes no input and returns []byte, error
func (p *PackageInformation) ToJSON() ([]byte, error) {
	content, err := json.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to json : %v", err)
	}
	return content, nil
}

// ToPrettyJSON func takes no input and returns []byte, error
func (p *PackageInformation) ToPrettyJSON() ([]byte, error) {
	content, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to convert to json : %v", err)
	}
	return content, nil
}

// end func takes string splits it on separator and returns string
func end(s, sep string) string {
	ss := strings.Split(s, sep)
	return ss[len(ss)-1]
}

// begin func takes string splits it on separator and returns string
func begin(s, sep string) string {
	ss := strings.Split(s, sep)
	ss = ss[:len(ss)-1]
	return strings.Join(ss, sep)
}

// ExtractCommands func takes doc *goquery.Document input and returns []Command, error
func ExtractCommands(doc *goquery.Document) ([]Command, error) {
	commands := make([]Command, 0)
	var index int
	doc.Find("kbd").Each(func(i int, s *goquery.Selection) {
		command := Command{
			Index: index,
			Cmd:   s.Text(),
		}
		commands = append(commands, command)
		index = index + 1
	})
	return commands, nil
}

// ExtractSources func takes doc *goquery.Document input and returns []Source, error
func ExtractSources(doc *goquery.Document) ([]Source, error) {
	sources := make([]Source, 0)
	doc.Find(".package .itemizedlist .compact").Each(func(i int, s *goquery.Selection) {
		source := Source{}
		s.Find("p").Each(func(i int, s *goquery.Selection) {
			block := s.Text()
			switch true {
			case strings.Contains(block, "(HTTP)"):
				link := s.Find(".ulink").Text()
				source.Archive = strings.TrimSpace(link)
			case strings.Contains(block, "MD5"):
				md5 := strings.Split(block, ":")[1]
				source.MD5Sum = strings.TrimSpace(md5)
			case strings.Contains(block, "size"):
				size := strings.Split(block, ":")[1]
				source.Size = strings.TrimSpace(size)
			case strings.Contains(block, "disk space"):
				ondisk := strings.Split(block, ":")[1]
				source.OnDisk = strings.TrimSpace(ondisk)
			case strings.Contains(block, "build time"):
				bt := strings.Split(block, ":")[1]
				source.BuildTime = strings.TrimSpace(bt)
			}

		})
		sources = append(sources, source)
	})
	return sources, nil

}

// ExtractApplication func takes doc *goquery.Document input and returns Application, error
func ExtractApplication(doc *goquery.Document) (Application, error) {
	application := Application{}
	doc.Find("title").Each(func(i int, s *goquery.Selection) {
		titlestr := s.Text()
		if titlestr != "" {
			title := strings.TrimSpace(titlestr)
			name := begin(title, "-")
			version := end(title, "-")
			if strings.Contains(title, " ") {
				name = strings.Join(strings.Split(title, " "), "-")
				version = "1.0.0"
			}
			if strings.Contains(title, "&nbsp;") {
				t := end(title, "&nbsp;")
				name = begin(t, "-")
				version = end(t, "-")
			}
			title = strconv.QuoteToASCII(strings.TrimSpace(titlestr))
			if strings.Contains(title, "\\u00a0") {
				t := end(title, "\\u00a0")
				name = begin(t, "-")
				version = end(t, "-")
				if strings.HasSuffix(version, `"`) {
					version = version[:len(version)-1]
				}
			}
			application.Name = strings.ToLower(strings.TrimSpace(name))
			application.Version = strings.TrimSpace(version)
		}
	})

	doc.Find(".package").Each(func(i int, s *goquery.Selection) {
		if application.Name == "" {
			name := s.Find(".application").First().Text()
			application.Name = strings.TrimSpace(strings.ToLower(name))
		}
		description := s.Find("p").First().Text()
		application.Description = strings.TrimSpace(description)
	})

	return application, nil

}

// ReadDoc func takes b []byte input and returns *goquery.Document, error
func ReadDoc(b []byte) (*goquery.Document, error) {
	p := bytes.NewReader(b)
	doc, err := goquery.NewDocumentFromReader(p)
	if err != nil {
		return doc, err
	}
	return doc, nil
}

// ExtractDependencies func takes doc *goquery.Document input and returns Dependencies, error
func ExtractDependencies(doc *goquery.Document) (Dependencies, error) {
	dependencies := Dependencies{}
	var requires []string
	var recommended []string
	var optional []string
	doc.Find(".package .required").Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(i int, a *goquery.Selection) {
			requires = append(requires, strings.ToLower(strings.TrimSpace(a.Text())))
		})
	})
	doc.Find(".package .recommended").Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(i int, a *goquery.Selection) {
			recommended = append(recommended, strings.ToLower(strings.TrimSpace(a.Text())))
		})
	})
	doc.Find(".package .optional").Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(i int, a *goquery.Selection) {
			optional = append(optional, strings.ToLower(strings.TrimSpace(a.Text())))
		})
	})
	dependencies.Requires = requires
	dependencies.Recommended = recommended
	dependencies.Optional = optional
	return dependencies, nil

}

// CreatePackageInformation func takes b []byte input and returns *PackageInformation, error
func CreatePackageInformation(b []byte) (*PackageInformation, error) {
	pkgInfo := &PackageInformation{}
	doc, err := ReadDoc(b)
	if err != nil {
		return pkgInfo, err
	}
	deps, err := ExtractDependencies(doc)
	if err != nil {
		return pkgInfo, err
	}
	pkgInfo.Dependencies = deps
	cmds, err := ExtractCommands(doc)
	if err != nil {
		return pkgInfo, err
	}
	pkgInfo.Commands = cmds
	srcs, err := ExtractSources(doc)
	if err != nil {
		return pkgInfo, err
	}
	pkgInfo.Sources = srcs
	app, err := ExtractApplication(doc)
	if err != nil {
		return pkgInfo, err
	}
	pkgInfo.Name = app.Name
	pkgInfo.Version = app.Version
	pkgInfo.Description = app.Description
	return pkgInfo, nil
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

// main func takes no input and returns
func main() {
	var (
		destdir  string
		asjson   bool
		asyaml   bool
		noindent bool
		write    bool
		debug    bool
	)
	flag.StringVar(&destdir, "destination", "/tmp/pkgs", "Path to write files to disk")
	flag.BoolVar(&asjson, "asjson", false, "Output JSON")
	flag.BoolVar(&asyaml, "asyaml", true, "Output YAML")
	flag.BoolVar(&noindent, "noindent", false, "No Indent for JSON")
	flag.BoolVar(&write, "write-to-disk", false, "Write files to disk")
	flag.BoolVar(&debug, "debug", false, "Turn debugging on")
	flag.Parse()
	args := flag.Args()
	if asjson {
		asyaml = false
	}
	if len(args) > 0 {
		if write {
			err := os.MkdirAll(destdir, 0755)
			check(err)
		}
		for _, filepath := range args {
			b, err := ioutil.ReadFile(filepath)
			check(err)
			pkgInfo, err := CreatePackageInformation(b)
			check(err)
			if asyaml {
				yml, err := pkgInfo.ToYAML()
				check(err)
				if write {
					filename := pkgInfo.Name + "-" + pkgInfo.Version + ".yml"
					filepath := path.Join(destdir, filename)
					err := ioutil.WriteFile(filepath, yml, 0644)
					check(err)
				} else {
					fmt.Printf("%s\n", yml)
				}
			}
			if asjson {
				jsn, err := pkgInfo.ToPrettyJSON()
				check(err)
				if noindent {
					jsn, err = pkgInfo.ToJSON()
					check(err)
				}
				if write {
					filename := pkgInfo.Name + "-" + pkgInfo.Version + ".json"
					filepath := path.Join(destdir, filename)
					err := ioutil.WriteFile(filepath, jsn, 0644)
					check(err)
				} else {
					fmt.Printf("%s\n", jsn)
				}
			}
		}
	}
}
