package styles

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
)

var Current Theme

//go:embed themes.json
var themesFromFile []byte

type Theme struct {
	Name       string `json:"name"`
	Covered    string `json:"green"`
	NotCovered string `json:"red"`
	Neutral    string `json:"selection"`
}

// DefaultTheme returns a default style.
func DefaultTheme() Theme {
	return Theme{
		Covered:    "#00ff00",
		NotCovered: "#ff0000",
		Neutral:    "#7f7f7f",
	}
}

// SetTheme sets the current theme.
func SetTheme() {
	name := os.Getenv("GOCOVSH_THEME")
	themes, _ := parseThemes(themesFromFile)
	Current = findTheme(themes, name)
}

// findTheme finds a theme.
func findTheme(themes []Theme, name string) Theme {
	for _, theme := range themes {
		if theme.Name == name {
			return theme
		}
	}

	return DefaultTheme()
}

// parseThemes converts the bytes from the json theme file to []Theme.
func parseThemes(bts []byte) ([]Theme, error) {
	var themes []Theme
	if err := json.Unmarshal(bts, &themes); err != nil {
		return nil, fmt.Errorf("could not load themes.json: %w", err)
	}

	return themes, nil
}
