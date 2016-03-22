// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/awoodbeck/acrostic/assets"
	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(assets.Assets, vfsgen.Options{
		BuildTags:    "!dev",
		PackageName:  "assets",
		VariableName: "Assets",
		Filename:     "assets/vfsgen.go",
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
