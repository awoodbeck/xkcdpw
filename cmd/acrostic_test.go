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

func TestRunAcrostic(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{"valid word", []string{"test"}, false},
		{"valid word with number", []string{"test", "5"}, false},
		{"invalid word with numbers", []string{"test123"}, true},
		{"invalid word with special chars", []string{"test!"}, true},
		{"valid uppercase", []string{"TEST"}, false},
		{"empty word", []string{""}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// reset flags to defaults
			acrosticDelim = "-"
			acrosticCapitalize = capitalizeModeFirst
			acrosticNumber = false
			acrosticNumberMin = 0
			acrosticNumberMax = defaultNumberMax
			acrosticNumberPos = numberPositionEnd

			err := runAcrostic(acrosticCmd, tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("runAcrostic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRunAcrosticWithOptions(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		delim       string
		capitalize  string
		number      bool
		numberMin   int
		numberMax   int
		numberPos   string
		wantErr     bool
		checkOutput func(string) bool
	}{
		{
			name:       "with delimiter",
			args:       []string{"test", "1"},
			delim:      "-",
			capitalize: "none",
			wantErr:    false,
		},
		{
			name:       "with first capitalization",
			args:       []string{"test", "1"},
			delim:      " ",
			capitalize: "first",
			wantErr:    false,
		},
		{
			name:       "with number",
			args:       []string{"test", "1"},
			delim:      " ",
			capitalize: "none",
			number:     true,
			numberMin:  10,
			numberMax:  99,
			numberPos:  "end",
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
			// set flags
			acrosticDelim = tt.delim
			acrosticCapitalize = tt.capitalize
			acrosticNumber = tt.number
			acrosticNumberMin = tt.numberMin
			acrosticNumberMax = tt.numberMax
			acrosticNumberPos = tt.numberPos

			err := runAcrostic(acrosticCmd, tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("runAcrostic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAcrosticValidation(t *testing.T) {
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
			wantErr: "WORD must only contain letters",
		},
		{
			name:    "word with special chars",
			args:    []string{"test@example"},
			wantErr: "WORD must only contain letters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// reset flags to defaults
			acrosticDelim = "-"
			acrosticCapitalize = capitalizeModeFirst
			acrosticNumber = false

			err := runAcrostic(acrosticCmd, tt.args)
			if err == nil {
				t.Errorf("runAcrostic() expected error containing %q, got nil", tt.wantErr)
				return
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("runAcrostic() error = %v, want error containing %q", err, tt.wantErr)
			}
		})
	}
}
