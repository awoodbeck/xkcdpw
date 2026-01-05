package cmd

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"regexp"
	"strings"

	"github.com/awoodbeck/acrostic"
	"github.com/spf13/cobra"
)

const (
	defaultNumberMax     = 99
	defaultRandomWordLen = 4
	capitalizeModeNone   = "none"
	capitalizeModeFirst  = "first"
	capitalizeModeAll    = "all"
	capitalizeModeRandom = "random"
	numberPositionEnd    = "end"
	numberPositionBegin  = "beginning"
	numberPositionRandom = "random"
	commonLetters        = "abcdefghilmnoprstw"
)

var (
	delim      string
	capitalize string
	number     bool
	numberMin  int
	numberMax  int
	numberPos  string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "xkcdpw [WORD] [NUMBER]",
	Short: "Generates easy-to-remember passphrases",
	Long: `Generates easy-to-remember passphrases.

This application is a simple passphrase generator based on the XKCD comic.

See http://www.explainxkcd.com/wiki/index.php/936:_Password_Strength for an
explanation. This app goes a step further toward making passphrases
memorable in that it uses a series of adjectives followed by a noun (see
https://gfycat.com/about#links).

If no WORD is provided, a random 4-letter word is used as the acrostic.
NUMBER specifies how many passphrases to generate (default: 10).`,
	Args: cobra.MaximumNArgs(2),
	RunE: runRoot,
}

func init() {
	RootCmd.Flags().StringVarP(&delim, "delim", "d", "-", "delimiter to use between words")
	RootCmd.Flags().StringVarP(&capitalize, "capitalize", "c", capitalizeModeFirst,
		"capitalization mode: none, first, all, random")
	RootCmd.Flags().BoolVarP(&number, "number", "n", false, "add a random number to each passphrase")
	RootCmd.Flags().IntVar(&numberMin, "number-min", 0,
		"minimum value for random number (requires --number)")
	RootCmd.Flags().IntVar(&numberMax, "number-max", defaultNumberMax,
		"maximum value for random number (requires --number)")
	RootCmd.Flags().StringVar(&numberPos, "number-position", numberPositionEnd,
		"number position: end, beginning, random (requires --number)")
}

func runRoot(_ *cobra.Command, args []string) error {
	word, num, err := parseArgs(args)
	if err != nil {
		return err
	}

	acro, err := acrostic.NewAcrostic(nil, nil)
	if err != nil {
		return err
	}

	opts, err := buildOptions()
	if err != nil {
		return err
	}

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

func parseArgs(args []string) (word string, num int, err error) {
	word = ""
	num = 10

	if len(args) > 0 {
		// Check if first arg is a word or a number
		if matched, _ := regexp.MatchString("^[a-zA-Z]+$", args[0]); matched {
			word = strings.ToLower(args[0])
			// Second arg is number of passphrases if present
			if len(args) > 1 {
				if _, err := fmt.Sscanf(args[1], "%d", &num); err != nil {
					return "", 0, fmt.Errorf("invalid number argument: %w", err)
				}
			}
		} else {
			// First arg is number of passphrases
			if _, err := fmt.Sscanf(args[0], "%d", &num); err != nil {
				return "", 0, fmt.Errorf("invalid argument: must be a word (letters only) or number")
			}
		}
	}

	// If no word provided, generate a random word
	if word == "" {
		word = generateRandomWord(defaultRandomWordLen)
	}

	// Validate number
	if num < 1 || num > 100 {
		return "", 0, fmt.Errorf("number must fall between 1 and 100, inclusive")
	}

	return word, num, nil
}

func buildOptions() ([]acrostic.Option, error) {
	opts := []acrostic.Option{
		acrostic.WithSeparator(delim),
	}

	capMode, err := parseCapitalizeMode(capitalize)
	if err != nil {
		return nil, err
	}
	if capMode != acrostic.CapitalizeNone {
		opts = append(opts, acrostic.WithCapitalization(capMode))
	}

	if number {
		opts = append(opts, acrostic.WithNumber(numberMin, numberMax))

		numPosMode, parseErr := parseNumberPosition(numberPos)
		if parseErr != nil {
			return nil, parseErr
		}
		if numPosMode != acrostic.NumberPositionEnd {
			opts = append(opts, acrostic.WithNumberPosition(numPosMode))
		}
	}

	return opts, nil
}

func generateRandomWord(length int) string {
	word := make([]byte, length)
	lettersLen := big.NewInt(int64(len(commonLetters)))

	for i := range word {
		n, err := rand.Int(rand.Reader, lettersLen)
		if err != nil {
			// Fallback to a safe default if random generation fails
			word[i] = 't'
			continue
		}
		word[i] = commonLetters[n.Int64()]
	}

	return string(word)
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
