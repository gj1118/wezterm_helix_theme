package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/BurntSushi/toml"
)

var commits = map[string]string{
	"Catppuccin Latte":     "catppuccin_latte",
	"Catppuccin Macchiato": "catppuccin_macchiato",
	"Catppuccin Mocha":     "catppuccin_mocha",
	"Everforest Light (Gogh)":    "everforest_light",
	"Macintosh (base16)":    "curzon",
	"Tokyo Night":          "tokyonight",
	"Tokyo Night Moon":     "tokyonight_moon",
	"Monokai Soda" : "monokai_pro_spectrum",
	"Kanagawa (Gogh)" : "kanagawa",
	"Modus-Operandi":"modus_operandi_tinted",
	"rose-pine-dawn":"rose_pine_dawn",
	"rose-pine":"rose_pine",
	"Tomorrow":"spacebones_light",
	"Sonokai (Gogh)":"sonokai",
}

const HELIX_TOML_THEME_FILE_LOCATION = "/Users/gjanjua/.config/helix/themes/wezterm_theme.toml"

func main() {
	programArgs := os.Args[1:]
	fmt.Println(programArgs)
	inheritedTheme := "default"
	if len(programArgs) > 0 {
		inheritedTheme = programArgs[0]
	}
	fmt.Println(inheritedTheme)
	content, err := os.ReadFile("/Users/gjanjua/.config/wezterm/helix_theme.toml")
	if err != nil {
		panic(fmt.Sprintf("Error reading file:%s", err))
	}

	// Parse the TOML content
	var data map[string]interface{}
	if _, err := toml.Decode(string(content), &data); err != nil {
		panic("Error while parsing toml")
	}
	value, exists := commits[inheritedTheme]
	if exists == false {
		fmt.Println("we haven't yet have a corresponding theme for →", inheritedTheme)
		value = "default"
	}
	fmt.Println(value)
	data["inherits"] = value
	// Format the TOML data
	formatted := formatTOML(data)

	// Write the formatted string to a file
	if err := os.WriteFile(HELIX_TOML_THEME_FILE_LOCATION, []byte(formatted), 0644); err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("TOML file updated successfully")
	// writeNewFile("config.toml", "inherits", "gagan", tree)
}

func formatTOML(data map[string]interface{}) string {
	var result strings.Builder
	keys := make([]string, 0, len(data))
	for k := range data {
		if k != "palette" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	for _, key := range keys {
		formatKeyValue(&result, key, data[key], 0)
	}

	// Handle the palette section
	if palette, ok := data["palette"].(map[string]interface{}); ok {
		result.WriteString("\n[palette]\n")
		paletteKeys := make([]string, 0, len(palette))
		for k := range palette {
			paletteKeys = append(paletteKeys, k)
		}
		sort.Strings(paletteKeys)
		for _, key := range paletteKeys {
			fmt.Fprintf(&result, "%s = %q\n", key, palette[key])
		}
	}

	return result.String()
}

func formatKeyValue(result *strings.Builder, key string, value interface{}, indent int) {
	switch v := value.(type) {
	case string:
		if strings.Contains(key, "ui.") {
			fmt.Println(key)
			fmt.Fprintf(result, "\n%s%q = %q", strings.Repeat(" ", indent), key, v)

		} else {
			fmt.Fprintf(result, "%s%q = %q", strings.Repeat(" ", indent), key, v)
		}
	case map[string]interface{}:
		fmt.Fprintf(result, "\n%s%q = {", strings.Repeat("", indent), key)
		keys := make([]string, 0, len(v))
		for k := range v {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for i, k := range keys {
			if i > 0 {
				result.WriteString(",")
			}
			formatKeyValue(result, k, v[k], indent+1)
		}
		result.WriteString("}\n")
	default:
		fmt.Fprintf(result, "%s%q = %v\n", strings.Repeat("", indent), key, v)
	}
}
