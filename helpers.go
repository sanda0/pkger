package main

import "strings"

func IsSelected(selected map[int]struct{}, index int) bool {
	_, ok := selected[index]
	return ok
}

func LowerCasePkgName(prefix string, name string) string {
	return strings.ToLower(prefix) + strings.ToLower(name)
}

func LowerCaseText(text string) string {
	return strings.ToLower(text)
}
