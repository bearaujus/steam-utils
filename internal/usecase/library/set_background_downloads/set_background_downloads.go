package set_background_downloads

import (
	"context"
	"errors"
	"fmt"
	"github.com/bearaujus/steam-utils/internal/config"
	"github.com/bearaujus/steam-utils/internal/model"
	"github.com/bearaujus/steam-utils/internal/pkg"
	"github.com/bearaujus/steam-utils/internal/usecase"
	"github.com/bearaujus/steam-utils/pkg/steam_acf"
	"github.com/bearaujus/steam-utils/pkg/steam_path"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

func NewCmdRunner(_ context.Context, cfg *config.Config) usecase.CmdRunner {
	return func(cmd *cobra.Command, args []string) error {
		se, err := steam_path.NewSteamPath(cfg.SteamPath)
		if err != nil {
			if errors.Is(err, steam_path.ErrEmptyPath) {
				return model.ErrSteamPathIsNotSet.New(config.PersistentFlagSteamPath)
			}
			return model.ErrSteamPathIsInvalid.New(err, config.PersistentFlagSteamPath)
		}

		files, err := os.ReadDir(se.SteamApps())
		if err != nil {
			return model.ErrReadDirectory.New(err)
		}

		var (
			aodwrUpdate, totalUpdate int
			aodwrTargets             = []string{"AppState", "AllowOtherDownloadsWhileRunning"}
			aodwrTargetsName         = strings.Join(aodwrTargets, ".")
		)
		pkg.PrintSep()
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if strings.HasSuffix(strings.ToLower(file.Name()), ".acf") {
				totalUpdate++
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

				fmt.Printf("Index\t: %v\nName\t: %v\nFile\t: %v\n", totalUpdate, appName, fileName)
				var aodwrPrevious string
				aodwrPrevious, err = sa.Update(aodwrTargets, cmd.Use)
				if err != nil {
					return model.ErrUpdateValueFromSteamACFFile.New(err)
				}

				if aodwrPrevious != cmd.Use {
					err = os.WriteFile(fileName, sa.Serialize(), os.ModePerm)
					if err != nil {
						return model.ErrWriteFile.New(err)
					}
				}
				if aodwrPrevious == cmd.Use {
					fmt.Printf("Action\t: No changes made. %v is already configured and up-to-date\n", aodwrTargetsName)
				} else {
					fmt.Printf("Action\t: Updated %v from %v -> %v\n", aodwrTargetsName, aodwrPrevious, cmd.Use)
					aodwrUpdate++
				}

				pkg.PrintSep()
			}
		}

		msg := fmt.Sprintf("Successfully updated %v: %d out of %d", aodwrTargetsName, aodwrUpdate, totalUpdate)
		if aodwrUpdate == 0 {
			msg = fmt.Sprintf("No files were updated for %v", aodwrTargetsName)
		}
		fmt.Println(msg)

		pkg.PrintSep()
		fmt.Printf("Applied\t: %v - %v\n", cmd.Use, cmd.Short)
		if aodwrUpdate != 0 {
			fmt.Println("To see the changes, please restart your Steam!")
		}

		return nil
	}
}
