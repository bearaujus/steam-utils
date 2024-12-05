package main

import (
	"fmt"
	"github.com/bearaujus/steam-utils/pkg/steam_acf"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Directory to scan
	dir := "D:/Program Files (x86)/Steam/steamapps"

	// Read the directory contents (does not recurse)
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// Iterate over the files in the directory
	for _, file := range files {
		// Skip if it's a directory
		if file.IsDir() {
			continue
		}

		// Check if the file has a .acf extension
		if strings.HasSuffix(file.Name(), ".acf") {
			// Process the .acf file (e.g., print its name)
			fmt.Println("Found .acf file:", filepath.Join(dir, file.Name()))
			err = UpdateACF(filepath.Join(dir, file.Name()))
			if err != nil {
				fmt.Println("Error updating ACF:", err)
			}
			fmt.Println("----------------------------")
		}
	}

	//var rootCmd = cmd.New(context.TODO(), &config.Config{})
	//rootCmd.Execute()
}

func UpdateACF(name string) error {
	// 0 always keep this game updated
	// 1 Only update this game when I launch it
	// 2 High Priority - Always auto-update this game before others
	data, err := os.ReadFile(name)
	if err != nil {
		return err
	}

	sa, err := steam_acf.Parse(data)
	if err != nil {
		return err
	}

	err = sa.Update([]string{"AppState", "AutoUpdateBehavior", "qwe"}, "1")
	err = sa.Update([]string{"AppState", "AutoUpdateBehavior"}, "1")
	if err != nil {
		return err
	}

	d, err := sa.Serialize()
	if err != nil {
		return err
	}

	err = os.WriteFile(fmt.Sprintf("./dev/%v", filepath.Base(name)), d, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
