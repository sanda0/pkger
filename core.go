package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func createConfig() {
	config := config{
		PkgRoot:   "packages/pkger",
		Prefix:    "X",
		Namespace: "Pkger",
	}

	file, err := json.MarshalIndent(config, "", "	")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = ioutil.WriteFile("config.pkger.json", file, 0666)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func readConfig() (*config, error) {
	data, err := ioutil.ReadFile("config.pkger.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	congifg := config{}
	err = json.Unmarshal(data, &congifg)

	return &congifg, err
}

func checkOrCreatePkgRoot(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
	}
}

func newPkgForm(flags NewPkgFlags) {
	conifg, err := readConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	checkOrCreatePkgRoot(conifg.PkgRoot)
	showForm(&flags, conifg)
}

func createFoldersAndFiles(f folderSelectForm) {
	if !f.Quit {
		//create
		for i, folder := range f.Choices {
			_, ok := f.Selected[i]
			if ok {
				os.MkdirAll(fmt.Sprintf("%s/%s%s/src/%s", f.Config.PkgRoot, f.Config.Prefix, f.PkgName, folder), 0755)
			}
		}
		fmt.Println("✅ Folders created successfully!")

	}
}

// show form using bubbletea
func initModel(flags *NewPkgFlags, config *config) folderSelectForm {
	return folderSelectForm{
		Choices:     []string{"Controllers", "Models", "migrations", "views", "Resources", "Requests", "config", "middleware", "seeders", "lang"},
		Cursor:      0,
		Selected:    map[int]struct{}{},
		PkgName:     flags.PkgName,
		Config:      config,
		AuthorName:  flags.AuthorName,
		AuthorEmail: flags.AuthorEmail,
	}
}

func showForm(flags *NewPkgFlags, config *config) {
	form := tea.NewProgram(initModel(flags, config))
	if nf, err := form.Run(); err != nil {
		log.Fatal(err.Error())
	} else {
		createFoldersAndFiles(nf.(folderSelectForm))
		createRoutesFile(nf.(folderSelectForm))
		createServiceProviderFile(nf.(folderSelectForm))
		createComposerJsonFile(nf.(folderSelectForm))
	}

}
