// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/awoodbeck/xkcdpw/server"
	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(server.Assets, vfsgen.Options{
		BuildTags:    "!dev",
		PackageName:  "server",
		VariableName: "Assets",
		Filename:     "server/assets/vfsgen.go",
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
