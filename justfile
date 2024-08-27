build:
	env GOOS=darwin GOARCH=amd64 go build
	chmod +x ./change_term_theme

relocate:
	cp change_term_theme /Users/gjanjua/.config/wezterm/
	# this will overrite the existing toml file and creates some confusion , hence disabling this for now
	# cp helix_theme.toml /Users/gjanjua/.config/wezterm

bl: build relocate
	echo "built and relocated"
