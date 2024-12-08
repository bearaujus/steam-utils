package steam_path

import (
	"fmt"
	"os"
	"path"
	"runtime"
)

type SteamPath interface {
	Base() string
	SteamApps() string
	String() string
}

type steamPath struct {
	basePath string
}

func NewSteamPath(location string) (SteamPath, error) {
	if location == "" {
		return nil, ErrEmptyPath
	}
	if _, err := os.Stat(location); err != nil {
		return nil, ErrInvalidPath
	}
	sp := steamPath{basePath: location}
	if err := sp.validate(); err != nil {
		return nil, err
	}
	return &sp, nil
}

func (sp *steamPath) Base() string {
	return sp.basePath
}

func (sp *steamPath) SteamApps() string {
	return path.Join(sp.basePath, "steamapps")
}

func (sp *steamPath) String() string {
	return sp.basePath
}

func (sp *steamPath) validate() error {
	paths := []string{
		sp.SteamApps(),
	}
	for _, v := range paths {
		if _, err := os.Stat(v); err != nil {
			return ErrInvalidPath
		}
	}
	return nil
}

func LoadDefaultSteamPath() (SteamPath, error) {
	var paths []string

	switch runtime.GOOS {
	case "windows":
		targetPath := []string{
			`Program Files (x86)/Steam`,
			`Program Files/Steam`,
		}
		for _, drive := range windowsListAvailableDrives() {
			for _, dir := range targetPath {
				paths = append(paths, path.Join(drive, dir))
			}
		}
	case "linux":
		paths = []string{
			os.ExpandEnv("$HOME/.steam/steam"),
			os.ExpandEnv("$HOME/.local/share/Steam"),
		}
	case "darwin": // macOS
		paths = []string{
			os.ExpandEnv("$HOME/Library/Application Support/Steam"),
		}
	default:
		return nil, ErrUnsupportedOS
	}

	for _, v := range paths {
		sp, err := NewSteamPath(v)
		if err != nil {
			continue
		}
		return sp, nil
	}

	return nil, ErrDefaultPathNotFound
}

func windowsListAvailableDrives() []string {
	var drives []string
	for letter := 'A'; letter <= 'Z'; letter++ {
		drive := fmt.Sprintf("%c:/", letter)
		if _, err := os.Stat(drive); err == nil {
			drives = append(drives, drive)
		}
	}
	return drives
}
