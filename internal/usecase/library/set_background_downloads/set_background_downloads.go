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
	"time"
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

		var fileTargets []os.DirEntry
		for _, file := range files {
			if !file.IsDir() && strings.HasSuffix(strings.ToLower(file.Name()), ".acf") {
				fileTargets = append(fileTargets, file)
			}
		}

		if len(fileTargets) == 0 {
			fmt.Printf("No .acf files detected in %v directory. Ensure that you have installed applications in your Steam library and try again.\n", se.SteamApps())
			return nil
		}

		var (
			aodwrUpdate      int
			aodwrTargets     = []string{"AppState", "AllowOtherDownloadsWhileRunning"}
			aodwrTargetsName = strings.Join(aodwrTargets, ".")
			bar              = pkg.NewProgressBar(len(fileTargets), cmd.Short)
		)
		for _, file := range fileTargets {
			time.Sleep(time.Millisecond * 50)
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
			bar.Add(appName)

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
				aodwrUpdate++
			}
		}

		bar.Finish()
		msg := fmt.Sprintf("Successfully updated %v: %d out of %d", aodwrTargetsName, aodwrUpdate, len(fileTargets))
		if aodwrUpdate == 0 {
			msg = fmt.Sprintf("No files were updated for %v", aodwrTargetsName)
		}
		fmt.Println(msg)

		if aodwrUpdate != 0 {
			pkg.PrintSep()
			fmt.Println("To see the changes, please restart your Steam!")
		}

		return nil
	}
}
