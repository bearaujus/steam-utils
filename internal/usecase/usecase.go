package usecase

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/bearaujus/steam-utils/internal/model"
	"github.com/bearaujus/steam-utils/pkg/steam_acf"
	"github.com/bearaujus/steam-utils/pkg/steam_path"
)

func ListLibraryMetadata(_ context.Context, sp steam_path.SteamPath) ([]os.DirEntry, error) {
	files, err := os.ReadDir(sp.SteamApps())
	if err != nil {
		return nil, model.ErrReadDirectory.New(err)
	}

	var ret []os.DirEntry
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(strings.ToLower(file.Name()), ".acf") {
			ret = append(ret, file)
		}
	}

	if len(ret) == 0 {
		return nil, model.ErrEmptyListLibraryMetadata.New(sp.SteamApps())
	}

	return ret, nil
}

func SetLibraryMetadataAutoUpdate(ctx context.Context, sp steam_path.SteamPath, behaviour string) error {
	fileTargets, err := ListLibraryMetadata(ctx, sp)
	if err != nil {
		return err
	}

	var aubTargets = []string{"AppState", "AutoUpdateBehavior"}
	var sauTargets = []string{"AppState", "ScheduledAutoUpdate"}
	for _, file := range fileTargets {
		fileName := filepath.Join(sp.SteamApps(), file.Name())
		data, err := os.ReadFile(fileName)
		if err != nil {
			return model.ErrReadFile.New(err)
		}

		sa, err := steam_acf.Parse(data)
		if err != nil {
			return model.ErrParseSteamACFFile.New(err)
		}

		var aubPrevious, sauPrevious string
		aubPrevious, err = sa.Update(aubTargets, behaviour)
		if err != nil {
			return model.ErrUpdateValueFromSteamACFFile.New(err)
		}

		if behaviour == model.LibraryAutoUpdateOnlyOnLaunch {
			sauPrevious, err = sa.Update(sauTargets, "0")
			if err != nil {
				return model.ErrUpdateValueFromSteamACFFile.New(err)
			}
		}

		if aubPrevious != behaviour || (behaviour == model.LibraryAutoUpdateOnlyOnLaunch && sauPrevious != "0") {
			err = os.WriteFile(fileName, sa.Serialize(), os.ModePerm)
			if err != nil {
				return model.ErrWriteFile.New(err)
			}
		}
	}

	return nil
}

func SetLibraryMetadataBackgroundDownloads(ctx context.Context, sp steam_path.SteamPath, behaviour string) error {
	fileTargets, err := ListLibraryMetadata(ctx, sp)
	if err != nil {
		return err
	}

	var aodwrTargets = []string{"AppState", "AllowOtherDownloadsWhileRunning"}
	for _, file := range fileTargets {
		fileName := filepath.Join(sp.SteamApps(), file.Name())
		data, err := os.ReadFile(fileName)
		if err != nil {
			return model.ErrReadFile.New(err)
		}

		sa, err := steam_acf.Parse(data)
		if err != nil {
			return model.ErrParseSteamACFFile.New(err)
		}

		var aodwrPrevious string
		aodwrPrevious, err = sa.Update(aodwrTargets, behaviour)
		if err != nil {
			return model.ErrUpdateValueFromSteamACFFile.New(err)
		}

		if aodwrPrevious != behaviour {
			err = os.WriteFile(fileName, sa.Serialize(), os.ModePerm)
			if err != nil {
				return model.ErrWriteFile.New(err)
			}
		}
	}

	return nil
}
