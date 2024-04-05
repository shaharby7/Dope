package main

import (
	"fmt"

	"github.com/shaharby7/Dope/pkg/build"
)

// tmp

func main() {
	err := build.BuildProject("./example/project.dope.yaml", "./example/build")
	if err != nil {
		panic(err)
	}
	fmt.Printf("build finished successfully")
}
