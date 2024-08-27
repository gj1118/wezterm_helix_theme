# Change Helix Theme tool

This is a small tool written in golang that changes the theme of `Helix` whenever the `Wezterm` theme is changed. 

## Building 
You will need [Just Task Runner](https://github.com/casey/just) (to make your life a lot easy). Make sure that `Just` is installed on your path 

On MacOs you can use `brew` to install `Just` using the following 
```brew install just```

If you have Just installed, run the command `just bl` after you have checked out this repo, to build and copy the artifacts to the wezterm folder. The helix folder that it will copy the generated `helix` theme to is `/Users/<userfolder>/.config/helix/themes`

## Configuring themes 
In the wezterm lua configuiration file (located in the .config folder of your home path - also listed above) you can add the function below 

```
  local function themeCycler(window, _)
  local currentScheme = window:effective_config().color_scheme
  local schemes = { "Catppuccin Latte","Macintosh (base16)", 'Catppuccin Macchiato', 'Catppuccin Mocha', 'Everforest Light (Gogh)','Tokyo Night','Tokyo Night Moon',"Monokai Soda","Kanagawa (Gogh)" }
  -- print("will print some data")
  -- print(data["theme"])
  for i = 1, #schemes, 1 do
    if schemes[i] == currentScheme then
      local overrides = window:get_config_overrides() or {}
      local next = i % #schemes + 1
      overrides.color_scheme = schemes[next]
      wezterm.log_info("Switched to: " .. schemes[next])
      window:set_config_overrides(overrides)

      -- set the helix theme , we are using a custom golang binary for the same. 
      -- we pass the name of the binary to the executable file and then let it 
      -- run the business 
      local success, stdout, stderr = wezterm.run_child_process { HELIX_THEME_CHANGER_LOCATION, schemes[next] }
      print(success)
      print(stdout)
      print(stderr)
      return
    end
  end
end
```

and then you can use the following keycombination to cycle through your themes 
```
  { key = "t",          mods = "ALT",  action = wezterm.action_callback(themeCycler) },
  
```

Notice that in `themeCycler` function above we have the list of the schemes that we are targetting.  


Similarly in the `main.go` file we have the following map 

```
  var commits = map[string]string{
	"Catppuccin Latte":     "catppuccin_latte",
	"Catppuccin Macchiato": "catppuccin_macchiato",
	"Catppuccin Mocha":     "catppuccin_mocha",
	"Everforest Light (Gogh)":    "solarized_light",
	"Macintosh (base16)":    "curzon",
	"Tokyo Night":          "tokyonight",
	"Tokyo Night Moon":     "tokyonight_moon",
	"Monokai Soda" : "monokai_soda",
	"Kanagawa (Gogh)" : "kanagawa",
}

```

This key is the theme that you are using in the wezterm.lua config file and in the `themecycler` function. The value is the `helix` theme you want to get applied when you switch to the corresponding `wezterm` theme.

When you change the wezterm theme using `Alt+t` (if you kept the same settings from this readme file) - then a helix theme (`wezterm_theme.toml`) is created/updated in the following location 
`/Users/<userfolder>/.config/helix/themes`
Now, in your helix editor, set yor configuration to load the following theme like so, 
`
  theme="wezterm_theme"
`
Now whever you change the wezterm theme the corresponding theme will be applied to your helix theme.


## Go version used 
1.23.0
