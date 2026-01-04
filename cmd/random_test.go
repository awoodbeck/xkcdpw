package cmd

import (
	"strings"
	"testing"
)

func TestRunRandom(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{"no args", []string{}, false},
		{"with word count", []string{"5"}, false},
		{"with word count and number", []string{"5", "10"}, false},
		{"word count at minimum", []string{"1"}, false},
		{"word count at maximum", []string{"10"}, false},
		{"number at minimum", []string{"4", "1"}, false},
		{"number at maximum", []string{"4", "100"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// reset flags to defaults
			randomDelim = "-"
			randomCapitalize = capitalizeModeFirst
			randomNumber = false
			randomNumberMin = 0
			randomNumberMax = defaultRandomNumberMax
			randomNumberPos = numberPositionEnd

			err := runRandom(randomCmd, tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("runRandom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRunRandomWithOptions(t *testing.T) {
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
			args:       []string{"4", "1"},
			delim:      "-",
			capitalize: "none",
			wantErr:    false,
		},
		{
			name:       "with all capitalization",
			args:       []string{"4", "1"},
			delim:      " ",
			capitalize: "all",
			wantErr:    false,
		},
		{
			name:       "with random capitalization",
			args:       []string{"4", "1"},
			capitalize: "random",
			wantErr:    false,
		},
		{
			name:       "with number at beginning",
			args:       []string{"4", "1"},
			capitalize: "none",
			number:     true,
			numberMin:  100,
			numberMax:  999,
			numberPos:  "beginning",
			wantErr:    false,
		},
		{
			name:       "with number at random position",
			args:       []string{"4", "1"},
			capitalize: "none",
			number:     true,
			numberPos:  "random",
			wantErr:    false,
		},
		{
			name:       "invalid capitalize mode",
			args:       []string{"4", "1"},
			capitalize: "invalid",
			wantErr:    true,
		},
		{
			name:      "invalid number position",
			args:      []string{"4", "1"},
			number:    true,
			numberPos: "middle",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// set flags
			randomDelim = tt.delim
			randomCapitalize = tt.capitalize
			randomNumber = tt.number
			randomNumberMin = tt.numberMin
			randomNumberMax = tt.numberMax
			randomNumberPos = tt.numberPos

			err := runRandom(randomCmd, tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("runRandom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRandomValidation(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr string
	}{
		{
			name:    "word count too low",
			args:    []string{"0"},
			wantErr: "word count must fall between 1 and 10",
		},
		{
			name:    "word count too high",
			args:    []string{"11"},
			wantErr: "word count must fall between 1 and 10",
		},
		{
			name:    "number too low",
			args:    []string{"4", "0"},
			wantErr: "number must fall between 1 and 100",
		},
		{
			name:    "number too high",
			args:    []string{"4", "101"},
			wantErr: "number must fall between 1 and 100",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// reset flags to defaults
			randomDelim = "-"
			randomCapitalize = capitalizeModeFirst
			randomNumber = false

			err := runRandom(randomCmd, tt.args)
			if err == nil {
				t.Errorf("runRandom() expected error containing %q, got nil", tt.wantErr)
				return
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("runRandom() error = %v, want error containing %q", err, tt.wantErr)
			}
		})
	}
}

func TestRandomConcurrency(t *testing.T) {
	// test that running multiple random generations sequentially works correctly
	// note: CLI commands are not designed to be run concurrently with different flag values
	t.Run("sequential generations", func(t *testing.T) {
		for i := range 10 {
			randomDelim = "-"
			randomCapitalize = capitalizeModeFirst
			randomNumber = false

			err := runRandom(randomCmd, []string{"4", "1"})
			if err != nil {
				t.Errorf("runRandom() iteration %d failed: %v", i, err)
			}
		}
	})
}
