package cmd

import (
	"strings"
	"testing"

	"github.com/awoodbeck/acrostic"
)

func TestParseCapitalizeMode(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    acrostic.CapitalizeMode
		wantErr bool
	}{
		{"none", "none", acrostic.CapitalizeNone, false},
		{"None uppercase", "None", acrostic.CapitalizeNone, false},
		{"first", "first", acrostic.CapitalizeFirst, false},
		{"all", "all", acrostic.CapitalizeAll, false},
		{"random", "random", acrostic.CapitalizeRandom, false},
		{"invalid", "invalid", acrostic.CapitalizeNone, true},
		{"empty", "", acrostic.CapitalizeNone, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseCapitalizeMode(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCapitalizeMode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseCapitalizeMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseNumberPosition(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    acrostic.NumberPosition
		wantErr bool
	}{
		{"end", "end", acrostic.NumberPositionEnd, false},
		{"End uppercase", "End", acrostic.NumberPositionEnd, false},
		{"beginning", "beginning", acrostic.NumberPositionBeginning, false},
		{"random", "random", acrostic.NumberPositionRandom, false},
		{"invalid", "middle", acrostic.NumberPositionEnd, true},
		{"empty", "", acrostic.NumberPositionEnd, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseNumberPosition(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseNumberPosition() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseNumberPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateRandomWord(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"4 letters", 4},
		{"1 letter", 1},
		{"10 letters", 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			word := generateRandomWord(tt.length)
			if len(word) != tt.length {
				t.Errorf("generateRandomWord() length = %d, want %d", len(word), tt.length)
			}
			// Verify all characters are lowercase letters
			for _, c := range word {
				if c < 'a' || c > 'z' {
					t.Errorf("generateRandomWord() contains non-letter character: %c", c)
				}
			}
		})
	}
}

func TestRunRoot(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{"no args", []string{}, false},
		{"just number", []string{"5"}, false},
		{"word only", []string{"test"}, false},
		{"word and number", []string{"test", "5"}, false},
		{"uppercase word", []string{"TEST"}, false},
		{"invalid word with numbers", []string{"test123"}, true},
		{"invalid word with special chars", []string{"test!"}, true},
		{"number too low", []string{"test", "0"}, true},
		{"number too high", []string{"test", "101"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset flags to defaults
			delim = "-"
			capitalize = capitalizeModeFirst
			number = false
			numberMin = 0
			numberMax = defaultNumberMax
			numberPos = numberPositionEnd

			err := runRoot(RootCmd, tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("runRoot() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRunRootWithOptions(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		delim      string
		capitalize string
		number     bool
		numberMin  int
		numberMax  int
		numberPos  string
		wantErr    bool
	}{
		{
			name:       "with delimiter",
			args:       []string{"test", "1"},
			delim:      "_",
			capitalize: capitalizeModeFirst,
			wantErr:    false,
		},
		{
			name:       "with all capitalization",
			args:       []string{"test", "1"},
			delim:      "-",
			capitalize: capitalizeModeAll,
			wantErr:    false,
		},
		{
			name:       "with number",
			args:       []string{"test", "1"},
			delim:      "-",
			capitalize: capitalizeModeFirst,
			number:     true,
			numberMin:  10,
			numberMax:  99,
			numberPos:  numberPositionEnd,
			wantErr:    false,
		},
		{
			name:       "random word with options",
			args:       []string{"5"},
			delim:      " ",
			capitalize: capitalizeModeNone,
			wantErr:    false,
		},
		{
			name:       "invalid capitalize mode",
			args:       []string{"test", "1"},
			capitalize: "invalid",
			wantErr:    true,
		},
		{
			name:      "invalid number position",
			args:      []string{"test", "1"},
			number:    true,
			numberPos: "invalid",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set flags
			delim = tt.delim
			capitalize = tt.capitalize
			number = tt.number
			numberMin = tt.numberMin
			numberMax = tt.numberMax
			numberPos = tt.numberPos

			err := runRoot(RootCmd, tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("runRoot() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRunRootValidation(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr string
	}{
		{
			name:    "number too low",
			args:    []string{"test", "0"},
			wantErr: "number must fall between 1 and 100",
		},
		{
			name:    "number too high",
			args:    []string{"test", "101"},
			wantErr: "number must fall between 1 and 100",
		},
		{
			name:    "word with digits",
			args:    []string{"test123"},
			wantErr: "invalid argument",
		},
		{
			name:    "word with special chars",
			args:    []string{"test@example"},
			wantErr: "invalid argument",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset flags to defaults
			delim = "-"
			capitalize = capitalizeModeFirst
			number = false

			err := runRoot(RootCmd, tt.args)
			if err == nil {
				t.Errorf("runRoot() expected error containing %q, got nil", tt.wantErr)
				return
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("runRoot() error = %v, want error containing %q", err, tt.wantErr)
			}
		})
	}
}
