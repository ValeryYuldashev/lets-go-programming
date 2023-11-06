package main

import (
	"fmt"
	"os"

	"ex3.2.2/internal"
	contains "github.com/ValeryYuldashev/substringFinder/pkg"
	yaml "gopkg.in/yaml.v3"
)

func main() {
	file, err := os.ReadFile("data.yml")

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	var data internal.Config

	err = yaml.Unmarshal(file, &data)

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	for _, v := range data.Files {
		fmt.Println(contains.Contains(v.Filename, v.Substing))
	}
}
