package interactive

import (
	"context"
	"fmt"

	"github.com/bearaujus/steam-utils/internal/config"
	"github.com/bearaujus/steam-utils/internal/model"
	"github.com/bearaujus/steam-utils/internal/pkg"
	"github.com/bearaujus/steam-utils/internal/usecase"
	"github.com/bearaujus/steam-utils/internal/view"
	"github.com/bearaujus/steam-utils/pkg/steam_path"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type interactive struct {
	cfg *config.Config
	app *tview.Application
}

func New(_ context.Context, cfg *config.Config) view.View {
	return &interactive{
		cfg: cfg,
		app: tview.NewApplication(),
	}
}

func (it *interactive) Run(ctx context.Context) error {
	if it.cfg.DefaultSteamPath != nil {
		it.useDefaultSteamPathPromptCmd(ctx, nil)
	} else {
		it.setSteamPathCmd(ctx, nil)
	}
	return it.app.Run()
}

func (it *interactive) useDefaultSteamPathPromptCmd(ctx context.Context, parent tview.Primitive) {
	cmd := tview.NewModal()
	cmd.SetText(fmt.Sprintf("Default Steam installation path detected at:\n%v\nDo you want to use this?", it.cfg.DefaultSteamPath.Base()))
	cmd.AddButtons([]string{"Yes", "No"})
	cmd.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonLabel == "Yes" {
			it.cfg.SteamPath = it.cfg.DefaultSteamPath
			it.rootCmd(ctx)
			return
		}
		it.setSteamPathCmd(ctx, parent)
	})
	it.setRoot(cmd)
}

func (it *interactive) setSteamPathCmd(ctx context.Context, parent tview.Primitive) {
	label := "Enter Steam Installation Path: "
	cmd := tview.NewForm()
	cmd.AddInputField(label, "", 0, nil, nil)
	cmd.AddButton("Save", func() {
		steamPath := cmd.GetFormItemByLabel(label).(*tview.InputField).GetText()
		sp, err := steam_path.NewSteamPath(steamPath)
		if err != nil {
			it.showErrorModal(ctx, model.ErrFailToInitializeSteamPath.New(err).Error(), cmd)
			return
		}
		it.cfg.SteamPath = sp
		it.rootCmd(ctx)
	})
	if it.cfg.DefaultSteamPath != nil {
		cmd.AddButton("Use Default", func() { it.useDefaultSteamPathPromptCmd(ctx, parent) })
	}
	if parent == nil {
		cmd.AddButton("Quit", func() { it.app.Stop() })
	} else {
		cmd.AddButton("Back", func() { it.setRoot(parent) })
	}
	enableFormFocusByArrow(cmd)
	it.setRoot(cmd)
}

func (it *interactive) rootCmd(ctx context.Context) {
	cmd := tview.NewList()
	cmd.AddItem(it.libraryCmd(ctx, cmd))
	cmd.AddItem(it.optionsCmd(ctx, cmd))
	it.addQuit(cmd)
	it.setRoot(cmd)
}

func (it *interactive) libraryCmd(ctx context.Context, parent tview.Primitive) (string, string, rune, func()) {
	return "Library", "Steam library utilities", '1', func() {
		cmd := tview.NewList()
		cmd.AddItem(it.librarySetAutoUpdateCmd(ctx, cmd)).
			AddItem(it.librarySetBackgroundDownloadsCmd(ctx, cmd))
		it.addBack(cmd, parent)
		it.setRoot(cmd)
	}
}

func (it *interactive) optionsCmd(ctx context.Context, parent tview.Primitive) (string, string, rune, func()) {
	return "Options", "", 'o', func() {
		cmd := tview.NewList()
		cmd.AddItem("Set Steam path", it.cfg.SteamPath.String(), '1', func() {
			it.setSteamPathCmd(ctx, cmd)
		})
		it.addBack(cmd, parent)
		it.setRoot(cmd)
	}
}

func (it *interactive) librarySetAutoUpdateCmd(ctx context.Context, parent tview.Primitive) (string, string, rune, func()) {
	return "Set auto update", "Set auto update behavior on all collections in your Steam library", '1', func() {
		cmd := tview.NewList()
		idx := 0
		for k, v := range model.LibraryAutoUpdate {
			idx++
			cmd.AddItem(v, "", rune(idx+'0'), func() {
				it.showModalWithoutButton(ctx, "Processing...")
				if err := usecase.SetLibraryMetadataAutoUpdate(ctx, it.cfg.SteamPath, k); err != nil {
					it.showErrorModal(ctx, err.Error(), cmd)
					return
				}
				it.showModal(ctx, "Success!\nTo see the changes, please restart your Steam.", nil)
			})
		}
		it.addBack(cmd, parent)
		it.setRoot(cmd)
	}
}

func (it *interactive) librarySetBackgroundDownloadsCmd(ctx context.Context, parent tview.Primitive) (string, string, rune, func()) {
	return "Set background downloads", "Set background downloads behavior on all collections in your Steam library", '2', func() {
		cmd := tview.NewList()
		idx := 0
		for k, v := range model.LibraryBackgroundDownloads {
			idx++
			cmd.AddItem(v, "", rune(idx+'0'), func() {
				it.showModalWithoutButton(ctx, "Processing...")
				if err := usecase.SetLibraryMetadataBackgroundDownloads(ctx, it.cfg.SteamPath, k); err != nil {
					it.showErrorModal(ctx, err.Error(), cmd)
					return
				}
				it.showModal(ctx, "Success!\nTo see the changes, please restart your Steam.", nil)
			})
		}
		it.addBack(cmd, parent)
		it.setRoot(cmd)
	}
}

func (it *interactive) setRoot(cmd tview.Primitive) {
	// Wrap the cmd in a Frame for padding
	frame := tview.NewFrame(cmd).
		SetBorders(1, 1, 0, 0, 1, 1) // Padding: Top, Bottom, Left, Right, Border Width

	// Configure the flex container
	flex := tview.NewFlex()
	flex.SetTitle(fmt.Sprintf(" %v ", pkg.GetTitleRaw(it.cfg)))
	flex.SetDirection(tview.FlexRow)
	flex.SetBorder(true)
	flex.AddItem(frame, 0, 1, false)

	// Set the root and focus
	it.app.SetRoot(flex, true)
	it.app.SetFocus(cmd)
}

func (it *interactive) addBack(cmd *tview.List, backPage tview.Primitive) {
	cmd.AddItem("Back", "", '0', func() {
		it.setRoot(backPage)
	})
}

func (it *interactive) addQuit(cmd *tview.List) {
	cmd.AddItem("Quit", "", 'q', func() {
		it.app.Stop()
	})
}

func (it *interactive) showModalWithoutButton(_ context.Context, text string) {
	cmd := tview.NewModal()
	cmd.SetText(text)
	it.setRoot(cmd)
}

func (it *interactive) showModal(ctx context.Context, text string, backPage tview.Primitive) {
	cmd := tview.NewModal()
	cmd.SetText(text)
	cmd.AddButtons([]string{"Ok"})
	cmd.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if backPage == nil {
			it.rootCmd(ctx)
			return
		}
		it.setRoot(backPage)
	})
	it.setRoot(cmd)
}

func (it *interactive) showErrorModal(ctx context.Context, text string, backPage tview.Primitive) {
	it.showModal(ctx, fmt.Sprintf("Error: %v", text), backPage)
}

func enableFormFocusByArrow(cmd *tview.Form) {
	i, fc, bc := 0, cmd.GetFormItemCount(), cmd.GetButtonCount()
	t := fc + bc // Total focusable items
	cmd.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyUp:
			// Move focus upwards
			if i > 0 {
				// Skip over buttons if they're at the bottom row
				if i >= fc {
					i = fc - 1 // Jump to the last form item
				} else {
					i-- // Move up within form items
				}
				cmd.SetFocus(i)
			}
		case tcell.KeyDown:
			// Move focus downwards
			if i < t-1 {
				// Jump to buttons if currently at the last form item
				if i < fc-1 {
					i++
				} else {
					i = fc // Jump to the first button
				}
				cmd.SetFocus(i)
			}
		case tcell.KeyRight:
			// Move right through buttons
			if i >= fc && i < t-1 {
				i++
				cmd.SetFocus(i)
			}
		case tcell.KeyLeft:
			// Move left through buttons
			if i > fc {
				i--
				cmd.SetFocus(i)
			}
		default:
			return event
		}
		return nil
	})
}
