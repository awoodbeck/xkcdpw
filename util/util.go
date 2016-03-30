package util

import (
	"fmt"
	"os"

	"github.com/jawher/mow.cli"
)

// CmdError accepts a mow.cli Cmd pointer and an error object.  It then
// displays the error and prints the command's help.
func CmdError(cmd *cli.Cmd, err error) {
	fmt.Println("Error:", err)
	cmd.PrintHelp()
	os.Exit(1)
}

// ExitWhenError accepts an error and only exits if the error is not nil.
func ExitWhenError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
