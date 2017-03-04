package main

import "io/ioutil"

// EnvironmentData contains basic info from the environment... such as path etc
type EnvironmentData struct {
}

// ReadTemplate takes a filename and gets the content from it
func ReadTemplate(templateFile string) string {
	templateBytes, err := ioutil.ReadFile(templateFile)
	return string(templateBytes)
}
