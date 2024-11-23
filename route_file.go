package main

import (
	"fmt"
	"log"
	"os"
)

func createRoutesFile(f folderSelectForm) {
	content := `<?php

use Illuminate\Support\Facades\Route;
`

	file, err := os.Create(fmt.Sprintf("%s/%s%s/src/routes.php", f.Config.PkgRoot, f.Config.Prefix, f.PkgName))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err.Error())
	}
}
