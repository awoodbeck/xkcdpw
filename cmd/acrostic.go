package cmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/awoodbeck/acrostic"
	"github.com/spf13/cobra"
)

const (
	defaultNumberMax     = 99
	capitalizeModeNone   = "none"
	capitalizeModeFirst  = "first"
	capitalizeModeAll    = "all"
	capitalizeModeRandom = "random"
	numberPositionEnd    = "end"
	numberPositionBegin  = "beginning"
	numberPositionRandom = "random"
)

var (
	acrosticDelim      string
	acrosticCapitalize string
	acrosticNumber     bool
	acrosticNumberMin  int
	acrosticNumberMax  int
	acrosticNumberPos  string
)

var acrosticCmd = &cobra.Command{
	Use:   "acrostic WORD [NUMBER]",
	Short: "Generate acrostical passphrases from a given word",
	Long: `Generate acrostical passphrases from a given word.

The WORD argument must contain only letters (case-insensitive).
NUMBER specifies how many passphrases to generate (1-100, default: 10).`,
	Args: cobra.RangeArgs(1, 2),
	RunE: runAcrostic,
}

func init() {
	RootCmd.AddCommand(acrosticCmd)

	acrosticCmd.Flags().StringVarP(&acrosticDelim, "delim", "d", "-", "delimiter to use between words")
	acrosticCmd.Flags().StringVarP(&acrosticCapitalize, "capitalize", "c", capitalizeModeFirst,
		"capitalization mode: none, first, all, random")
	acrosticCmd.Flags().BoolVarP(&acrosticNumber, "number", "n", false, "add a random number to each passphrase")
	acrosticCmd.Flags().IntVar(&acrosticNumberMin, "number-min", 0,
		"minimum value for random number (requires --number)")
	acrosticCmd.Flags().IntVar(&acrosticNumberMax, "number-max", defaultNumberMax,
		"maximum value for random number (requires --number)")
	acrosticCmd.Flags().StringVar(&acrosticNumberPos, "number-position", numberPositionEnd,
		"number position: end, beginning, random (requires --number)")
}

func runAcrostic(_ *cobra.Command, args []string) error {
	word := strings.ToLower(args[0])
	num := 10

	if len(args) > 1 {
		if _, err := fmt.Sscanf(args[1], "%d", &num); err != nil {
			return fmt.Errorf("invalid number argument: %w", err)
		}
	}

	// validate word
	matched, err := regexp.MatchString("^[a-zA-Z]+$", word)
	if err != nil {
		return err
	}
	if !matched {
		return fmt.Errorf("WORD must only contain letters")
	}

	// validate number
	if num < 1 || num > 100 {
		return fmt.Errorf("number must fall between 1 and 100, inclusive")
	}

	// create acrostic generator
	acro, err := acrostic.NewAcrostic(nil, nil)
	if err != nil {
		return err
	}

	// build options
	opts := []acrostic.Option{
		acrostic.WithSeparator(acrosticDelim),
	}

	// add capitalization option
	capMode, err := parseCapitalizeMode(acrosticCapitalize)
	if err != nil {
		return err
	}
	if capMode != acrostic.CapitalizeNone {
		opts = append(opts, acrostic.WithCapitalization(capMode))
	}

	// add number options
	if acrosticNumber {
		opts = append(opts, acrostic.WithNumber(acrosticNumberMin, acrosticNumberMax))

		numPos, parseErr := parseNumberPosition(acrosticNumberPos)
		if parseErr != nil {
			return parseErr
		}
		if numPos != acrostic.NumberPositionEnd {
			opts = append(opts, acrostic.WithNumberPosition(numPos))
		}
	}

	// generate passphrases
	phrases, err := acro.GenerateAcrostics(word, num, opts...)
	if err != nil {
		return err
	}

	fmt.Println()
	for _, p := range phrases {
		fmt.Printf("%s\n\n", p)
	}

	return nil
}

func parseCapitalizeMode(mode string) (acrostic.CapitalizeMode, error) {
	switch strings.ToLower(mode) {
	case capitalizeModeNone:
		return acrostic.CapitalizeNone, nil
	case capitalizeModeFirst:
		return acrostic.CapitalizeFirst, nil
	case capitalizeModeAll:
		return acrostic.CapitalizeAll, nil
	case capitalizeModeRandom:
		return acrostic.CapitalizeRandom, nil
	default:
		return acrostic.CapitalizeNone, fmt.Errorf("invalid capitalize mode: %s (must be none, first, all, or random)", mode)
	}
}

func parseNumberPosition(pos string) (acrostic.NumberPosition, error) {
	switch strings.ToLower(pos) {
	case numberPositionEnd:
		return acrostic.NumberPositionEnd, nil
	case numberPositionBegin:
		return acrostic.NumberPositionBeginning, nil
	case numberPositionRandom:
		return acrostic.NumberPositionRandom, nil
	default:
		return acrostic.NumberPositionEnd, fmt.Errorf("invalid number position: %s (must be end, beginning, or random)", pos)
	}
}
