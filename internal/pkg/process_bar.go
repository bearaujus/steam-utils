package pkg

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/schollz/progressbar/v3"
	"strings"
)

type ProgressBar interface {
	Add(msg string)
	Finish()
}

type progressBar struct {
	pName string
	bar   *progressbar.ProgressBar
}

func NewProgressBar(max int, pName string) ProgressBar {
	return &progressBar{
		pName: pName,
		bar:   progressbar.NewOptions(max + 1),
	}
}

func (p *progressBar) Add(msg string) {
	p.setDescription(fmt.Sprintf("Processing: %s...", msg))
	_ = p.bar.Add(1)
}

func (p *progressBar) Finish() {
	p.setDescription("All tasks completed successfully!")
	_ = p.bar.Add(1)
	_ = p.bar.Finish()
	fmt.Println(fmt.Sprintf("\nTask: %s", color.New(color.FgHiBlue).Sprint(p.pName)))
	PrintSep()
}

func (p *progressBar) setDescription(msg string) {
	m := maxWidth - 48
	if len(msg) > m {
		p.bar.Describe(msg[:m])
		return
	}
	p.bar.Describe(msg + strings.Repeat(" ", m-len(msg)))
}
