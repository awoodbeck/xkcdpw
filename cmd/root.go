package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "xkcdpw",
	Short: "Generates easy-to-remember passphrases",
	Long: `Generates easy-to-remember passphrases.

This application is a simple passphrase generator based on the XKCD comic.

See http://www.explainxkcd.com/wiki/index.php/936:_Password_Strength for an
explanation. This app goes a step further toward making passphrases
memorable in that it uses a series of adjectives followed by a noun (see
https://gfycat.com/about#links).`,
}
