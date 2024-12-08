package cli

import (
	"context"
	"fmt"

	"github.com/bearaujus/steam-utils/internal/config"
	"github.com/bearaujus/steam-utils/internal/model"
	"github.com/bearaujus/steam-utils/internal/pkg"
	"github.com/bearaujus/steam-utils/internal/usecase"
	"github.com/bearaujus/steam-utils/internal/view"
	"github.com/bearaujus/steam-utils/pkg/steam_path"
	"github.com/spf13/cobra"
)

const (
	PersistentFlagSteamPath = "steam-path"
)

type cli struct {
	cfg          *config.Config
	root         *cobra.Command
	rawSteamPath string
}

func New(ctx context.Context, cfg *config.Config) view.View {
	app := &cli{cfg: cfg}
	app.root = app.rootCmd(ctx)
	var rawDefaultSteamPath string
	if app.cfg.DefaultSteamPath != nil {
		rawDefaultSteamPath = app.cfg.DefaultSteamPath.String()
	}
	app.root.PersistentFlags().StringVar(&app.rawSteamPath, PersistentFlagSteamPath, rawDefaultSteamPath, "Path to steam installation directory")
	return app
}

func (c *cli) Run(ctx context.Context) error {
	pkg.PrintTitle(c.cfg)
	c.root.SetContext(ctx)
	if err := c.root.Execute(); err != nil {
		return err
	}
	return nil
}

func (c *cli) rootCmd(ctx context.Context) *cobra.Command {
	root := &cobra.Command{
		Use:               c.cfg.LdFlags.Name,
		Args:              cobra.NoArgs,
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true, DisableNoDescFlag: true, DisableDescriptions: true, HiddenDefaultCmd: true},
	}
	root.AddCommand(c.libraryCmd(ctx))
	return root
}

func (c *cli) libraryCmd(ctx context.Context) *cobra.Command {
	libraryCmd := &cobra.Command{
		Use:   "library",
		Short: "Steam library utilities",
		Args:  cobra.NoArgs,
	}
	libraryCmd.AddCommand(
		c.librarySetAutoUpdateCmd(ctx),
		c.librarySetBackgroundDownloadsCmd(ctx),
	)
	return libraryCmd
}

func (c *cli) librarySetAutoUpdateCmd(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-auto-update",
		Short: "Set auto update behavior on all collections in your Steam library",
		Args:  cobra.NoArgs,
	}
	for k, v := range model.LibraryAutoUpdate {
		runner := func(cmd *cobra.Command, args []string) error {
			sp, err := c.getSteamPath()
			if err != nil {
				return err
			}
			err = usecase.SetLibraryMetadataAutoUpdate(ctx, sp, k)
			if err != nil {
				return err
			}
			printSuccessMessage(fmt.Sprintf("%v (%v)", cmd.CommandPath(), v))
			return nil
		}
		cmd.AddCommand(c.cmdRunner(k, v, runner))
	}
	return cmd
}

func (c *cli) librarySetBackgroundDownloadsCmd(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-background-downloads",
		Short: "Set background downloads behavior on all collections in your Steam library",
		Args:  cobra.NoArgs,
	}
	for k, v := range model.LibraryBackgroundDownloads {
		runner := func(cmd *cobra.Command, args []string) error {
			sp, err := c.getSteamPath()
			if err != nil {
				return err
			}
			err = usecase.SetLibraryMetadataBackgroundDownloads(ctx, sp, k)
			if err != nil {
				return err
			}
			printSuccessMessage(v)
			return nil
		}
		cmd.AddCommand(c.cmdRunner(k, v, runner))
	}
	return cmd
}

func (c *cli) cmdRunner(option string, description string, runner func(cmd *cobra.Command, args []string) error) *cobra.Command {
	return &cobra.Command{
		Use:   option,
		Short: description,
		Args:  cobra.NoArgs,
		RunE:  runner,
	}
}

func (c *cli) getSteamPath() (steam_path.SteamPath, error) {
	if c.rawSteamPath == "" && c.cfg.DefaultSteamPath != nil {
		return c.cfg.DefaultSteamPath, nil
	}
	if c.rawSteamPath == "" && c.cfg.DefaultSteamPath != nil {
		return nil, model.ErrFailToInitializeSteamPath.New(fmt.Sprintf("application is unable to determine steam path. please specify the steam path using flag '%v'", PersistentFlagSteamPath))
	}
	sp, err := steam_path.NewSteamPath(c.rawSteamPath)
	if err != nil {
		return nil, model.ErrFailToInitializeSteamPath.New(err)
	}
	return sp, nil
}

func printSuccessMessage(task string) {
	fmt.Println(fmt.Sprintf("Applied: '%v'", task))
	fmt.Println("Success! To see the changes, please restart your Steam.")
	fmt.Println()
}
