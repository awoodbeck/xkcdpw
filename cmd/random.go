package cmd

import (
	"fmt"

	"github.com/awoodbeck/acrostic"
	"github.com/spf13/cobra"
)

const (
	defaultRandomNumberMax = 99
)

var (
	randomDelim      string
	randomCapitalize string
	randomNumber     bool
	randomNumberMin  int
	randomNumberMax  int
	randomNumberPos  string
)

var randomCmd = &cobra.Command{
	Use:   "random [WORDS] [NUMBER]",
	Short: "Generate random acrostical passphrases",
	Long: `Generate random acrostical passphrases.

WORDS specifies words per passphrase (1-10, default: 4).
NUMBER specifies how many passphrases to generate (1-100, default: 10).`,
	Args: cobra.RangeArgs(0, 2),
	RunE: runRandom,
}

func init() {
	RootCmd.AddCommand(randomCmd)

	randomCmd.Flags().StringVarP(&randomDelim, "delim", "d", "-", "delimiter to use between words")
	randomCmd.Flags().StringVarP(&randomCapitalize, "capitalize", "c", capitalizeModeFirst,
		"capitalization mode: none, first, all, random")
	randomCmd.Flags().BoolVarP(&randomNumber, "number", "n", false, "add a random number to each passphrase")
	randomCmd.Flags().IntVar(&randomNumberMin, "number-min", 0,
		"minimum value for random number (requires --number)")
	randomCmd.Flags().IntVar(&randomNumberMax, "number-max", defaultRandomNumberMax,
		"maximum value for random number (requires --number)")
	randomCmd.Flags().StringVar(&randomNumberPos, "number-position", numberPositionEnd,
		"number position: end, beginning, random (requires --number)")
}

func runRandom(_ *cobra.Command, args []string) error {
	wordCount, num, err := parseRandomArgs(args)
	if err != nil {
		return err
	}

	acro, err := acrostic.NewAcrostic(nil, nil)
	if err != nil {
		return err
	}

	opts, err := buildRandomOptions()
	if err != nil {
		return err
	}

	phrases, err := acro.GenerateRandomAcrostics(wordCount, num, opts...)
	if err != nil {
		return err
	}

	fmt.Println()
	for _, p := range phrases {
		fmt.Printf("%s\n\n", p)
	}

	return nil
}

func parseRandomArgs(args []string) (wordCount, num int, err error) {
	wordCount = 4
	num = 10

	if len(args) > 0 {
		if _, err := fmt.Sscanf(args[0], "%d", &wordCount); err != nil {
			return 0, 0, fmt.Errorf("invalid word count argument: %w", err)
		}
	}

	if len(args) > 1 {
		if _, err := fmt.Sscanf(args[1], "%d", &num); err != nil {
			return 0, 0, fmt.Errorf("invalid number argument: %w", err)
		}
	}

	if wordCount < 1 || wordCount > 10 {
		return 0, 0, fmt.Errorf("word count must fall between 1 and 10, inclusive")
	}

	if num < 1 || num > 100 {
		return 0, 0, fmt.Errorf("number must fall between 1 and 100, inclusive")
	}

	return wordCount, num, nil
}

func buildRandomOptions() ([]acrostic.Option, error) {
	opts := []acrostic.Option{
		acrostic.WithSeparator(randomDelim),
	}

	capMode, err := parseCapitalizeMode(randomCapitalize)
	if err != nil {
		return nil, err
	}
	if capMode != acrostic.CapitalizeNone {
		opts = append(opts, acrostic.WithCapitalization(capMode))
	}

	if randomNumber {
		opts = append(opts, acrostic.WithNumber(randomNumberMin, randomNumberMax))

		numPos, parseErr := parseNumberPosition(randomNumberPos)
		if parseErr != nil {
			return nil, parseErr
		}
		if numPos != acrostic.NumberPositionEnd {
			opts = append(opts, acrostic.WithNumberPosition(numPos))
		}
	}

	return opts, nil
}
