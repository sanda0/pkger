package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func createServiceProviderFile(f folderSelectForm) {

	//  0  Controllers
	//  1 Models
	//  2 migrations
	//  3 views
	//  4 Resources
	//  5 Requests
	//  6 config
	//  7 middleware
	//  8 seeders
	//  9 lang

	content := `<?php

namespace {{.Config.Namespace}}\{{.Config.Prefix}}{{.PkgName}};

use Illuminate\Support\ServiceProvider;

class {{.Config.Prefix}}{{.PkgName}}ServiceProvider extends ServiceProvider
{
    /**
     * Register services.
     *
     * @return void
     */
    public function register()
    {
			{{if isSelected .Selected 6}}
			// Merges the configuration file 'config/config.php' into the '{{lowerCasePkgName .Config.Prefix .PkgName}}' configuration namespace.
			$this->mergeConfigFrom(__DIR__.'/config/config.php', '{{lowerCasePkgName .Config.Prefix .PkgName}}');
			{{end}}
    }

    /**
     * Bootstrap services.
     *
     * @return void
     */
    public function boot()
    {		
			// Loads routes defined in 'routes.php' file.
			$this->loadRoutesFrom(__DIR__.'/routes.php');
			{{if isSelected .Selected 2}}
			// Automatically loads migrations from the 'migrations' directory.
			$this->loadMigrationsFrom(__DIR__.'/migrations');
			{{end}}
			{{if isSelected .Selected 9}}
			// Loads translation files from the 'lang' directory with the namespace '{{lowerCasePkgName .Config.Prefix .PkgName}}'.
			$this->loadTranslationsFrom(__DIR__.'/lang', '{{lowerCasePkgName .Config.Prefix .PkgName}}');
			{{end}}
			{{if isSelected .Selected 3}}
			// Loads views from the 'views' directory with the namespace '{{lowerCasePkgName .Config.Prefix .PkgName}}'.
			$this->loadViewsFrom(__DIR__.'/views', '{{lowerCasePkgName .Config.Prefix .PkgName}}');
			{{end}}
    }
}
`
	tmpl := template.New("serviceProvider")

	// define helper function
	tmpl.Funcs(template.FuncMap{
		"isSelected":       IsSelected,
		"lowerCasePkgName": LowerCasePkgName,
	})

	//parse
	tmpl, err := tmpl.Parse(content)
	if err != nil {
		log.Fatal(err.Error())
	}

	file, err := os.Create(fmt.Sprintf("%s/%s%s/src/%s%sServiceProvider.php", f.Config.PkgRoot, f.Config.Prefix, f.PkgName, f.Config.Prefix, f.PkgName))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	err = tmpl.Execute(file, f)
	if err != nil {
		log.Fatal(err.Error())
	}

}
