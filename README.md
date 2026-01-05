# xkcdpw

A passphrase generator inspired by the [XKCD "Password Strength" comic](http://www.explainxkcd.com/wiki/index.php/936:_Password_Strength), creating high-entropy, easy-to-remember passphrases using acrostical word combinations.

This tool goes beyond simple random word selection by generating passphrases where each word starts with a specific letter, creating memorable patterns while maintaining strong security through high entropy.

## Features

- **Acrostic passphrases** - Generate phrases from a word (e.g., "test" becomes "Terrible-Exhausting-Scary-Tornado")
- **Random passphrases** - Create completely random acrostical phrases
- **Customizable formatting**:
  - Multiple capitalization modes (none, first, all, random)
  - Custom word separators
  - Optional number injection for additional entropy
- **High quality** - Uses curated word lists of adjectives and nouns
- **Secure** - Cryptographically random word selection

## Installation

### From Source

```bash
git clone https://github.com/awoodbeck/xkcdpw.git
cd xkcdpw
go build
```

### Requirements

- Go 1.25.5 or later

## Usage

### Quick Start

Generate passphrases with a random 4-letter word:

```bash
xkcdpw
```

Output:

```
Gorgeous-Radiant-Eager-Youthful
```

Generate a passphrase from a specific word:

```bash
xkcdpw hello
```

Output:

```
Handsome-Eager-Lively-Locked-Optimistic
```

Generate multiple passphrases:

```bash
xkcdpw 5
```

Output:

```
Brilliant-Zealous-Happy-Caring

Bright-Zesty-Harmonious-Charming

Bold-Zippy-Honest-Cheerful

Beautiful-Zany-Helpful-Creative

Brave-Zippy-Hearty-Calm
```

### Basic Syntax

```bash
xkcdpw [WORD] [NUMBER] [flags]
```

- **WORD** (optional): Word to use as acrostic. If omitted, a random 4-letter word is generated.
- **NUMBER** (optional): Number of passphrases to generate (default: 10, max: 100)

### Examples

**Generate 5 passphrases from "security":**

```bash
xkcdpw security 5
```

Output:

```
Safe-Energetic-Cheerful-Unique-Rapid-Intelligent-Thankful-Youthful

Stellar-Enormous-Creative-Upbeat-Resilient-Impressive-Timely-Young

Strong-Elegant-Careful-Unstoppable-Ready-Innovative-Thrilling-Yielding

Solid-Exciting-Confident-Ultimate-Reliable-Inspiring-Trustworthy-Yummy

Sturdy-Excellent-Calm-Useful-Robust-Interesting-Tremendous-Yellow
```

**Use spaces and lowercase:**

```bash
xkcdpw test 3 -d " " -c none
```

Output:

```
terrible exhausting scary tornado

tough energetic simple tremendous

tiny elegant sharp therapeutic
```

**Add random numbers for extra security:**

```bash
xkcdpw pass 3 -n --number-min 100 --number-max 999
```

Output:

```
Perfect-Awesome-Strong-Secure-847

Powerful-Amazing-Super-Safe-293

Practical-Agile-Solid-Stable-621
```

**Use underscores and all caps:**

```bash
xkcdpw safe 2 -d "_" -c all
```

Output:

```
STELLAR_AMAZING_FANTASTIC_EXCELLENT

SUPER_AWESOME_FABULOUS_EPIC
```

**Add numbers at the beginning:**

```bash
xkcdpw code 2 -n --number-min 1000 --number-max 9999 --number-position beginning
```

Output:

```
7284-Creative-Outstanding-Dependable-Excellent

3901-Confident-Optimal-Daring-Energetic
```

### Available Flags

| Flag                | Short | Default | Description                                           |
| ------------------- | ----- | ------- | ----------------------------------------------------- |
| `--delim`           | `-d`  | `-`     | Delimiter between words                               |
| `--capitalize`      | `-c`  | `first` | Capitalization mode: `none`, `first`, `all`, `random` |
| `--number`          | `-n`  | `false` | Add a random number to the passphrase                 |
| `--number-min`      |       | `0`     | Minimum value for random number                       |
| `--number-max`      |       | `99`    | Maximum value for random number                       |
| `--number-position` |       | `end`   | Number position: `end`, `beginning`, `random`         |

### Capitalization Modes

**none** - All lowercase

```bash
xkcdpw acrostic test 1 -c none
# terrible-exhausting-scary-tornado
```

**first** - Title case (default)

```bash
xkcdpw acrostic test 1 -c first
# Terrible-Exhausting-Scary-Tornado
```

**all** - All uppercase

```bash
xkcdpw acrostic test 1 -c all
# TERRIBLE-EXHAUSTING-SCARY-TORNADO
```

**random** - Random capitalization

```bash
xkcdpw acrostic test 1 -c random
# tErRiBlE-ExHaUsTiNg-sCaRy-ToRnAdO
```

## Use Cases

- **Password managers** - Generate memorable master passwords
- **System passphrases** - Easy to type, hard to crack
- **Mnemonic devices** - Create memorable phrases from acronyms
- **Creative inspiration** - Generate interesting word combinations

## Security Considerations

- Uses cryptographically secure random number generation
- High entropy through large word lists
- Each word adds approximately 10-13 bits of entropy
- A 4-word passphrase provides ~40-52 bits of entropy
- Adding numbers increases entropy further

**Example entropy calculation:**

- 4 random words: ~48 bits
- With 3-digit number (000-999): ~58 bits
- 5 words with 4-digit number: ~75 bits

## Credits

- Inspired by [XKCD #936: Password Strength](https://xkcd.com/936/)

## License

See [LICENSE](LICENSE) file for details.
