package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/awoodbeck/acrostic"
	"github.com/jawher/mow.cli"
)

// ErrInvalidWord is returned when a given WORD contains characters other than alpha characters.
var ErrInvalidWord = errors.New("WORD must only contain letters")

// ErrInvalidNum is returned when a given number is less than 1 or greater than 100.
var ErrInvalidNum = errors.New("number must fall between 1 and 100, inclusive")

func init() {
	app.Command("acrostic", "generate acrostical passphrases from a given word", func(cmd *cli.Cmd) {
		defaultWord := ""
		defaultNum := 10

		word := cmd.StringArg("WORD", defaultWord, "case-insensitive word from which to generate acrostical passphrases")
		num := cmd.IntArg("NUMBER", defaultNum, "number of passphrases to generate, up to 100")
		delim := cmd.StringOpt("d delim", " ", "delimiter to use between words")

		cmd.Spec = "[OPTIONS] WORD [NUMBER]"

		cmd.Before = func() {
			matched, err := regexp.MatchString("^[a-zA-Z]+$", *word)
			exitWhenError(err)

			if !matched {
				*word = defaultWord
				cmdError(cmd, ErrInvalidWord)
			}

			if *num < 1 || *num > 100 {
				*num = defaultNum
				cmdError(cmd, ErrInvalidNum)
			}

			*word = strings.ToLower(*word)
		}

		cmd.Action = func() {
			var phrases []string

			acro, err := acrostic.NewAcrostic(nil, nil)
			exitWhenError(err)

			phrases, err = acro.GenerateAcrostics(*word, *num)

			fmt.Println()
			for _, p := range phrases {
				fmt.Printf("%s\n\n", strings.Join(strings.Split(p, " "), *delim))
			}
		}
	})
}
