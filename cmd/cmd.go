package cmd

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/awoodbeck/acrostic"
	"github.com/awoodbeck/xkcdpw/util"
	"github.com/jawher/mow.cli"
)

// ErrInvalidNum is returned when a given number is less than 1 or greater than 100.
var ErrInvalidNum = errors.New("number must fall between 1 and 100, inclusive")

// ErrInvalidWord is returned when a given WORD contains characters other than alpha characters.
var ErrInvalidWord = errors.New("WORD must only contain letters")

// ErrInvalidWordCount is returned when a given word count is less than 1 or greater than 10.
var ErrInvalidWordCount = errors.New("word count must fall between 1 and 10, inclusive")

// RegisterAcrostic accepts a pointer to a mow.cli Cli object and registers the "acrostic" command.
func RegisterAcrostic(app *cli.Cli) {
	app.Command("acrostic", "generate acrostical passphrases from a given word", func(cmd *cli.Cmd) {
		defaultWord := ""
		defaultNum := 10

		word := cmd.StringArg("WORD", defaultWord, "case-insensitive word from which to generate acrostical passphrases")
		num := cmd.IntArg("NUMBER", defaultNum, "number of passphrases to generate, up to 100")
		delim := cmd.StringOpt("d delim", " ", "delimiter to use between words")

		cmd.Spec = "[OPTIONS] WORD [NUMBER]"

		cmd.Before = func() {
			matched, err := regexp.MatchString("^[a-zA-Z]+$", *word)
			util.ExitWhenError(err)

			if !matched {
				*word = defaultWord
				util.CmdError(cmd, ErrInvalidWord)
			}

			if *num < 1 || *num > 100 {
				*num = defaultNum
				util.CmdError(cmd, ErrInvalidNum)
			}

			*word = strings.ToLower(*word)
		}

		cmd.Action = func() {
			var phrases []string

			a, err := acrostic.NewAcrostic(nil, nil)
			util.ExitWhenError(err)

			phrases, err = a.GenerateAcrostics(*word, *num)

			fmt.Println()
			for _, p := range phrases {
				fmt.Printf("%s\n\n", strings.Join(strings.Split(p, " "), *delim))
			}
		}
	})
}

// RegisterRandom accepts a pointer to a mow.cli Cli object and registers the "random" command.
func RegisterRandom(app *cli.Cli) {
	app.Command("random", "generate random acrostical passphrases", func(cmd *cli.Cmd) {
		defaultWordCount := 4
		defaultNum := 10
		count := cmd.IntArg("WORDS", defaultWordCount, "words per passphrase, up to 10")
		num := cmd.IntArg("NUMBER", defaultNum, "number of passphrases to generate, up to 100")
		delim := cmd.StringOpt("d delim", " ", "delimiter to use between words")

		cmd.Spec = "[OPTIONS] [WORDS] [NUMBER]"

		cmd.Before = func() {
			if *count < 1 || *num > 10 {
				*count = defaultWordCount
				util.CmdError(cmd, ErrInvalidWordCount)
			}

			if *num < 1 || *num > 100 {
				*num = defaultNum
				util.CmdError(cmd, ErrInvalidNum)
			}
		}

		cmd.Action = func() {
			var phrases []string

			a, err := acrostic.NewAcrostic(nil, nil)
			util.ExitWhenError(err)

			fmt.Println()

			for i := 0; i < *num; i++ {
				phrases, err = a.GenerateRandomAcrostics(*count, 1)
				for _, p := range phrases {
					fmt.Printf("%s\n\n", strings.Join(strings.Split(p, " "), *delim))
				}
			}
		}
	})
}

// RegisterServer accepts a pointer to a mow.cli Cli object and registers the "server" command.
func RegisterServer(app *cli.Cli) {
	app.Command("server", "run a web server capable of generating acrostical phrases", func(cmd *cli.Cmd) {

	})
}
