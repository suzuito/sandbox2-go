package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"

	"gopkg.in/yaml.v3"
)

func usage() {
	// Headline of usage
	fmt.Fprintf(os.Stderr, "サブプロジェクトがディレクトリ構造のルールを満たしているかどうか？チェックするコマンド。\n")
	// Print command line option list
	flag.PrintDefaults()
}

// サブプロジェクト配下のソースコードが、ディレクトリ構造のルールを満たしているかどうか？チェックするコマンド。
func main() {
	ruleFilePath := ""
	flag.StringVar(&ruleFilePath, "rule-file", "", "ルールファイルのパス")
	flag.Usage = usage
	flag.Parse()
	if ruleFilePath == "" {
		usage()
		os.Exit(1)
	}
	ruleFileBytes, err := os.ReadFile(ruleFilePath)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(2)
	}
	rules := []Rule{}
	if err := yaml.Unmarshal(ruleFileBytes, &rules); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(2)
	}
	errs := []*InvalidRuleError{}
	filepath.WalkDir("./", func(dirPath string, info fs.DirEntry, _ error) error {
		if !info.IsDir() {
			return nil
		}
		for _, rule := range rules {
			matched, err := regexp.MatchString(rule.Path, dirPath)
			if err != nil {
				errs = append(errs, &InvalidRuleError{
					Path:    dirPath,
					Message: fmt.Sprintf("%+v", err),
				})
				continue
			}
			if !matched {
				continue
			}
			entries, err := os.ReadDir(dirPath)
			if err != nil {
				errs = append(errs, &InvalidRuleError{
					Path:    dirPath,
					Message: fmt.Sprintf("%+v", err),
				})
				continue
			}
			for _, entry := range entries {
				if !entry.IsDir() {
					continue
				}
				matched3 := false
				for _, allowed := range rule.Allowed {
					matched2, err := regexp.MatchString(allowed, entry.Name())
					if err != nil {
						errs = append(errs, &InvalidRuleError{
							Path:    dirPath,
							Message: fmt.Sprintf("%+v", err),
						})
						continue
					}
					if matched2 {
						matched3 = true
					}
				}
				if !matched3 {
					errs = append(errs, &InvalidRuleError{
						Path:    dirPath,
						Message: fmt.Sprintf("Rule:%s\n  Forbidden dir '%s'", rule.Name, entry.Name()),
					})
				}
			}
		}
		return nil
	})
	for _, err := range errs {
		fmt.Printf("%+v\n", err)
	}
	if len(errs) > 0 {
		os.Exit(2)
	}
	/*
		if err := readDir("./"); err != nil {
			fmt.Printf("%+v\n", err)
			os.Exit(1)
		}
	*/
}

type InvalidRuleError struct {
	Path    string
	Message string
}

func (t *InvalidRuleError) Error() string {
	return fmt.Sprintf("%s : %s", t.Path, t.Message)
}

type Rule struct {
	Name    string   `yaml:"name"`
	Path    string   `yaml:"path"`
	Allowed []string `yaml:"allowed"`
}
