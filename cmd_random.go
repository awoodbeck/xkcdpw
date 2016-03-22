package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/awoodbeck/acrostic"
	"github.com/jawher/mow.cli"
)

// ErrInvalidWordCount is returned when a given word count is less than 1 or greater than 10.
var ErrInvalidWordCount = errors.New("word count must fall between 1 and 10, inclusive")

func init() {
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
				cmdError(cmd, ErrInvalidWordCount)
			}

			if *num < 1 || *num > 100 {
				*num = defaultNum
				cmdError(cmd, ErrInvalidNum)
			}
		}

		cmd.Action = func() {
			var phrases []string

			acro, err := acrostic.NewAcrostic(nil, nil)
			exitWhenError(err)

			fmt.Println()

			for i := 0; i < *num; i++ {
				phrases, err = acro.GenerateRandomAcrostics(*count, 1)
				for _, p := range phrases {
					fmt.Printf("%s\n\n", strings.Join(strings.Split(p, " "), *delim))
				}
			}
		}
	})
}
