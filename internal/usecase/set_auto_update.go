package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/bearaujus/steam-utils/internal/config"
	"github.com/bearaujus/steam-utils/internal/model"
	"github.com/bearaujus/steam-utils/internal/pkg"
	"github.com/bearaujus/steam-utils/pkg/steam_acf"
	"github.com/bearaujus/steam-utils/pkg/steam_path"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

func NewSetAutoUpdate(_ context.Context, cfg *config.Config) CmdRunner {
	return func(cmd *cobra.Command, args []string) error {
		se, err := steam_path.NewSteamPath(cfg.SteamPath)
		if err != nil {
			if errors.Is(err, steam_path.ErrEmptyPath) {
				return model.ErrSteamPathIsNotSet.New(err, config.PersistentFlagSteamPath)
			}
			return model.ErrInvalidSteamPath.New(err, config.PersistentFlagSteamPath)
		}
		files, err := os.ReadDir(se.SteamApps())
		if err != nil {
			return model.ErrReadDirectory.New(err)
		}
		var totalUpdate, totalUpdateTarget int
		pkg.PrintSep()
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if strings.HasSuffix(file.Name(), ".acf") {
				totalUpdateTarget++
				fileName := filepath.Join(se.SteamApps(), file.Name())
				data, err := os.ReadFile(fileName)
				if err != nil {
					return model.ErrReadFile.New(err)
				}
				sa, err := steam_acf.Parse(data)
				if err != nil {
					return model.ErrParseSteamACFFile.New(err)
				}
				appName, err := sa.Get([]string{"AppState", "name"})
				if err != nil {
					return model.ErrGetValueFromSteamACFFile.New(err)
				}
				fmt.Printf("Index\t: %v\nName\t: %v\nFile\t: %v\n", totalUpdateTarget, appName, fileName)
				var updateTargets = []string{"AppState", "AutoUpdateBehavior"}
				previousValue, err := sa.Update(updateTargets, cmd.Use)
				if err != nil {
					return model.ErrUpdateValueFromSteamACFFile.New(err)
				}
				if previousValue == cmd.Use {
					fmt.Println("Action\t: No changes made. The file is already configured and up-to-date.")
					pkg.PrintSep()
					continue
				}
				err = os.WriteFile(fileName, sa.Serialize(), os.ModePerm)
				if err != nil {
					return model.ErrWriteFile.New(err)
				}
				totalUpdate++
				fmt.Printf("Action\t: Updated %v from \"%v\" to \"%v\".\n", strings.Join(updateTargets, "."), previousValue, cmd.Use)
				pkg.PrintSep()
			}
		}
		msg := fmt.Sprintf("Successfully updated %d out of %d files!\n", totalUpdate, totalUpdateTarget)
		if totalUpdate == 0 {
			msg = "No files were updated."
		}
		fmt.Println(msg)
		return nil
	}
}
