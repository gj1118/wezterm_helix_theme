# Change Helix Theme tool

This is a small tool written in golang that changes the theme of `Helix` whenever the `Wezterm` theme is changed. 

## Prerequisites 
You will need the following installed
1. Wezterm
2. Helix
3. Golang
4. [Just Task Runner](https://github.com/casey/just) (to make your life a lot easy).Also please make sure that `Just` is installed on your path 

On MacOs you can use `brew` to install `Just` using the following 
```brew install just```
## Configuring themes 
In the wezterm lua configuration file (located in the .config folder of your home path - e.g. `/Users/<userfolder>/.config/wezterm/wezterm.lua`) you can add the function below 

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

Notice that in `themeCycler` function above we have the list of the schemes that we will be targetting.  


Similarly in the `main.go` (in the folder where you checked out this repo) file we have the following map (Future verisons of this tool might make this configurable)

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

The key in the above map are the themes that you are using in the `themecycler` function. Their corresponding values are the `helix` theme you want to get applied when you switch to that `wezterm` theme.

## Building

After making sure that you have the prerequistes up and running , run the command `just bl`after you have checked out this repo.
`just bl` will perform the following functions 
1. Create a macos build
2. Make it executable
3. Copy the generated executable to the wezterm folder (`/Users/<userfolder>/.config/wezterm`)




When you change the wezterm theme using `Alt+t` (if you kept the same settings from this readme file) - then a helix theme (`wezterm_theme.toml`) is created/updated in the following location 
`/Users/<userfolder>/.config/helix/themes`
Now, in your helix editor, set yor configuration to load the following theme like so, 
`
  theme="wezterm_theme"
`
Now whever you change the wezterm theme the corresponding theme will be applied to your helix theme.


## Go version used 
1.23.0
