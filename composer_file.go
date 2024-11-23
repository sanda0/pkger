package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func createComposerJsonFile(f folderSelectForm) {
	content := `{
		"name": "{{lowerCaseTest .Config.Namespace}}/{{lowerCasePkgName .Config.Prefix .PkgName}}",
		"type": "project",
		"license": "MIT",
		"autoload": {
				"psr-4": {
						"{{.Config.Namespace}}\\{{.Config.Prefix}}{{.PkgName}}\\": "src/"
				}
		},
		"authors": [
				{
						"name": "{{.AuthorName}}",
						"email": "{{.AuthorEmail}}"
				}
		],
		"minimum-stability": "dev",
		"require": {},
		"extra": {
				"laravel": {
						"providers": [
								"{{.Config.Namespace}}\\{{.Config.Prefix}}{{.PkgName}}\\{{.Config.Prefix}}{{.PkgName}}ServiceProvider"
						]
				}
		}
	}
	`

	tmpl := template.New("composer.json")
	tmpl.Funcs(template.FuncMap{
		"isSelected":       IsSelected,
		"lowerCasePkgName": LowerCasePkgName,
		"lowerCaseText":    LowerCaseText,
	})

	tmpl, err := tmpl.Parse(content)
	if err != nil {
		log.Fatal(err.Error())
	}

	file, err := os.Create(fmt.Sprintf("%s/%s%s/composer.json", f.Config.PkgRoot, f.Config.Prefix, f.PkgName))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	err = tmpl.Execute(file, f)
	if err != nil {
		log.Fatal(err.Error())
	}

}
