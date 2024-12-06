package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/bearaujus/steam-utils/internal/config"
	"github.com/bearaujus/steam-utils/internal/model"
	"github.com/bearaujus/steam-utils/pkg/steam_acf"
	"github.com/bearaujus/steam-utils/pkg/steam_path"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

func NewSetAutoUpdate(ctx context.Context, cfg *config.Config) CmdRunner {
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
			return fmt.Errorf("error reading directory: %v", err)
		}
		var totalUpdate int
		var totalUpdateTarget int
		fmt.Println("-------------------------------------------------------------")
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if strings.HasSuffix(file.Name(), ".acf") {
				totalUpdateTarget++
				filename := filepath.Join(se.SteamApps(), file.Name())
				data, err := os.ReadFile(filename)
				if err != nil {
					return err
				}
				sa, err := steam_acf.Parse(data)
				if err != nil {
					return err
				}
				previousValue, err := sa.Update([]string{"AppState", "AutoUpdateBehavior"}, cmd.Use)
				if err != nil {
					return err
				}
				if previousValue == cmd.Use {
					fmt.Printf("File %v has already same configuration. Skipping...\n", filename)
					fmt.Println("-------------------------------------------------------------")
					continue
				}
				fmt.Printf("Updating %v...\n", filename)
				err = os.WriteFile(filename, sa.Serialize(), os.ModePerm)
				if err != nil {
					return err
				}
				totalUpdate++
				fmt.Printf("Changed %v -> %v\n", previousValue, cmd.Use)
				fmt.Println("-------------------------------------------------------------")
			}
		}
		msg := fmt.Sprintf("Sucessfully updated [%v/%v] files!\n", totalUpdate, totalUpdateTarget)
		if totalUpdate == 0 {
			msg = "No file were updated"
		}
		fmt.Println(msg)
		return nil
	}
}
