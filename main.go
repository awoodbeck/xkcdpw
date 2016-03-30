//go:generate go run -tags=dev server/assets_generate.go

package main

import (
	"fmt"
	"os"

	"github.com/jawher/mow.cli"
)

var (
	app = cli.App("xkdcpw", usage)
)

var usage = `Generates easy-to-remember passphrases.

This application is a simple passphrase generator based on the XKCD comic.

See http://www.explainxkcd.com/wiki/index.php/936:_Password_Strength for an
explanation.  This app goes a step further toward making passphrases
memorable in that it uses a series of adjectives followed by an noun (see
https://gfycat.com/about#links).`

func main() {
	exitWhenError(app.Run(os.Args))
}

func cmdError(cmd *cli.Cmd, err error) {
	fmt.Println("Error:", err)
	cmd.PrintHelp()
	os.Exit(1)
}

func exitWhenError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
