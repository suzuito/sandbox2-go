package main

import (
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/mod/modfile"
	"gopkg.in/yaml.v3"
)

type Package struct {
	FullName string
	Files    []File
}

type Import struct {
	Value string
}

type File struct {
	Path    string
	Imports []Import
}

type Rule struct {
	Name      string   `yaml:"name"`
	Src       []string `yaml:"src"`
	Forbidden []string `yaml:"forbidden"`
}

func usage() {
	// Headline of usage
	fmt.Fprintf(os.Stderr, "サブプロジェクト配下のソースコードが、importのルールを満たしているかどうか？チェックするコマンド。\n")
	// Print command line option list
	flag.PrintDefaults()
}

func main() {
	ruleFilePath := ""
	flag.StringVar(&ruleFilePath, "rule-file", "", "ルールファイルのパス")
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

	// Parse go.mod
	contentGoMod, err := os.ReadFile("go.mod")
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	goModFile, err := modfile.Parse("go.mod", contentGoMod, nil)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	modName := goModFile.Module.Mod.Path
	// fmt.Println(modName)

	// Parse all .go files in this module
	packages := []Package{}
	filepath.WalkDir("./", func(dirPath string, info fs.DirEntry, _ error) error {
		if !info.IsDir() {
			return nil
		}
		fset := token.NewFileSet()
		pkgs, err := parser.ParseDir(fset, dirPath, func(fi fs.FileInfo) bool { return true }, 0)
		if err != nil {
			return err
		}
		for pkgName, pkg := range pkgs {
			fullPackageName := path.Join(modName, path.Dir(dirPath), pkgName)
			// fmt.Println(fullPackageName)
			files := []File{}
			for filePath, goFile := range pkg.Files {
				// fmt.Println("  " + filePath)
				imports := []Import{}
				for _, imp := range goFile.Imports {
					importPath := strings.Trim(imp.Path.Value, "\"")
					// fmt.Println("    " + importPath)
					imports = append(imports, Import{Value: importPath})
				}
				files = append(files, File{
					Path:    filePath,
					Imports: imports,
				})
			}
			packages = append(packages, Package{
				FullName: fullPackageName,
				Files:    files,
			})
		}
		return nil
	})

	// Validate
	errs := []InvalidRuleError{}
	for _, pkg := range packages {
		for _, rule := range rules {
			validate(&rule, &pkg, &errs)
		}
	}
	for _, err := range errs {
		fmt.Printf("%s: %s\n", err.SourcePath, err.Message)
	}
	if len(errs) > 0 {
		os.Exit(2)
	}
}

type InvalidRuleError struct {
	SourcePath string
	Message    string
}

func (t *InvalidRuleError) Error() string {
	return fmt.Sprintf("%s : %s", t.SourcePath, t.Message)
}

func validate(rule *Rule, pkg *Package, errs *[]InvalidRuleError) {
	matched := false
	for _, src := range rule.Src {
		matchedInFor, err := regexp.MatchString(src, pkg.FullName)
		// fmt.Println(src, pkg.FullName, matchedInFor)
		if err != nil {
			*errs = append(*errs, InvalidRuleError{
				SourcePath: pkg.FullName,
				Message:    fmt.Sprintf("%+v", err),
			})
			continue
		}
		if matchedInFor {
			matched = true
		}
	}
	if !matched {
		return
	}
	for _, pkgFile := range pkg.Files {
		for _, imp := range pkgFile.Imports {
			for _, forbiddenImportRule := range rule.Forbidden {
				matched, err := regexp.MatchString(forbiddenImportRule, imp.Value)
				if err != nil {
					*errs = append(*errs, InvalidRuleError{
						SourcePath: pkg.FullName,
						Message:    fmt.Sprintf("%+v", err),
					})
					continue
				}
				if matched {
					*errs = append(*errs, InvalidRuleError{
						SourcePath: pkgFile.Path,
						Message:    fmt.Sprintf("Rule:%s\n  Forbidden import %s", rule.Name, imp.Value),
					})
				}
			}
		}
	}
}
